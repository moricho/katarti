package main

import (
	"github.com/moricho/katariti/config"
	"github.com/moricho/katariti/internal/driver/logger"
)

func main() {
	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	logger, err := logger.New()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	return
}
