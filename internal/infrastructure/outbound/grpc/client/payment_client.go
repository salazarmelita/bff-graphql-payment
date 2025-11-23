package client

import (
	"context"
	"fmt"
	"graphql-payment-bff/internal/application/ports"
	"graphql-payment-bff/internal/domain/exception"
	"graphql-payment-bff/internal/domain/model"
	"graphql-payment-bff/internal/infrastructure/outbound/grpc/dto"
	"graphql-payment-bff/internal/infrastructure/outbound/grpc/mapper"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PaymentServiceGRPCClient implementa PaymentInfraRepository usando gRPC
type PaymentServiceGRPCClient struct {
	conn    *grpc.ClientConn
	mapper  *mapper.PaymentInfraGRPCMapper
	timeout time.Duration
}

// NewPaymentServiceGRPCClient crea un nuevo cliente gRPC para el servicio de pagos
func NewPaymentServiceGRPCClient(serverAddress string, timeout time.Duration) (*PaymentServiceGRPCClient, error) {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to payment service: %w", err)
	}

	return &PaymentServiceGRPCClient{
		conn:    conn,
		mapper:  mapper.NewPaymentInfraGRPCMapper(),
		timeout: timeout,
	}, nil
}

// GetPaymentInfraByID implementa PaymentInfraRepository.GetPaymentInfraByID
func (c *PaymentServiceGRPCClient) GetPaymentInfraByID(ctx context.Context, paymentRackID string) (*model.PaymentInfra, error) {
	// Crear contexto con timeout
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	// Crear request
	request := c.mapper.ToCreateRequest(paymentRackID)

	// Respuesta mock por ahora (ya que no tenemos el servidor gRPC real)
	// En una implementación real, esto llamaría al servicio gRPC actual
	response := c.mockGRPCCall(request)

	// Manejar errores gRPC
	if response == nil {
		return nil, exception.ErrPaymentInfraServiceUnavailable
	}

	if response.Response != nil && response.Response.Status == dto.PaymentManagerResponseStatus_RESPONSE_STATUS_ERROR {
		return nil, exception.ErrPaymentRackNotFound
	}

	// Mapear respuesta a modelo de dominio
	return c.mapper.ToDomain(response), nil
}

// Close cierra la conexión gRPC
func (c *PaymentServiceGRPCClient) Close() error {
	return c.conn.Close()
}

// mockGRPCCall simula una llamada gRPC para propósitos de desarrollo/testing
// En producción, esto se reemplazaría con la llamada real al cliente gRPC
func (c *PaymentServiceGRPCClient) mockGRPCCall(request *dto.GetPaymentInfraByIDRequest) *dto.GetPaymentInfraByIDResponse {
	// Simular diferentes respuestas basadas en el ID del rack para testing
	if request.PaymentRackId == "" {
		return &dto.GetPaymentInfraByIDResponse{
			Response: &dto.PaymentManagerGenericResponse{
				TransactionId: time.Now().Format("20060102150405"),
				Message:       "ID de rack inválido",
				Status:        dto.PaymentManagerResponseStatus_RESPONSE_STATUS_ERROR,
			},
		}
	}

	// Respuesta mock exitosa
	return &dto.GetPaymentInfraByIDResponse{
		Response: &dto.PaymentManagerGenericResponse{
			TransactionId: time.Now().Format("20060102150405"),
			Message:       "Success",
			Status:        dto.PaymentManagerResponseStatus_RESPONSE_STATUS_OK,
		},
		PaymentRack: &dto.PaymentRackRecord{
			Id:      1,
			Address: "Chicureo",
		},
		Installation: &dto.PaymentInstallationRecord{
			Id:       1,
			Name:     "DEV PAGO",
			Address:  "Chicureo",
			ImageUrl: "https://www.image.cl/image.jpg",
		},
		BookingTimes: []*dto.PaymentBookingTimeRecord{
			{
				Id:              1,
				Name:            "Express (1 día)",
				UnitMeasurement: dto.UnitMeasurement_DAY,
				Amount:          1,
			},
			{
				Id:              2,
				Name:            "Normal (3 días)",
				UnitMeasurement: dto.UnitMeasurement_DAY,
				Amount:          3,
			},
		},
	}
}

// mapGRPCError mapea errores gRPC a errores de dominio
func (c *PaymentServiceGRPCClient) mapGRPCError(err error) error {
	if err == nil {
		return nil
	}

	statusErr, ok := status.FromError(err)
	if !ok {
		return exception.ErrPaymentInfraServiceUnavailable
	}

	switch statusErr.Code() {
	case codes.NotFound:
		return exception.ErrPaymentRackNotFound
	case codes.InvalidArgument:
		return exception.ErrInvalidPaymentRackID
	case codes.Unavailable:
		return exception.ErrPaymentInfraServiceUnavailable
	default:
		return exception.ErrPaymentInfraServiceUnavailable
	}
}

// Asegurar que PaymentServiceGRPCClient implementa PaymentInfraRepository
var _ ports.PaymentInfraRepository = (*PaymentServiceGRPCClient)(nil)
