package session

import "lazzytchk/council/internal/model"

type Redis struct {
}

func (r Redis) Save(user model.User) string {
	return ""
}
