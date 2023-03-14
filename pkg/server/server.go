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
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/anaxaim/tui/pkg/common"
	"github.com/anaxaim/tui/pkg/config"
	"github.com/anaxaim/tui/pkg/database"
	"github.com/anaxaim/tui/pkg/repository"
	"github.com/anaxaim/tui/pkg/utils"
	"github.com/anaxaim/tui/pkg/version"
)

func New(conf *config.Config, logger *logrus.Logger) (*Server, error) {
	db, err := database.NewMongoClient(&conf.DB)
	if err != nil {
		return nil, errors.Wrap(err, "db init failed")
	}

	repo := repository.NewRepository(db)

	gin.SetMode(conf.Server.ENV)

	e := gin.New()
	e.Use(gin.Recovery())

	return &Server{
		engine:     e,
		config:     conf,
		logger:     logger,
		repository: repo,
	}, nil
}

type Server struct {
	engine *gin.Engine
	config *config.Config
	logger *logrus.Logger

	repository repository.Repository
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
	if err := s.repository.Close(); err != nil {
		s.logger.Warnf("failed to close repository, %v", err)
	}
}

func (s *Server) initRouter() {
	root := s.engine

	root.GET("/", common.WrapFunc(s.getRoutes))
	root.GET("/version", common.WrapFunc(version.Get))
	root.GET("/healthz", common.WrapFunc(s.Ping))
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

type Status struct {
	Ping         bool `json:"ping"`
	DBRepository bool `json:"dbRepository"`
}

func (s *Server) Ping() *Status {
	status := &Status{Ping: true}

	ctx, cannel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cannel()

	if err := s.repository.Ping(ctx); err == nil {
		status.DBRepository = true
	}

	return status
}
