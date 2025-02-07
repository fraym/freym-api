package config

import (
	"time"

	"github.com/Becklyn/go-wire-core/env"
)

type Config struct {
	Address           string
	KeepaliveInterval time.Duration
	KeepaliveTimeout  time.Duration
	AppPrefix         string
	GrpcPort          string
}

func NewDefaultConfig() *Config {
	return &Config{
		KeepaliveInterval: 40 * time.Second,
		KeepaliveTimeout:  3 * time.Second,
		GrpcPort:          "9000",
	}
}

func NewEnvConfig() *Config {
	return &Config{
		Address:           env.String("SYNC_CLIENT_ADDRESS"),
		KeepaliveInterval: time.Duration(env.IntWithDefault("SYNC_CLIENT_KEEPALIVE_INTERVAL", 40)) * time.Second,
		KeepaliveTimeout:  time.Duration(env.IntWithDefault("SYNC_CLIENT_KEEPALIVE_TIMEOUT", 3)) * time.Second,
		AppPrefix:         env.StringWithDefault("SYNC_CLIENT_APP_PREFIX", ""),
		GrpcPort:          env.StringWithDefault("SYNC_CLIENT_GRPC_PORT", "9000"),
	}
}
