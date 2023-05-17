package main

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"lazzytchk/council/grpc/auth"
	"log"
	"net"
	"os"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Can't load .env file: %s", envErr)
	}

	host := os.Getenv("HOST")
	addr := ":" + os.Getenv("PORT")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %s", addr, err)
	}

	s := grpc.NewServer()

	api := &auth.Server{}

	auth.RegisterUserStorageServer(s, api)
	log.Printf("Starting gRPC server on %s%s", host, addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
