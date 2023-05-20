package auth

import (
	"context"
	"github.com/lazzytchik/council/internal/data"
	"github.com/lazzytchik/council/internal/model"
	"github.com/lazzytchik/council/internal/session"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Server struct {
	UserStorageServer
	identifier data.Identifier
	registrar  data.Registrar
	session    session.Session

	errLogger *log.Logger
}

func (s *Server) Auth(ctx context.Context, cred *Credentials) (*Session, error) {
	user, err := s.identifier.Identify(cred.Email, cred.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while identifying user", err)
	}

	token, err := s.session.Save(session.Token{
		User: user,
	})

	return &Session{Token: token}, status.New(codes.OK, "").Err()
}

func (s *Server) Register(ctx context.Context, user *User) (*UserId, error) {

	u := &model.User{
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}
	u.Hashed()

	id, err := s.registrar.Register(*u)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot register user", err)
	}

	return &UserId{Id: uint32(id)}, status.New(codes.OK, "").Err()
}
