package service

import (
	"context"
	"graphql-payment-bff/internal/application/ports"
	"graphql-payment-bff/internal/domain/exception"
	"graphql-payment-bff/internal/domain/model"
	"strings"
)

// PaymentInfraService implements the payment infrastructure use cases
type PaymentInfraService struct {
	repo ports.PaymentInfraRepository
}

// NewPaymentInfraService creates a new payment infrastructure service
func NewPaymentInfraService(repo ports.PaymentInfraRepository) *PaymentInfraService {
	return &PaymentInfraService{
		repo: repo,
	}
}

// GetPaymentInfraByID retrieves payment infrastructure by ID
func (s *PaymentInfraService) GetPaymentInfraByID(ctx context.Context, paymentRackID string) (*model.PaymentInfra, error) {
	// Validate input
	if strings.TrimSpace(paymentRackID) == "" {
		return nil, exception.ErrInvalidPaymentRackID
	}

	// Call repository
	paymentInfra, err := s.repo.GetPaymentInfraByID(ctx, paymentRackID)
	if err != nil {
		return nil, err
	}

	return paymentInfra, nil
}
