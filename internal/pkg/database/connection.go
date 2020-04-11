package database

import (
	"context"
	"time"
)

type Connection interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Name() string
}

type Config struct {
	Host              string
	Port              string
	ConnectionTimeout time.Duration
}

func NewDatabaseConfig(host string, port string, connectionTimeout time.Duration) *Config {
	return &Config{Host: host, Port: port, ConnectionTimeout: connectionTimeout}
}
