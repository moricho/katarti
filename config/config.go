package config

import (
	"os"
)

// XdsConfig represents the configuration for xDS server
type XdsConfig struct {
	Addr string
}

// CacheConfig represents the configuration for cache
type CacheConfig struct {
	Addr string
}

// Config represents application config.
type Config struct {
	Xds   XdsConfig
	Redis CacheConfig
}

// New return Config struct.
func New() (Config, error) {
	// xDS configuration
	xdsConfig := XdsConfig{Addr: os.Getenv("XDS_ADDR")}

	// redis configuration
	redisConf := RedisConfig{Addr: os.Getenv("REDIS_ADDR")}

	conf := Config{
		Xds:   xdsConfig,
		Redis: redisConf,
	}
	return conf, nil
}
