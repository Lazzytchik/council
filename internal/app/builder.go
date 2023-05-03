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

	Identifier data.Identifier
	Registrar  data.Registrar
	Session    Session

	HttpServer *http.Server
}

func (b *ServerBuilder) ConfigureServer() {

}

func (b *ServerBuilder) Build() *Server {
	return b.Server
}
