package http

import (
	"lazzytchk/council/internal/model"
	"time"
)

type Token struct {
	User       model.User
	ExpireTime time.Time
}

type Session interface {
	Get(token string) (Token, error)
	Save(token Token) (string, error)
}
