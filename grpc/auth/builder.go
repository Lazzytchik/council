package auth

import (
	"lazzytchk/council/internal/data"
	"lazzytchk/council/internal/session"
	"log"
)

type Builder interface {
	Build() *Server
}

type ServerBuilder struct {
	Server Server
}

func (b *ServerBuilder) ConfigureIdentifier(identifier data.Identifier) {
	b.Server.identifier = identifier
}

func (b *ServerBuilder) ConfigureRegistrar(registrar data.Registrar) {
	b.Server.registrar = registrar
}

func (b *ServerBuilder) ConfigureSession(s session.Session) {
	b.Server.session = s
}

func (b *ServerBuilder) ConfigureLogger(l *log.Logger) {
	b.Server.errLogger = l
}

func (b *ServerBuilder) Build() *Server {
	return &b.Server
}
