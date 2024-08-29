package main

import (
	"joaosalless/challenge-starkbank/bootstrap"
	"joaosalless/challenge-starkbank/src/cron"
	"log"
)

func main() {
	container := bootstrap.Initialize()

	err := container.Invoke(cron.NewInvoiceCreateCron)
	if err != nil {
		log.Fatal("Failed to run invoice create cron:", err)
	}
}
