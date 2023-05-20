package data

import (
	"github.com/lazzytchik/council/internal/model"
	"math/rand"
)

type Identifier interface {
	Identify(username, password string) (model.User, error)
}

type Registrar interface {
	Register(user model.User) (uint, error)
}

type Mock struct {
}

func (m Mock) Identify(email, password string) (uint, error) {
	return uint(rand.Int()), nil
}

func (m Mock) Register(user model.User) (uint, error) {
	return uint(rand.Int()), nil
}
