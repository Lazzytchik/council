package main

import (
	"context"
	"lazzytchk/council/auth/grpc"
)

type server struct {
}

func (s *server) Auth(ctx context.Context, cred grpc.Credentials) (grpc.Session, error) {
	return grpc.Session{}, nil
}

func (s *server) Register(ctx context.Context, user grpc.User) (grpc.UserId, error) {
	return grpc.UserId{}, nil
}
