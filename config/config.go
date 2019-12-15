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
	Cache CacheConfig
}

// New return Config struct.
func New() (Config, error) {
	// xDS configuration
	xdsConfig := XdsConfig{Addr: os.Getenv("XDS_ADDR")}

	// cache configuration
	cacheConf := CacheConfig{Addr: os.Getenv("CACHE_ADDR")}

	conf := Config{
		Xds:   xdsConfig,
		Cache: cacheConf,
	}
	return conf, nil
}
