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

// GetAvailableLockers retrieves available lockers by rack ID and booking time
func (s *PaymentInfraService) GetAvailableLockers(ctx context.Context, paymentRackID int, bookingTimeID int) (*model.AvailableLockers, error) {
	// Validate input
	if paymentRackID <= 0 {
		return nil, exception.ErrInvalidPaymentRackID
	}

	if bookingTimeID <= 0 {
		return nil, exception.ErrInvalidBookingTimeID
	}

	// Call repository
	lockers, err := s.repo.GetAvailableLockers(ctx, paymentRackID, bookingTimeID)
	if err != nil {
		return nil, err
	}

	return lockers, nil
}

// ValidateDiscountCoupon validates a discount coupon
func (s *PaymentInfraService) ValidateDiscountCoupon(ctx context.Context, couponCode string) (*model.DiscountCouponValidation, error) {
	// Validate input
	if strings.TrimSpace(couponCode) == "" {
		return nil, exception.ErrInvalidCouponCode
	}

	// Call repository
	validation, err := s.repo.ValidateDiscountCoupon(ctx, couponCode)
	if err != nil {
		return nil, err
	}

	return validation, nil
}

// GeneratePurchaseOrder generates a purchase order
func (s *PaymentInfraService) GeneratePurchaseOrder(ctx context.Context, groupID int, couponCode *string, userEmail string, userPhone string) (*model.PurchaseOrder, error) {
	// Validate input
	if groupID <= 0 {
		return nil, exception.ErrInvalidGroupID
	}

	if strings.TrimSpace(userEmail) == "" {
		return nil, exception.ErrInvalidEmail
	}

	if strings.TrimSpace(userPhone) == "" {
		return nil, exception.ErrInvalidPhone
	}

	// Call repository
	order, err := s.repo.GeneratePurchaseOrder(ctx, groupID, couponCode, userEmail, userPhone)
	if err != nil {
		return nil, err
	}

	return order, nil
}
