package config

import (
	"time"

	"github.com/Becklyn/go-wire-core/env"
)

type Config struct {
	Address           string
	KeepaliveInterval time.Duration
	KeepaliveTimeout  time.Duration
	SendTimeout       time.Duration
	Namespace         string
}

func NewDefaultConfig() *Config {
	return &Config{
		KeepaliveInterval: 40 * time.Second,
		KeepaliveTimeout:  3 * time.Second,
		SendTimeout:       1 * time.Second,
		Namespace:         "",
	}
}

func NewEnvConfig() *Config {
	return &Config{
		Address:           env.String("DEPLOYMENTS_CLIENT_ADDRESS"),
		KeepaliveInterval: time.Duration(env.IntWithDefault("DEPLOYMENTS_CLIENT_KEEPALIVE_INTERVAL", 40)) * time.Second,
		KeepaliveTimeout:  time.Duration(env.IntWithDefault("DEPLOYMENTS_CLIENT_KEEPALIVE_TIMEOUT", 3)) * time.Second,
		SendTimeout:       time.Duration(env.IntWithDefault("DEPLOYMENTS_CLIENT_SEND_TIMEOUT", 1)) * time.Second,
		Namespace:         env.String("DEPLOYMENTS_NAMESPACE"),
	}
}
