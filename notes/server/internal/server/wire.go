package server

import (
	"github.com/google/wire"
	"github.com/shaned24/tough-notes-storage/notes/server/internal/notes"
)

var Providers = wire.NewSet(
	notes.Providers,
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
