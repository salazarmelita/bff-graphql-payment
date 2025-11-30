package ports

import (
	"context"
	"graphql-payment-bff/internal/domain/model"
)

// PaymentInfraRepository define la interfaz del repositorio para datos de infraestructura de pagos
type PaymentInfraRepository interface {
	GetPaymentInfraByQrValue(ctx context.Context, qrValue string) (*model.PaymentInfra, error)
	GetAvailableLockers(ctx context.Context, paymentRackID int, bookingTimeID int) (*model.AvailableLockers, error)
	ValidateDiscountCoupon(ctx context.Context, couponCode string) (*model.DiscountCouponValidation, error)
	GeneratePurchaseOrder(ctx context.Context, groupID int, couponCode *string, userEmail string, userPhone string) (*model.PurchaseOrder, error)
}
