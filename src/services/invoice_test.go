package services

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/joaosalless/challenge-starkbank-backend/src/domain"
	"github.com/mvrilo/go-cpf"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/joaosalless/challenge-starkbank-backend/src/dtos"
	"github.com/joaosalless/challenge-starkbank-backend/src/mocks"
	"github.com/stretchr/testify/assert"
)

func TestInvoiceService_CreateInvoice(t *testing.T) {
	type deps struct {
		logger      *mocks.MockLogger
		bankGateway *mocks.MockBankGateway
	}

	type args struct {
		ctx   context.Context
		input dtos.CreateInvoiceInput
	}

	type want struct {
		out dtos.CreateInvoiceOutput
		err error
	}

	type setup struct {
		deps deps
		args args
		want want
	}

	tests := []struct {
		name  string
		setup func(ctrl *gomock.Controller, ctx context.Context, deps deps, invoice dtos.CreateInvoiceInput) setup
	}{
		{
			name: "should succeed when invoice is created successfully",
			setup: func(ctrl *gomock.Controller, ctx context.Context, deps deps, invoice dtos.CreateInvoiceInput) setup {
				output := dtos.CreateInvoiceOutput{
					Data: invoice.Data,
				}

				deps.logger.EXPECT().
					Infow("InvoiceService.CreateInvoice", "input", gomock.Any()).
					Times(1)

				deps.bankGateway.EXPECT().
					CreateInvoice(gomock.Any(), invoice).
					Return(output, nil).
					Times(1)

				return setup{
					deps: deps,
					args: args{ctx: ctx, input: invoice},
					want: want{
						out: output,
						err: nil,
					},
				}
			},
		},
		{
			name: "should fail when bank gateway returns error",
			setup: func(ctrl *gomock.Controller, ctx context.Context, deps deps, invoice dtos.CreateInvoiceInput) setup {
				deps.logger.EXPECT().Infow("InvoiceService.CreateInvoice", "input", gomock.Any()).Times(1)

				deps.bankGateway.EXPECT().
					CreateInvoice(gomock.Any(), invoice).
					Return(dtos.CreateInvoiceOutput{}, assert.AnError).
					Times(1)

				return setup{
					deps: deps,
					args: args{ctx: ctx, input: invoice},
					want: want{
						out: dtos.CreateInvoiceOutput{},
						err: assert.AnError,
					},
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			due := gofakeit.DateRange(
				time.Now().AddDate(0, 0, 3),
				time.Now().AddDate(0, 0, 60))

			invoice := dtos.CreateInvoiceInput{
				Data: []domain.Invoice{
					{
						Amount:     gofakeit.Number(10000, 20000),
						Due:        &due,
						Expiration: 1,
						Name:       fmt.Sprintf("%s %s", gofakeit.Person().FirstName, gofakeit.Person().LastName),
						TaxId:      cpf.GeneratePretty(),
					},
				},
			}

			setup := tt.setup(
				ctrl,
				context.Background(),
				deps{
					logger:      mocks.NewMockLogger(ctrl),
					bankGateway: mocks.NewMockBankGateway(ctrl),
				},
				invoice,
			)

			s := InvoiceService{
				logger:      setup.deps.logger,
				bankGateway: setup.deps.bankGateway,
			}

			out, err := s.CreateInvoice(setup.args.ctx, setup.args.input)
			assert.Equal(t, setup.want.err, err)
			assert.EqualValues(t, setup.want.out, out)
		})
	}
}
