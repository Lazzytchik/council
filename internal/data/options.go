package data

import "fmt"

type ConnOptions struct {
	Name     string
	Host     string
	User     string
	Port     string
	Password string
}

func (c ConnOptions) ConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
	)
}

func (c ConnOptions) Describe() string {
	return fmt.Sprintf(
		"NAME: %s \nHOST: %s \nPORT: %s \nUSER: %s \nPASSWORD: %s",
		c.Name,
		c.Host,
		c.Port,
		c.User,
		c.Password,
	)
}
