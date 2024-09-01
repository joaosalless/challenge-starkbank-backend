package main

import (
	"github.com/joaosalless/challenge-starkbank-backend/cmd"
	"github.com/joaosalless/challenge-starkbank-backend/src/api/http"
	"log"
)

func main() {
	container := cmd.Initialize()

	err := container.Invoke(http.NewServer)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
