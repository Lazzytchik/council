package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"lazzytchk/council/internal/app"
	"lazzytchk/council/internal/data"
	"log"
	"net/http"
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

	builder := app.ServerBuilder{
		Server: &app.Server{},
	}

	postgres := data.NewPostgres(options, el)

	builder.ConfigureIdentifier(postgres)
	builder.ConfigureRegistrar(postgres)
	builder.ConfigureServer(&http.Server{
		Addr:     ":8080",
		ErrorLog: el,
	})

	s := builder.Build()
	s.ListenAndServe()
}
