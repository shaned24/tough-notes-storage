package server

import (
	"github.com/google/wire"
	"github.com/shaned24/tough-notes-storage/notes/server/internal/notes"
	"github.com/shaned24/tough-notes-storage/notes/server/internal/pkg/database"
)

var Providers = wire.NewSet(
	wire.Struct(new(Server), "*"),
	ProvideServerConfig,
	database.Providers,
	notes.Providers,
)

func ProvideServerConfig() Config {
	return Config{
		Host:     "0.0.0.0",
		Port:     "50051",
		Protocol: "tcp",
	}
}
