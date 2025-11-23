package config

import (
	"graphql-payment-bff/internal/application/service"
	"graphql-payment-bff/internal/domain/ports"
	"graphql-payment-bff/internal/infrastructure/inbound/graphql/resolver"
	"graphql-payment-bff/internal/infrastructure/outbound/grpc/client"
	"time"
)

// Config holds all configuration for the application
type Config struct {
	Server ServerConfig
	GRPC   GRPCConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// GRPCConfig holds gRPC client configuration
type GRPCConfig struct {
	PaymentServiceAddress string
	Timeout               time.Duration
}

// Container holds all application dependencies
type Container struct {
	// Services
	PaymentInfraService ports.PaymentInfraService

	// Resolvers
	GraphQLResolver *resolver.Resolver

	// Infrastructure
	PaymentServiceClient *client.PaymentServiceGRPCClient
}

// NewContainer creates a new dependency injection container
func NewContainer(config Config) (*Container, error) {
	container := &Container{}

	// Initialize gRPC client
	paymentClient, err := client.NewPaymentServiceGRPCClient(
		config.GRPC.PaymentServiceAddress,
		config.GRPC.Timeout,
	)
	if err != nil {
		return nil, err
	}
	container.PaymentServiceClient = paymentClient

	// Initialize services
	container.PaymentInfraService = service.NewPaymentInfraService(paymentClient)

	// Initialize resolvers
	container.GraphQLResolver = resolver.NewResolver(container.PaymentInfraService)

	return container, nil
}

// Close closes all resources
func (c *Container) Close() error {
	if c.PaymentServiceClient != nil {
		return c.PaymentServiceClient.Close()
	}
	return nil
}

// DefaultConfig returns default configuration
func DefaultConfig() Config {
	return Config{
		Server: ServerConfig{
			Port:         "8080",
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
		GRPC: GRPCConfig{
			PaymentServiceAddress: "localhost:50051",
			Timeout:               10 * time.Second,
		},
	}
}
