package dto

// gRPC DTOs based on the protobuf definitions

// PaymentManagerGenericResponse represents the generic response wrapper
type PaymentManagerGenericResponse struct {
	TransactionId string                       `json:"transaction_id"`
	Message       string                       `json:"message"`
	Status        PaymentManagerResponseStatus `json:"status"`
}

// PaymentManagerResponseStatus enum
type PaymentManagerResponseStatus int32

const (
	PaymentManagerResponseStatus_RESPONSE_STATUS_UNSPECIFIED PaymentManagerResponseStatus = 0
	PaymentManagerResponseStatus_RESPONSE_STATUS_OK          PaymentManagerResponseStatus = 1
	PaymentManagerResponseStatus_RESPONSE_STATUS_ERROR       PaymentManagerResponseStatus = 2
)

// GetPaymentInfraByIDRequest represents the request for getting payment infra by ID
type GetPaymentInfraByIDRequest struct {
	PaymentRackId string `json:"payment_rack_id"`
}

// GetPaymentInfraByIDResponse represents the response for getting payment infra by ID
type GetPaymentInfraByIDResponse struct {
	Response     *PaymentManagerGenericResponse `json:"response"`
	PaymentRack  *PaymentRackRecord             `json:"payment_rack"`
	Installation *PaymentInstallationRecord     `json:"installation"`
	BookingTimes []*PaymentBookingTimeRecord    `json:"booking_times"`
}

// PaymentInstallationRecord represents installation data
type PaymentInstallationRecord struct {
	Id       int32  `json:"id"`
	Name     string `json:"name"`
	Region   string `json:"region"`
	City     string `json:"city"`
	Address  string `json:"address"`
	ImageUrl string `json:"image_url"`
}

// PaymentRackRecord represents payment rack data
type PaymentRackRecord struct {
	Id          int32  `json:"id"`
	Description string `json:"description"`
	Address     string `json:"address"`
}

// PaymentBookingTimeRecord represents booking time configuration
type PaymentBookingTimeRecord struct {
	Id              int32           `json:"id"`
	Name            string          `json:"name"`
	UnitMeasurement UnitMeasurement `json:"unit_measurement"`
	Amount          int32           `json:"amount"`
}

// UnitMeasurement enum
type UnitMeasurement int32

const (
	UnitMeasurement_UNSPECIFIED UnitMeasurement = 0
	UnitMeasurement_HOUR        UnitMeasurement = 1
	UnitMeasurement_DAY         UnitMeasurement = 2
	UnitMeasurement_WEEK        UnitMeasurement = 3
	UnitMeasurement_MONTH       UnitMeasurement = 4
)
