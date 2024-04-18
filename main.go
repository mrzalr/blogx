package main

import (
	"blogx/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	s := server.New()

	port := os.Getenv("PORT")
	if len(port) > 0 {
		s.WithPort(port)
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
