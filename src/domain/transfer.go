package domain

import "time"

type Transfer struct {
	Id                 string                 `json:",omitempty"`
	Amount             int                    `json:",omitempty"`
	Name               string                 `json:",omitempty"`
	TaxId              string                 `json:",omitempty"`
	BankCode           string                 `json:",omitempty"`
	BranchCode         string                 `json:",omitempty"`
	AccountNumber      string                 `json:",omitempty"`
	AccountType        string                 `json:",omitempty"`
	ExternalId         string                 `json:",omitempty"`
	Scheduled          *time.Time             `json:",omitempty"`
	Description        string                 `json:",omitempty"`
	DisplayDescription string                 `json:",omitempty"`
	Tags               []string               `json:",omitempty"`
	Rules              []RuleInt              `json:",omitempty"`
	Fee                int                    `json:",omitempty"`
	Status             string                 `json:",omitempty"`
	TransactionIds     []string               `json:",omitempty"`
	Metadata           map[string]interface{} `json:",omitempty"`
	Created            *time.Time             `json:",omitempty"`
	Updated            *time.Time             `json:",omitempty"`
}
