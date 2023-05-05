package session

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"lazzytchk/council/internal/model"
	"time"
)

type Token struct {
	User       model.User
	ExpireTime time.Duration
}

func (t Token) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t Token) Hash() string {
	str := t.User.Email
	result, _ := bcrypt.GenerateFromPassword([]byte(str), 0)

	return string(result)
}
