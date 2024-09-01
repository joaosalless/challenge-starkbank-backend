package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	"github.com/joaosalless/challenge-starkbank-backend/config"
	"github.com/joaosalless/challenge-starkbank-backend/src/domain"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/joaosalless/challenge-starkbank-backend/src/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransferService_calculateTransferAmount(t *testing.T) {
	type args struct {
		invoice domain.Invoice
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sem desconto e sem taxa",
			args: args{
				invoice: domain.Invoice{
					NominalAmount:  10000,
					DiscountAmount: 0,
					Fee:            0,
				},
			},
			want: 10000,
		},
		{
			name: "Com desconto e sem taxa",
			args: args{
				invoice: domain.Invoice{
					NominalAmount:  10000,
					DiscountAmount: 1000,
					Fee:            0,
				},
			},
			want: 9000,
		},
		{
			name: "Sem desconto e com taxa",
			args: args{
				invoice: domain.Invoice{
					NominalAmount:  10000,
					DiscountAmount: 0,
					Fee:            500,
				},
			},
			want: 9500,
		},
		{
			name: "Com desconto e com taxa",
			args: args{
				invoice: domain.Invoice{
					NominalAmount:  10000,
					DiscountAmount: 1000,
					Fee:            500,
				},
			},
			want: 8500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := TransferService{}
			if got := i.calculateTransferAmount(tt.args.invoice); got != tt.want {
				t.Errorf("calculateTransferAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransferService_CreateTransferFromInvoice(t *testing.T) {
	type deps struct {
		logger         *mocks.MockLogger
		bankGateway    *mocks.MockBankGateway
		transferConfig config.Transfer
	}

	type args struct {
		ctx   context.Context
		input dtos.CreateTransferFromInvoiceInput
	}

	type want struct {
		out dtos.CreateTransferOutput
		err error
	}

	type setup struct {
		deps deps
		args args
		want want
	}

	tests := []struct {
		name  string
		setup func(ctrl *gomock.Controller, ctx context.Context, deps deps, invoice domain.Invoice) setup
	}{
		{
			name: "should fail when invoice status is not paid",
			setup: func(ctrl *gomock.Controller, ctx context.Context, deps deps, invoice domain.Invoice) setup {
				invoice.Status = domain.InvoiceStatusCreated

				deps.logger.EXPECT().Infow(gomock.Any(), gomock.Any()).Times(1)
				deps.logger.EXPECT().Errorw(gomock.Any(), gomock.Any()).Times(1)

				return setup{
					deps: deps,
					args: args{ctx: ctx, input: dtos.CreateTransferFromInvoiceInput{Data: invoice}},
					want: want{
						out: dtos.CreateTransferOutput{},
						err: errors.New("invalid invoice status"),
					},
				}
			},
		},
		{
			name: "should fail when bank gateway returns error",
			setup: func(ctrl *gomock.Controller, ctx context.Context, deps deps, invoice domain.Invoice) setup {
				invoice.Status = domain.InvoiceStatusPaid
				invoice.NominalAmount = 10000
				invoice.DiscountAmount = 0
				invoice.Fee = 0
				invoice.ExternalId = "fake-external-id"
				invoice.DisplayDescription = "Fake invoice description"

				deps.logger.EXPECT().
					Infow(gomock.Any(), gomock.Any()).
					AnyTimes()

				deps.bankGateway.EXPECT().
					CreateTransfer(gomock.Any(), dtos.CreateTransferInput{Data: []domain.Transfer{
						{
							Amount:        TransferService{}.calculateTransferAmount(invoice),
							Name:          deps.transferConfig.BankAccount.Name,
							TaxId:         deps.transferConfig.BankAccount.TaxId,
							BankCode:      deps.transferConfig.BankAccount.BankCode,
							BranchCode:    deps.transferConfig.BankAccount.BranchCode,
							AccountNumber: deps.transferConfig.BankAccount.AccountNumber,
							AccountType:   deps.transferConfig.BankAccount.AccountType,
							Tags: []string{
								fmt.Sprintf("urn:invoice:%s", invoice.ExternalId),
							},
							Description: fmt.Sprintf("Payment for invoice #%s - %s", invoice.ExternalId, invoice.DisplayDescription),
						},
					}}).
					Times(1).
					Return(dtos.CreateTransferOutput{}, assert.AnError)

				deps.logger.EXPECT().Errorw(gomock.Any(), gomock.Any()).AnyTimes()

				return setup{
					deps: deps,
					args: args{ctx: ctx, input: dtos.CreateTransferFromInvoiceInput{Data: invoice}},
					want: want{
						out: dtos.CreateTransferOutput{},
						err: fmt.Errorf("error when call bankGateway.CreateTransfer: %w", assert.AnError),
					},
				}
			},
		},
		{
			name: "should success when bank gateway returns success",
			setup: func(ctrl *gomock.Controller, ctx context.Context, deps deps, invoice domain.Invoice) setup {
				invoice.Status = domain.InvoiceStatusPaid
				invoice.NominalAmount = 10000
				invoice.DiscountAmount = 0
				invoice.Fee = 0
				invoice.ExternalId = "fake-external-id"
				invoice.DisplayDescription = "Fake invoice description"

				deps.logger.EXPECT().
					Infow(gomock.Any(), gomock.Any()).
					AnyTimes()

				deps.bankGateway.EXPECT().
					CreateTransfer(gomock.Any(), dtos.CreateTransferInput{Data: []domain.Transfer{
						{
							Amount:        TransferService{}.calculateTransferAmount(invoice),
							Name:          deps.transferConfig.BankAccount.Name,
							TaxId:         deps.transferConfig.BankAccount.TaxId,
							BankCode:      deps.transferConfig.BankAccount.BankCode,
							BranchCode:    deps.transferConfig.BankAccount.BranchCode,
							AccountNumber: deps.transferConfig.BankAccount.AccountNumber,
							AccountType:   deps.transferConfig.BankAccount.AccountType,
							Tags: []string{
								fmt.Sprintf("urn:invoice:%s", invoice.ExternalId),
							},
							Description: fmt.Sprintf("Payment for invoice #%s - %s", invoice.ExternalId, invoice.DisplayDescription),
						},
					}}).
					Times(1).
					Return(dtos.CreateTransferOutput{}, nil)

				return setup{
					deps: deps,
					args: args{ctx: ctx, input: dtos.CreateTransferFromInvoiceInput{Data: invoice}},
					want: want{
						out: dtos.CreateTransferOutput{},
						err: nil,
					},
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			invoice := domain.Invoice{}

			setup := tt.setup(
				ctrl,
				context.Background(),
				deps{
					logger:      mocks.NewMockLogger(ctrl),
					bankGateway: mocks.NewMockBankGateway(ctrl),
					transferConfig: config.Transfer{
						BankAccount: config.BankAccount{
							Name:          fmt.Sprintf("%s %s", gofakeit.Person().FirstName, gofakeit.Person().LastName),
							TaxId:         "99.999.999/9999-99",
							BankCode:      "99999999",
							BranchCode:    "9999",
							AccountNumber: "9999999999999999",
							AccountType:   "payment",
						},
					},
				},
				invoice,
			)

			s := TransferService{
				logger:         setup.deps.logger,
				bankGateway:    setup.deps.bankGateway,
				transferConfig: setup.deps.transferConfig,
			}

			out, err := s.CreateTransferFromInvoice(setup.args.ctx, setup.args.input)
			assert.Equal(t, setup.want.err, err)
			assert.EqualValues(t, setup.want.out, out)
		})
	}
}
