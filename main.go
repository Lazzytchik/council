package main

import (
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/lazzytchik/council/internal/app"
	"github.com/lazzytchik/council/internal/data"
	"github.com/lazzytchik/council/internal/session"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Printf("Can't load .env file: %s", envErr)
	}

	host := os.Getenv("HOST")
	addr := os.Getenv("PORT")

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

	el.Println(options.Describe())

	postgres := data.NewPostgres(options, el)
	ses := session.NewRedis(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT_NUMBER"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	},
		time.Hour*4,
	)

	builder.ConfigureIdentifier(postgres)
	builder.ConfigureRegistrar(postgres)
	builder.ConfigureSession(ses)
	builder.ConfigureServer(&http.Server{
		Addr:     ":" + addr,
		ErrorLog: el,
	})

	s := builder.Build()

	log.Printf("Server is running on %s:%s.", host, addr)
	if err := s.ListenAndServe(); err != nil {
		log.Println("cannot start server:", err)
	}

}
