package model

// PaymentInfra represents the aggregate of payment infrastructure data
type PaymentInfra struct {
	TransactionID string
	Message       string
	Status        ResponseStatus
	PaymentRack   *PaymentRack
	Installation  *PaymentInstallation
	BookingTimes  []PaymentBookingTime
}

// PaymentRack represents a payment rack entity
type PaymentRack struct {
	ID          int
	Description string
	Address     string
}

// PaymentInstallation represents an installation entity
type PaymentInstallation struct {
	ID       int
	Name     string
	Region   string
	City     string
	Address  string
	ImageURL string
}

// PaymentBookingTime represents booking time configuration
type PaymentBookingTime struct {
	ID              int
	Name            string
	UnitMeasurement UnitMeasurement
	Amount          int
}

// ResponseStatus enum
type ResponseStatus string

const (
	ResponseStatusUnspecified ResponseStatus = "RESPONSE_STATUS_UNSPECIFIED"
	ResponseStatusOK          ResponseStatus = "RESPONSE_STATUS_OK"
	ResponseStatusError       ResponseStatus = "RESPONSE_STATUS_ERROR"
)

// UnitMeasurement enum
type UnitMeasurement string

const (
	UnitMeasurementUnspecified UnitMeasurement = "UNSPECIFIED"
	UnitMeasurementHour        UnitMeasurement = "HOUR"
	UnitMeasurementDay         UnitMeasurement = "DAY"
	UnitMeasurementWeek        UnitMeasurement = "WEEK"
	UnitMeasurementMonth       UnitMeasurement = "MONTH"
)
