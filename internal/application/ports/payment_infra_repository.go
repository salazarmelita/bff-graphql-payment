package ports

import (
	"context"
	"graphql-payment-bff/internal/domain/model"
)

// PaymentInfraRepository defines the repository interface for payment infrastructure data
type PaymentInfraRepository interface {
	GetPaymentInfraByID(ctx context.Context, paymentRackID string) (*model.PaymentInfra, error)
}
