package main

import (
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/lazzytchik/council/grpc/auth"
	"github.com/lazzytchik/council/internal/data"
	"github.com/lazzytchik/council/internal/session"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Printf("Can't load .env file: %s", envErr)
	}

	host := os.Getenv("HOST")
	addr := ":" + os.Getenv("PORT")

	options := data.ConnOptions{
		Name:     os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	el := log.New(os.Stdout, "[gRPC-SERVER]: ", log.Lmicroseconds)

	el.Println(options)

	builder := auth.ServerBuilder{}

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
	builder.ConfigureLogger(el)

	api := builder.Build()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", host+addr, err)
	}

	s := grpc.NewServer()

	auth.RegisterUserStorageServer(s, api)
	log.Printf("Starting gRPC server on %s%s", host, addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
