package main

import (
	"github.com/joaosalless/challenge-starkbank-backend/cmd"
	"github.com/joaosalless/challenge-starkbank-backend/src/schedule"
	"log"
)

func main() {
	container := cmd.Initialize()

	err := container.Invoke(schedule.NewScheduledTasks)
	if err != nil {
		log.Fatal("Failed to start scheduler:", err)
	}
}
