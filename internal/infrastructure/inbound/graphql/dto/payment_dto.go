package dto

// GraphQL DTOs for input/output transformation

// GetPaymentInfraByIDInput represents the GraphQL input for getting payment infrastructure
type GetPaymentInfraByIDInput struct {
	PaymentRackID string `json:"paymentRackId"`
}

// PaymentInfraResponse represents the GraphQL response for payment infrastructure
type PaymentInfraResponse struct {
	TransactionID string                 `json:"transactionId"`
	Message       string                 `json:"message"`
	Status        string                 `json:"status"`
	PaymentRack   *PaymentRackResponse   `json:"paymentRack,omitempty"`
	Installation  *InstallationResponse  `json:"installation,omitempty"`
	BookingTimes  []*BookingTimeResponse `json:"bookingTimes"`
}

// PaymentRackResponse represents a payment rack in GraphQL response
type PaymentRackResponse struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

// InstallationResponse represents an installation in GraphQL response
type InstallationResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Region   string `json:"region"`
	City     string `json:"city"`
	Address  string `json:"address"`
	ImageURL string `json:"imageUrl"`
}

// BookingTimeResponse represents booking time in GraphQL response
type BookingTimeResponse struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	UnitMeasurement string `json:"unitMeasurement"`
	Amount          int    `json:"amount"`
}
