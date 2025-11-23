package exception

import "errors"

var (
	// ErrPaymentRackNotFound is returned when a payment rack is not found
	ErrPaymentRackNotFound = errors.New("payment rack not found")

	// ErrInvalidPaymentRackID is returned when the payment rack ID is invalid
	ErrInvalidPaymentRackID = errors.New("invalid payment rack ID")

	// ErrPaymentInfraServiceUnavailable is returned when the payment infrastructure service is unavailable
	ErrPaymentInfraServiceUnavailable = errors.New("payment infrastructure service unavailable")
)
