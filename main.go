package main

import (
	"lazzytchk/council/internal/http"
	"log"
	"os"
)

func main() {
	el := log.New(os.Stdout, "[SERVER]: ", log.Lmicroseconds)

	s := http.NewServer(":8080", el)
	s.ListenAndServe()
}
