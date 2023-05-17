package auth

import (
	"context"
)

type Server struct {
	UserStorageServer
}

func (s *Server) Auth(ctx context.Context, cred *Credentials) (*Session, error) {
	return &Session{}, nil
}

func (s *Server) Register(ctx context.Context, user *User) (*UserId, error) {
	return &UserId{}, nil
}
