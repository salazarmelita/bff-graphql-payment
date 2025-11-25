package config

import (
	"graphql-payment-bff/internal/application/service"
	"graphql-payment-bff/internal/domain/ports"
	"graphql-payment-bff/internal/infrastructure/inbound/graphql/resolver"
	"graphql-payment-bff/internal/infrastructure/outbound/grpc/client"
	"time"
)

// Config contiene toda la configuración de la aplicación
type Config struct {
	Server ServerConfig
	GRPC   GRPCConfig
}

// ServerConfig contiene la configuración del servidor
type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// GRPCConfig contiene la configuración del cliente gRPC
type GRPCConfig struct {
	PaymentServiceAddress string
	Timeout               time.Duration
}

// Container contiene todas las dependencias de la aplicación
type Container struct {
	// Servicios
	PaymentInfraService ports.PaymentInfraService

	// Resolvers
	GraphQLResolver *resolver.Resolver

	// Infraestructura
	PaymentServiceClient *client.PaymentServiceGRPCClient
}

// NewContainer crea un nuevo contenedor de inyección de dependencias
func NewContainer(config Config) (*Container, error) {
	container := &Container{}

	// Inicializar cliente gRPC
	paymentClient, err := client.NewPaymentServiceGRPCClient(
		config.GRPC.PaymentServiceAddress,
		config.GRPC.Timeout,
	)
	if err != nil {
		return nil, err
	}
	container.PaymentServiceClient = paymentClient

	// Inicializar servicios
	container.PaymentInfraService = service.NewPaymentInfraService(paymentClient)

	// Inicializar resolvers
	container.GraphQLResolver = resolver.NewResolver(container.PaymentInfraService)

	return container, nil
}

// Close cierra todos los recursos
func (c *Container) Close() error {
	if c.PaymentServiceClient != nil {
		return c.PaymentServiceClient.Close()
	}
	return nil
}

// DefaultConfig devuelve la configuración por defecto
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
