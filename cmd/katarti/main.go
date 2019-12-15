package main

import (
	"github.com/moricho/katarti/config"
	"github.com/moricho/katarti/internal/driver/logger"
)

func main() {
	_, err := config.New()
	if err != nil {
		panic(err)
	}

	logger := logger.New()
	defer logger.Sync()

	return
}
