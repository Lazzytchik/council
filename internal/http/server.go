package http

import (
	"encoding/json"
	"errors"
	"lazzytchk/council/internal/data"
	"lazzytchk/council/internal/model"
	"log"
	"net/http"
)

type AuthServer interface {
	Auth(i data.Identifier) http.HandlerFunc
	Register(rg data.Registrar) http.HandlerFunc
}

type Server struct {
	*http.Server
}

func NewServer(addr string, errLogger *log.Logger) *Server {
	s := Server{}

	router := http.NewServeMux()
	router.HandleFunc("/auth", s.Auth(data.Mock{}))
	router.HandleFunc("/register", s.Register(data.Mock{}))

	s.Server = &http.Server{
		Addr:     addr,
		Handler:  router,
		ErrorLog: errLogger,
	}

	return &s
}

func (s *Server) Auth(i data.Identifier) http.HandlerFunc {
	type request struct {
		email    string `json:"email"`
		password string `json:"password"`
	}

	type response struct {
		Status string `json:"status"`
		Id     uint   `json:"result"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, 400, errors.New("invalid parameters"))
		}

		id, err := i.Identify(req.email, req.password)
		if err != nil {
			s.error(w, r, 403, errors.New("no user with given credentials"))
		}

		s.respond(w, r, 200, response{
			Status: "OK",
			Id:     id,
		})
	}
}

func (s *Server) Register(rg data.Registrar) http.HandlerFunc {
	type response struct {
		Status string `json:"status"`
		Id     uint   `json:"result"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			s.error(w, r, 400, errors.New("invalid parameters"))
		}

		id, err := rg.Register(*user)
		if err != nil {
			s.error(w, r, 403, err)
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
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
