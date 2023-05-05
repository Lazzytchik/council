package session

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	client *redis.Client
	Expire time.Duration
}

func NewRedis(options *redis.Options, expire time.Duration) *Redis {
	client := redis.NewClient(options)

	return &Redis{
		client: client,
		Expire: expire,
	}
}

func (r *Redis) Save(token Token) (string, error) {
	key := token.Hash()

	return key, r.client.Set(key, token, r.Expire).Err()
}

func (r *Redis) Get(hash string) (Token, error) {
	var token Token

	res, err := r.client.Get(hash).Result()

	if err != nil {
		return token, errors.Join(errors.New("redis read error"), err)
	}

	err = json.Unmarshal([]byte(res), &token)

	if err != nil {
		return token, errors.Join(errors.New("cannot unmarshal redis result into token"), err)
	}

	return token, err
}
