//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/shaned24/tough-notes-storage/notes/server/internal/server"
)

func setupServer() (*server.Server, error) {
	wire.Build(
		server.Providers,
	)
	return nil, nil
}
