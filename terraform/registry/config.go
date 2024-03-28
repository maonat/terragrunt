package registry

import (
	"net"
	"strconv"
	"time"
)

const (
	defaultShutdownTimeout = time.Second * 30
)

type Option func(Config) Config

func WithHostname(hostname string) Option {
	return func(cfg Config) Config {
		if hostname != "" {
			cfg.hostname = hostname
		}
		return cfg
	}
}

func WithPort(port int) Option {
	return func(cfg Config) Config {
		if port != 0 {
			cfg.port = port
		}
		return cfg
	}
}

func WithToken(token string) Option {
	return func(cfg Config) Config {
		cfg.token = token
		return cfg
	}
}

func WithProviderCacheDir(cacheDir string) Option {
	return func(cfg Config) Config {
		cfg.providerCacheDir = cacheDir
		return cfg
	}
}

func WithProviderCompleteLock(completeLock bool) Option {
	return func(cfg Config) Config {
		cfg.providerCompleteLock = completeLock
		return cfg
	}
}

type Config struct {
	hostname        string
	port            int
	token           string
	shutdownTimeout time.Duration

	providerCacheDir     string
	providerCompleteLock bool
}

func NewConfig(opts ...Option) *Config {
	cfg := &Config{
		shutdownTimeout: defaultShutdownTimeout,
	}

	return cfg.WithOptions(opts...)
}

func (config *Config) WithOptions(opts ...Option) *Config {
	for _, opt := range opts {
		*config = opt(*config)
	}

	return config
}

func (cfg *Config) Addr() string {
	return net.JoinHostPort(cfg.hostname, strconv.Itoa(cfg.port))
}