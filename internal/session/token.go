package session

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
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
	str := t.User.Email + t.User.Password

	return hex.EncodeToString(md5.New().Sum([]byte(str)))
}
