package model

// ResponseStatus enum
type ResponseStatus string

const (
	ResponseStatusResponseStatusUnspecified ResponseStatus = "RESPONSE_STATUS_UNSPECIFIED"
	ResponseStatusResponseStatusOk          ResponseStatus = "RESPONSE_STATUS_OK"
	ResponseStatusResponseStatusError       ResponseStatus = "RESPONSE_STATUS_ERROR"
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
