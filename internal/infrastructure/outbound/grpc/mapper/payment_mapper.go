package mapper

import (
	"graphql-payment-bff/internal/domain/model"
	"graphql-payment-bff/internal/infrastructure/outbound/grpc/dto"
)

// PaymentInfraGRPCMapper handles mapping between domain models and gRPC DTOs
type PaymentInfraGRPCMapper struct{}

// NewPaymentInfraGRPCMapper creates a new mapper instance
func NewPaymentInfraGRPCMapper() *PaymentInfraGRPCMapper {
	return &PaymentInfraGRPCMapper{}
}

// ToCreateRequest maps domain input to gRPC request
func (m *PaymentInfraGRPCMapper) ToCreateRequest(paymentRackID string) *dto.GetPaymentInfraByIDRequest {
	return &dto.GetPaymentInfraByIDRequest{
		PaymentRackId: paymentRackID,
	}
}

// ToDomain maps gRPC response to domain model
func (m *PaymentInfraGRPCMapper) ToDomain(response *dto.GetPaymentInfraByIDResponse) *model.PaymentInfra {
	if response == nil {
		return nil
	}

	paymentInfra := &model.PaymentInfra{}

	// Map response metadata
	if response.Response != nil {
		paymentInfra.TransactionID = response.Response.TransactionId
		paymentInfra.Message = response.Response.Message
		paymentInfra.Status = m.mapResponseStatus(response.Response.Status)
	}

	// Map payment rack
	if response.PaymentRack != nil {
		paymentInfra.PaymentRack = &model.PaymentRack{
			ID:          int(response.PaymentRack.Id),
			Description: response.PaymentRack.Description,
			Address:     response.PaymentRack.Address,
		}
	}

	// Map installation
	if response.Installation != nil {
		paymentInfra.Installation = &model.PaymentInstallation{
			ID:       int(response.Installation.Id),
			Name:     response.Installation.Name,
			Region:   response.Installation.Region,
			City:     response.Installation.City,
			Address:  response.Installation.Address,
			ImageURL: response.Installation.ImageUrl,
		}
	}

	// Map booking times
	if response.BookingTimes != nil {
		paymentInfra.BookingTimes = make([]model.PaymentBookingTime, len(response.BookingTimes))
		for i, bt := range response.BookingTimes {
			paymentInfra.BookingTimes[i] = model.PaymentBookingTime{
				ID:              int(bt.Id),
				Name:            bt.Name,
				UnitMeasurement: m.mapUnitMeasurement(bt.UnitMeasurement),
				Amount:          int(bt.Amount),
			}
		}
	}

	return paymentInfra
}

// mapResponseStatus converts gRPC response status to domain status
func (m *PaymentInfraGRPCMapper) mapResponseStatus(status dto.PaymentManagerResponseStatus) model.ResponseStatus {
	switch status {
	case dto.PaymentManagerResponseStatus_RESPONSE_STATUS_OK:
		return model.ResponseStatusOK
	case dto.PaymentManagerResponseStatus_RESPONSE_STATUS_ERROR:
		return model.ResponseStatusError
	default:
		return model.ResponseStatusUnspecified
	}
}

// mapUnitMeasurement converts gRPC unit measurement to domain unit measurement
func (m *PaymentInfraGRPCMapper) mapUnitMeasurement(unit dto.UnitMeasurement) model.UnitMeasurement {
	switch unit {
	case dto.UnitMeasurement_HOUR:
		return model.UnitMeasurementHour
	case dto.UnitMeasurement_DAY:
		return model.UnitMeasurementDay
	case dto.UnitMeasurement_WEEK:
		return model.UnitMeasurementWeek
	case dto.UnitMeasurement_MONTH:
		return model.UnitMeasurementMonth
	default:
		return model.UnitMeasurementUnspecified
	}
}
