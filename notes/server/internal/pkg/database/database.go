package database

import (
	"context"
	"time"
)

type Database interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
}

type Config struct {
	Host              string
	Port              string
	ConnectionTimeout time.Duration
}

func NewDatabaseConfig(host string, port string, connectionTimeout time.Duration) *Config {
	return &Config{Host: host, Port: port, ConnectionTimeout: connectionTimeout}
}
