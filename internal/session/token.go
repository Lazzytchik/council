package session

import (
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
