package main

import (
	"flag"

	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"

	"github.com/anaxaim/tui/backend/pkg/config"
	"github.com/anaxaim/tui/backend/pkg/server"
)

var appConfig = flag.String("config", "config/app.yaml", "application config path")

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
