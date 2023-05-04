package app

import (
	"lazzytchk/council/internal/data"
	"net/http"
)

type Builder interface {
	Build() *Server
}

type ServerBuilder struct {
	Server *Server
}

func (b *ServerBuilder) ConfigureServer(s *http.Server) {
	router := http.NewServeMux()

	router.HandleFunc("/auth", b.Server.Auth())
	router.HandleFunc("/register", b.Server.Register())

	s.Handler = router

	b.Server.Server = s
}

func (b *ServerBuilder) ConfigureIdentifier(identifier data.Identifier) {
	b.Server.identifier = identifier
}

func (b *ServerBuilder) ConfigureRegistrar(registrar data.Registrar) {
	b.Server.registrar = registrar
}

func (b *ServerBuilder) Build() *Server {
	return b.Server
}
