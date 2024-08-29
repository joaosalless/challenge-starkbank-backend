package config

import (
	"os"
)

type Config struct {
	AppPort                   string
	StarkBankPublicKey        string
	StarkBankPrivateKey       string
	InvoiceCreateCronInterval string
}

func LoadConfig() *Config {
	return &Config{
		AppPort:                   getEnv("APP_PORT", "8080"),
		StarkBankPublicKey:        getEnv("STARK_BANK_PUBLIC_KEY", ""),
		StarkBankPrivateKey:       getEnv("STARK_BANK_PRIVATE_KEY", ""),
		InvoiceCreateCronInterval: getEnv("INVOICE_CREATE_CRON_INTERVAL", ""),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
