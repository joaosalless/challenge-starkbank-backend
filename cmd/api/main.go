package main

import (
	"joaosalless/challenge-starkbank/cmd"
	"joaosalless/challenge-starkbank/src/api/http"
	"log"
)

func main() {
	container := cmd.Initialize()

	err := container.Invoke(http.NewServer)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
