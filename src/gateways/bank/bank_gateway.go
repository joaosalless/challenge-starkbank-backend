package bank

import (
	"context"
	"fmt"
	"github.com/joaosalless/challenge-starkbank-backend/pkg/app"
	"github.com/joaosalless/challenge-starkbank-backend/src/domain"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	StarkInvoice "github.com/starkbank/sdk-go/starkbank/invoice"
	StarkInvoiceRule "github.com/starkbank/sdk-go/starkbank/invoice/rule"
	StarkTransfer "github.com/starkbank/sdk-go/starkbank/transfer"
	StarkTransferRule "github.com/starkbank/sdk-go/starkbank/transfer/rule"
	"github.com/starkinfra/core-go/starkcore/user/project"
	"github.com/starkinfra/core-go/starkcore/user/user"
)

// BankGateway implements interfaces.BankGateway
type BankGateway struct {
	user user.User
}

type Dependencies struct {
	app.Dependencies
}

func NewBankGateway(deps Dependencies) *BankGateway {
	return &BankGateway{
		user: project.Project{
			Id:          deps.Config.BankProvider.Starkbank.ProjectId,
			Environment: deps.Config.BankProvider.Starkbank.Environment,
			PrivateKey:  deps.Config.BankProvider.Starkbank.PrivateKey,
		},
	}
}

func (g BankGateway) CreateInvoice(ctx context.Context, input dtos.CreateInvoiceInput) (output dtos.CreateInvoiceOutput, err error) {
	var payload []StarkInvoice.Invoice

	for _, item := range input.Data {
		var rules []StarkInvoiceRule.Rule

		for _, r := range item.Rules {
			rules = append(rules, StarkInvoiceRule.Rule{Key: r.Key, Value: r.Value})
		}

		payload = append(payload, StarkInvoice.Invoice{
			Amount:       item.Amount,
			Name:         item.Name,
			TaxId:        item.TaxId,
			Due:          item.Due,
			Expiration:   item.Expiration,
			Fine:         item.Fine,
			Interest:     item.Interest,
			Discounts:    item.Discounts,
			Tags:         item.Tags,
			Rules:        rules,
			Descriptions: item.Descriptions,
		})
	}

	invoices, stackErr := StarkInvoice.Create(payload, g.user)

	if stackErr.Errors != nil {
		return output, fmt.Errorf("error creating invoice: %+v", stackErr)
	}

	for _, item := range invoices {
		var rules []domain.RuleStrings

		for _, r := range item.Rules {
			rules = append(rules, domain.RuleStrings{Key: r.Key, Value: r.Value})
		}

		output.Data = append(output.Data, domain.Invoice{
			Id:                 item.Id,
			Amount:             item.Amount,
			Name:               item.Name,
			TaxId:              item.TaxId,
			Due:                item.Due,
			Expiration:         item.Expiration,
			Fine:               item.Fine,
			Interest:           item.Interest,
			Discounts:          item.Discounts,
			Tags:               item.Tags,
			Descriptions:       item.Descriptions,
			DisplayDescription: item.DisplayDescription,
			Pdf:                item.Pdf,
			Link:               item.Link,
			NominalAmount:      item.NominalAmount,
			FineAmount:         item.FineAmount,
			InterestAmount:     item.InterestAmount,
			DiscountAmount:     item.DiscountAmount,
			Brcode:             item.Brcode,
			Status:             item.Status,
			Fee:                item.Fee,
			TransactionIds:     item.TransactionIds,
			Created:            item.Created,
			Updated:            item.Updated,
		})
	}

	return output, err
}

func (g BankGateway) CreateTransfer(ctx context.Context, input dtos.CreateTransferInput) (output dtos.CreateTransferOutput, err error) {
	var payload []StarkTransfer.Transfer

	for _, item := range input.Data {
		var rules []StarkTransferRule.Rule

		for _, r := range item.Rules {
			rules = append(rules, StarkTransferRule.Rule{Key: r.Key, Value: r.Value})
		}

		payload = append(payload, StarkTransfer.Transfer{
			Amount:             item.Amount,
			Name:               item.Name,
			TaxId:              item.TaxId,
			BankCode:           item.BankCode,
			BranchCode:         item.BranchCode,
			AccountNumber:      item.AccountNumber,
			AccountType:        item.AccountType,
			ExternalId:         item.ExternalId,
			Scheduled:          item.Scheduled,
			Description:        item.Description,
			DisplayDescription: item.DisplayDescription,
			Tags:               item.Tags,
			Rules:              rules,
			Fee:                item.Fee,
		})
	}

	transfers, stackErr := StarkTransfer.Create(payload, g.user)

	if stackErr.Errors != nil {
		return output, fmt.Errorf("error creating transfers: %+v", stackErr)
	}

	for _, item := range transfers {
		var rules []domain.RuleInt
		for _, r := range item.Rules {
			rules = append(rules, domain.RuleInt{Key: r.Key, Value: r.Value})
		}

		output.Data = append(output.Data, domain.Transfer{
			Id:                 item.Id,
			Amount:             item.Amount,
			Name:               item.Name,
			TaxId:              item.TaxId,
			BankCode:           item.BankCode,
			BranchCode:         item.BranchCode,
			AccountNumber:      item.AccountNumber,
			AccountType:        item.AccountType,
			ExternalId:         item.ExternalId,
			Scheduled:          item.Scheduled,
			Description:        item.Description,
			DisplayDescription: item.DisplayDescription,
			Tags:               item.Tags,
			Rules:              rules,
			Fee:                item.Fee,
			Status:             item.Status,
			TransactionIds:     item.TransactionIds,
			Metadata:           item.Metadata,
			Created:            item.Created,
			Updated:            item.Updated,
		})
	}

	return output, err
}
