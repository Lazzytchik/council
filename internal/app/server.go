package app

import (
	"encoding/json"
	"errors"
	"github.com/lazzytchk/council/internal/data"
	"github.com/lazzytchk/council/internal/model"
	"github.com/lazzytchk/council/internal/session"
	"net/http"
)

type AuthServer interface {
	Auth() http.HandlerFunc
	Register() http.HandlerFunc
}

type Server struct {
	*http.Server
	identifier data.Identifier
	registrar  data.Registrar
	session    session.Session
}

func (s *Server) Auth() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type response struct {
		Status  string `json:"status"`
		Session string `json:"result"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, 400, errors.New("invalid parameters"))
			return
		}

		cookie, err := r.Cookie("session")
		if err == nil {
			_, err := s.session.Get(cookie.Value)
			if err != nil {
				s.ErrorLog.Println("session reading error: ", err)
			} else {
				s.respond(w, r, 200, response{
					Status:  "OK",
					Session: cookie.Value,
				})
				return
			}

		} else if !errors.Is(err, http.ErrNoCookie) {
			s.error(w, r, 500, errors.New("internal session error"))
			return
		}

		user, err := s.identifier.Identify(req.Email, req.Password)
		if err != nil {
			s.error(w, r, 403, errors.New("no user with given credentials"))
			return
		}

		token, err := s.session.Save(session.Token{
			User: user,
		})

		if err != nil {
			s.error(w, r, 500, errors.New("cant make session"))
		}

		s.setSessionCookie(w, r, token)
		s.respond(w, r, 200, response{
			Status:  "OK",
			Session: token,
		})
	}
}

func (s *Server) Register() http.HandlerFunc {
	type response struct {
		Status string `json:"status"`
		Id     uint   `json:"result"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			s.error(w, r, 400, errors.New("invalid parameters"))
			return
		}

		user.Hashed()

		id, err := s.registrar.Register(*user)
		if err != nil {
			s.error(w, r, 403, err)
			return
		}

		s.respond(w, r, 200, response{
			Status: "OK",
			Id:     id,
		})
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *Server) setSessionCookie(w http.ResponseWriter, r *http.Request, token string) {
	cookie := http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "/auth",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
}
