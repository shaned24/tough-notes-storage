package server

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	wire.Struct(new(Server), "*"),
	ProvideServerConfig,
)

func ProvideServerConfig() Config {
	return Config{
		Host:     "0.0.0.0",
		Port:     "50051",
		Protocol: "tcp",
	}
}
