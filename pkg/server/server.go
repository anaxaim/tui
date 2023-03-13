package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/anaxaim/tui/pkg/common"
	"github.com/anaxaim/tui/pkg/config"
	"github.com/anaxaim/tui/pkg/utils"
	"github.com/anaxaim/tui/pkg/version"
)

func New(conf *config.Config, logger *logrus.Logger) (*Server, error) {
	gin.SetMode(conf.Server.ENV)

	e := gin.New()
	e.Use(gin.Recovery())

	return &Server{
		engine: e,
		config: conf,
		logger: logger,
	}, nil
}

type Server struct {
	engine *gin.Engine
	config *config.Config
	logger *logrus.Logger
}

func (s *Server) Run() error {
	defer s.Close()

	s.initRouter()

	addr := fmt.Sprintf("%s:%d", s.config.Server.Address, s.config.Server.Port)
	s.logger.Infof("Start server on: %s", addr)

	server := &http.Server{
		Addr:              addr,
		Handler:           s.engine,
		ReadHeaderTimeout: 2 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			s.logger.Fatalf("Failed to start server, %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.config.Server.GracefulShutdownPeriod)*time.Second)
	defer cancel()

	ch := <-sig
	s.logger.Infof("Receive signal: %s", ch)

	return server.Shutdown(ctx)
}

func (s *Server) Close() {
	// close dbs ...
}

func (s *Server) initRouter() {
	root := s.engine

	root.GET("/", common.WrapFunc(s.getRoutes))
	root.GET("/version", common.WrapFunc(version.Get))
}

func (s *Server) getRoutes() []string {
	paths := utils.NewString()
	for _, r := range s.engine.Routes() {
		if r.Path != "" {
			paths.Insert(r.Path)
		}
	}
	return paths.Slice()
}
