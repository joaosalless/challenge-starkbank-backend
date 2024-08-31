package config

import (
	"os"
)

type Config struct {
	App          App
	Api          Api
	Clock        Clock
	Scheduler    Scheduler
	BankProvider BankProvider
}

type App struct {
	Env     string
	Name    string
	Version string
}

type Api struct {
	Name    string
	Port    string
	Spec    string
	Version string
}

type Clock struct {
	Location string
}

type Scheduler struct {
	Enabled                    bool
	InvoiceCreateScheduledTime string
}

type BankProvider struct {
	Starkbank Starkbank
}

type Starkbank struct {
	ProjectId   string
	PublicKey   string
	PrivateKey  string
	Environment string
}

func LoadConfig() *Config {
	return &Config{
		App: App{
			Env:     getEnv("APP_ENV", "local"),
			Name:    getEnv("APP_NAME", ""),
			Version: getEnv("APP_VERSION", ""),
		},
		Api: Api{
			Name:    getEnv("API_NAME", ""),
			Port:    getEnv("API_PORT", "8080"),
			Spec:    getEnv("API_SPEC", ""),
			Version: getEnv("API_VERSION", ""),
		},
		Clock: Clock{
			Location: getEnv("CLOCK_LOCATION", "America/Sao_Paulo"),
		},
		BankProvider: BankProvider{
			Starkbank: Starkbank{
				ProjectId:   getEnv("STARKBANK_PROJECT_ID", ""),
				PublicKey:   getEnv("STARKBANK_PUBLIC_KEY", ""),
				PrivateKey:  getEnv("STARKBANK_PRIVATE_KEY", ""),
				Environment: getEnv("STARKBANK_ENVIRONMENT", ""),
			},
		},
		Scheduler: Scheduler{
			Enabled:                    getEnv("SCHEDULER_ENABLED", "false") == "true",
			InvoiceCreateScheduledTime: getEnv("INVOICE_CREATE_SCHEDULED_TIME", ""),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
