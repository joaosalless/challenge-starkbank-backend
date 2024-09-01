package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	App          App
	Api          Api
	Clock        Clock
	Scheduler    Scheduler
	Invoice      Invoice
	Transfer     Transfer
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

type Invoice struct {
	ExpirationDays          int
	RandomInvoicesNumberMin int
	RandomInvoicesNumberMax int
}

type Transfer struct {
	BankAccount BankAccount
}

type BankAccount struct {
	Name          string
	TaxId         string
	BankCode      string
	BranchCode    string
	AccountNumber string
	AccountType   string
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
		Invoice: Invoice{
			ExpirationDays:          getEnvInt("INVOICE_EXPIRATION_DAYS", 1),
			RandomInvoicesNumberMin: getEnvInt("INVOICE_RANDOM_INVOICES_NUMBER_MIN", 8),
			RandomInvoicesNumberMax: getEnvInt("INVOICE_RANDOM_INVOICES_NUMBER_MAX", 12),
		},
		Transfer: Transfer{
			BankAccount: BankAccount{
				Name:          getEnv("BANK_ACCOUNT_NAME", ""),
				TaxId:         getEnv("BANK_ACCOUNT_TAX_ID", ""),
				BankCode:      getEnv("BANK_ACCOUNT_BANK_CODE", ""),
				BranchCode:    getEnv("BANK_ACCOUNT_BRANCH_CODE", ""),
				AccountNumber: getEnv("BANK_ACCOUNT_ACCOUNT_NUMBER", ""),
				AccountType:   getEnv("BANK_ACCOUNT_ACCOUNT_TYPE", ""),
			},
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

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(fmt.Printf("failed to convert env value to int: %v\n", err))
		}
		return intValue
	}
	return defaultValue
}
