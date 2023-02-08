package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	URL := os.Getenv("CHUCK_NORRIS_JOKES")
	if URL == "" {
		log.Fatal("get .env failed")
	}

	svc := NewChuckNorrisJokesService(URL)
	svc = NewLoggingMiddleware(svc)

	server := NewApiServer(svc)

	http.HandleFunc("/", server.GetRandomChuckNorrisJokesHandler)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("cannot connect to :3000")
	}
}
