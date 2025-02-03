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
	GroupId           string
}

func NewDefaultConfig() *Config {
	return &Config{
		KeepaliveInterval: 40 * time.Second,
		KeepaliveTimeout:  3 * time.Second,
		SendTimeout:       1 * time.Second,
	}
}

func NewEnvConfig() *Config {
	return &Config{
		Address:           env.String("STREAMS_CLIENT_ADDRESS"),
		KeepaliveInterval: time.Duration(env.IntWithDefault("STREAMS_CLIENT_KEEPALIVE_INTERVAL", 40)) * time.Second,
		KeepaliveTimeout:  time.Duration(env.IntWithDefault("STREAMS_CLIENT_KEEPALIVE_TIMEOUT", 3)) * time.Second,
		SendTimeout:       time.Duration(env.IntWithDefault("STREAMS_CLIENT_SEND_TIMEOUT", 1)) * time.Second,
		GroupId:           env.String("STREAMS_CLIENT_GROUP_ID"),
	}
}
