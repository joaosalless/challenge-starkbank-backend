package main

import (
	"joaosalless/challenge-starkbank/bootstrap"
	"joaosalless/challenge-starkbank/src/http"
	"log"
)

func main() {
	container := bootstrap.Initialize()

	err := container.Invoke(http.NewServer)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
