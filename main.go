package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"lazzytchk/council/internal/data"
	"lazzytchk/council/internal/http"
	"log"
	"os"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Can't load .env file: %s", envErr)
	}

	options := data.ConnOptions{
		Name:     os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	el := log.New(os.Stdout, "[SERVER]: ", log.Lmicroseconds)

	s := http.NewServer(":8080", el, options)
	s.ListenAndServe()
}
