package ports

import (
	"context"
	"graphql-payment-bff/internal/domain/model"
)

// PaymentInfraService defines the use case for payment infrastructure operations
type PaymentInfraService interface {
	GetPaymentInfraByID(ctx context.Context, paymentRackID string) (*model.PaymentInfra, error)
}
