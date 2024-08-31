package main

import (
	"joaosalless/challenge-starkbank/cmd"
	"joaosalless/challenge-starkbank/src/schedule"
	"log"
)

func main() {
	container := cmd.Initialize()

	err := container.Invoke(schedule.NewScheduledTasks)
	if err != nil {
		log.Fatal("Failed to start scheduler:", err)
	}
}
