//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/shaned24/tough-notes-storage/internal/server"
)

func setupServer() (*server.Server, error) {
	wire.Build(wireProviders)
	return nil, nil
}
