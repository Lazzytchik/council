package data

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"lazzytchk/council/internal/model"
	"log"
)

func NewPostgres(options ConnOptions, l *log.Logger) *Postgres {
	return &Postgres{options, l}
}

type Postgres struct {
	ConnOptions
	Logger *log.Logger
}

func (p Postgres) Identify(email, password string) (uint, error) {
	db, err := p.Connect()
	if err != nil {
		p.Logger.Println(err)
		return 0, err
	}
	defer db.Close()

	sql := "SELECT * FROM users WHERE email = $1 LIMIT 1"

	var user model.User
	err = db.Get(&user, sql, email)
	if err != nil {
		p.Logger.Println(err)
		return 0, errors.New("query error")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		p.Logger.Println(err)
		return 0, errors.New("wrong username or password")
	}

	return user.Id, nil
}

func (p Postgres) Register(user model.User) (uint, error) {
	db, err := p.Connect()
	if err != nil {
		p.Logger.Println(err)
		return 0, err
	}
	defer db.Close()

	sql := "INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING id"

	var id uint

	tx := db.MustBegin()
	err = tx.QueryRowx(sql, user.Email, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, errors.New("user with this email already exists")
	}

	if tx.Commit() != nil {
		return 0, errors.New("transaction error")
	}

	return id, nil
}

func (p Postgres) Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", p.ConnString())
	if err != nil {
		return nil, errors.New("cannot connect to db")
	}

	return db, nil
}
