package server

import (
	"github.com/google/wire"
	"github.com/shaned24/tough-notes-storage/internal/pkg/config"
	"github.com/shaned24/tough-notes-storage/internal/pkg/gateway"
)

var Providers = wire.NewSet(
	wire.Struct(new(Server), "*"),
	ProvideServerConfig,
	ProvideHttpGateway,
)

func ProvideServerConfig() Config {
	port := config.Getenv("GRPC_PORT", "50051")
	host := config.Getenv("GRPC_HOST", "0.0.0.0")
	protocol := config.Getenv("GRPC_PROTOCOL", "tcp")

	return Config{
		Host:     host,
		Port:     port,
		Protocol: protocol,
	}
}

func ProvideHttpGateway(c Config) *gateway.HttpGateway {
	port := config.Getenv("GATEWAY_PORT", "8080")
	return gateway.NewHttpGateway(port, c.Host, c.Protocol)
}
