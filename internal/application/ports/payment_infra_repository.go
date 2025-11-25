package ports

import (
	"context"
	"graphql-payment-bff/internal/domain/model"
)

// PaymentInfraRepository defines the repository interface for payment infrastructure data
type PaymentInfraRepository interface {
	GetPaymentInfraByID(ctx context.Context, paymentRackID string) (*model.PaymentInfra, error)
	GetAvailableLockers(ctx context.Context, paymentRackID int, bookingTimeID int) (*model.AvailableLockers, error)
	ValidateDiscountCoupon(ctx context.Context, couponCode string) (*model.DiscountCouponValidation, error)
	GeneratePurchaseOrder(ctx context.Context, groupID int, couponCode *string, userEmail string, userPhone string) (*model.PurchaseOrder, error)
}
