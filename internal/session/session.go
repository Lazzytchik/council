package session

type Session interface {
	Get(token string) (Token, error)
	Save(token Token) (string, error)
}
