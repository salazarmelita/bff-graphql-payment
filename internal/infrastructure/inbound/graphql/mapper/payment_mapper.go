package mapper

import (
	"graphql-payment-bff/graph/model"
	domainModel "graphql-payment-bff/internal/domain/model"
)

// PaymentInfraGraphQLMapper handles mapping between domain models and GraphQL DTOs
type PaymentInfraGraphQLMapper struct{}

// NewPaymentInfraGraphQLMapper creates a new mapper instance
func NewPaymentInfraGraphQLMapper() *PaymentInfraGraphQLMapper {
	return &PaymentInfraGraphQLMapper{}
}

// ToGraphQLResponse maps domain model to GraphQL response model
func (m *PaymentInfraGraphQLMapper) ToGraphQLResponse(paymentInfra *domainModel.PaymentInfra) *model.PaymentInfraResponse {
	if paymentInfra == nil {
		return nil
	}

	response := &model.PaymentInfraResponse{
		TransactionID: paymentInfra.TransactionID,
		Message:       paymentInfra.Message,
		Status:        m.mapResponseStatus(paymentInfra.Status),
		BookingTimes:  []*model.PaymentBookingTime{},
	}

	// Map payment rack
	if paymentInfra.PaymentRack != nil {
		response.PaymentRack = &model.PaymentRack{
			ID:          paymentInfra.PaymentRack.ID,
			Description: paymentInfra.PaymentRack.Description,
			Address:     paymentInfra.PaymentRack.Address,
		}
	}

	// Map installation
	if paymentInfra.Installation != nil {
		response.Installation = &model.PaymentInstallation{
			ID:       paymentInfra.Installation.ID,
			Name:     paymentInfra.Installation.Name,
			Region:   paymentInfra.Installation.Region,
			City:     paymentInfra.Installation.City,
			Address:  paymentInfra.Installation.Address,
			ImageURL: paymentInfra.Installation.ImageURL,
		}
	}

	// Map booking times
	for _, bt := range paymentInfra.BookingTimes {
		response.BookingTimes = append(response.BookingTimes, &model.PaymentBookingTime{
			ID:              bt.ID,
			Name:            bt.Name,
			UnitMeasurement: m.mapUnitMeasurement(bt.UnitMeasurement),
			Amount:          bt.Amount,
		})
	}

	return response
}

// mapResponseStatus converts domain response status to GraphQL status
func (m *PaymentInfraGraphQLMapper) mapResponseStatus(status domainModel.ResponseStatus) model.ResponseStatus {
	switch status {
	case domainModel.ResponseStatusOK:
		return model.ResponseStatusResponseStatusOk
	case domainModel.ResponseStatusError:
		return model.ResponseStatusResponseStatusError
	default:
		return model.ResponseStatusResponseStatusUnspecified
	}
}

// mapUnitMeasurement converts domain unit measurement to GraphQL unit measurement
func (m *PaymentInfraGraphQLMapper) mapUnitMeasurement(unit domainModel.UnitMeasurement) model.UnitMeasurement {
	switch unit {
	case domainModel.UnitMeasurementHour:
		return model.UnitMeasurementHour
	case domainModel.UnitMeasurementDay:
		return model.UnitMeasurementDay
	case domainModel.UnitMeasurementWeek:
		return model.UnitMeasurementWeek
	case domainModel.UnitMeasurementMonth:
		return model.UnitMeasurementMonth
	default:
		return model.UnitMeasurementUnspecified
	}
}
