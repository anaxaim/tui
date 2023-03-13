package main

import (
	"flag"

	"github.com/sirupsen/logrus"

	"github.com/anaxaim/tui/pkg/config"
	"github.com/anaxaim/tui/pkg/server"
)

var ( //nolint: gofumpt
	appConfig = flag.String("config", "config/app.yaml", "application config path")
)

func main() {
	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.JSONFormatter{})

	conf, err := config.Parse(*appConfig)
	if err != nil {
		logger.Fatalf("Failed to parse config: %v", err)
	}

	s, err := server.New(conf, logger)
	if err != nil {
		logger.Fatalf("Init server failed: %v", err)
	}

	if err := s.Run(); err != nil {
		logger.Fatalf("Run server failed: %v", err)
	}
}
