package service_test

import (
	"context"
	"graphql-payment-bff/internal/application/service"
	"graphql-payment-bff/internal/domain/exception"
	"graphql-payment-bff/internal/domain/model"
	"testing"
)

// MockPaymentInfraRepository implements PaymentInfraRepository for testing
type MockPaymentInfraRepository struct {
	mockFunc func(ctx context.Context, paymentRackID string) (*model.PaymentInfra, error)
}

func (m *MockPaymentInfraRepository) GetPaymentInfraByID(ctx context.Context, paymentRackID string) (*model.PaymentInfra, error) {
	if m.mockFunc != nil {
		return m.mockFunc(ctx, paymentRackID)
	}
	return nil, exception.ErrPaymentRackNotFound
}

func TestPaymentInfraService_GetPaymentInfraByID(t *testing.T) {
	tests := []struct {
		name           string
		paymentRackID  string
		mockResponse   *model.PaymentInfra
		mockError      error
		expectedError  error
		expectedResult *model.PaymentInfra
	}{
		{
			name:          "successful retrieval",
			paymentRackID: "valid-rack-id",
			mockResponse: &model.PaymentInfra{
				TransactionID: "tx-123",
				Message:       "Success",
				Status:        model.ResponseStatusOK,
			},
			mockError:     nil,
			expectedError: nil,
			expectedResult: &model.PaymentInfra{
				TransactionID: "tx-123",
				Message:       "Success",
				Status:        model.ResponseStatusOK,
			},
		},
		{
			name:           "empty rack ID",
			paymentRackID:  "",
			mockResponse:   nil,
			mockError:      nil,
			expectedError:  exception.ErrInvalidPaymentRackID,
			expectedResult: nil,
		},
		{
			name:           "whitespace only rack ID",
			paymentRackID:  "   ",
			mockResponse:   nil,
			mockError:      nil,
			expectedError:  exception.ErrInvalidPaymentRackID,
			expectedResult: nil,
		},
		{
			name:           "repository error",
			paymentRackID:  "valid-rack-id",
			mockResponse:   nil,
			mockError:      exception.ErrPaymentRackNotFound,
			expectedError:  exception.ErrPaymentRackNotFound,
			expectedResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockPaymentInfraRepository{
				mockFunc: func(ctx context.Context, paymentRackID string) (*model.PaymentInfra, error) {
					return tt.mockResponse, tt.mockError
				},
			}

			service := service.NewPaymentInfraService(mockRepo)

			result, err := service.GetPaymentInfraByID(context.Background(), tt.paymentRackID)

			// Check error
			if tt.expectedError != nil {
				if err == nil {
					t.Errorf("Expected error %v, got nil", tt.expectedError)
					return
				}
				if err != tt.expectedError {
					t.Errorf("Expected error %v, got %v", tt.expectedError, err)
					return
				}
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Check result
			if tt.expectedResult == nil && result != nil {
				t.Errorf("Expected nil result, got %v", result)
			} else if tt.expectedResult != nil && result == nil {
				t.Errorf("Expected result %v, got nil", tt.expectedResult)
			} else if tt.expectedResult != nil && result != nil {
				if result.TransactionID != tt.expectedResult.TransactionID {
					t.Errorf("Expected transaction ID %s, got %s", tt.expectedResult.TransactionID, result.TransactionID)
				}
				if result.Status != tt.expectedResult.Status {
					t.Errorf("Expected status %s, got %s", tt.expectedResult.Status, result.Status)
				}
			}
		})
	}
}
