package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint   `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func (u *User) Hashed() {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 0)
	u.Password = string(hashed)
}
