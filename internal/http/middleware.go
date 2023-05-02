package http

import "net/http"

func (s *Server) SessionCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Session check
		next(w, r)
	}
}
