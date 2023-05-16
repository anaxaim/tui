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

	"github.com/anaxaim/tui/backend/pkg/authentication"
	"github.com/anaxaim/tui/backend/pkg/common"
	"github.com/anaxaim/tui/backend/pkg/config"
	"github.com/anaxaim/tui/backend/pkg/container"
	"github.com/anaxaim/tui/backend/pkg/controller"
	"github.com/anaxaim/tui/backend/pkg/database"
	"github.com/anaxaim/tui/backend/pkg/middleware"
	"github.com/anaxaim/tui/backend/pkg/repository"
	"github.com/anaxaim/tui/backend/pkg/service"
	"github.com/anaxaim/tui/backend/pkg/utils"
	"github.com/anaxaim/tui/backend/pkg/version"
	"github.com/anaxaim/tui/backend/pkg/worker"
)

func New(conf *config.Config, logger *logrus.Logger) (*Server, error) {
	db, err := database.NewMongoClient(&conf.DB)
	if err != nil {
		return nil, errors.Wrap(err, "db init failed")
	}

	repo := repository.NewRepository(db)

	userService := service.NewUserService(repo.User())
	moduleService := service.NewModuleService(repo.Module(), repo.Credential())
	credentialService := service.NewCredentialService(repo.Credential())
	jwtService := authentication.NewJWTService(conf.Server.JWTSecret)
	terraformService := container.NewTerraformService()

	dispatcher, err := worker.NewDispatcher(&conf.Redis)
	if err != nil {
		return nil, errors.Wrap(err, "dispatcher init failed")
	}

	userController := controller.NewUserController(userService)
	moduleController := controller.NewModuleController(moduleService, terraformService, dispatcher.Client)
	credentialController := controller.NewCredentialController(credentialService)
	authController := controller.NewAuthController(userService, jwtService)

	controllers := []controller.Controller{userController, moduleController, credentialController, authController}

	gin.SetMode(conf.Server.ENV)

	engine := gin.New()
	engine.Use(
		gin.Recovery(),
		middleware.CORSMiddleware(),
		middleware.RequestInfoMiddleware(&utils.RequestInfoFactory{APIPrefixes: utils.NewString("api")}),
		middleware.LogMiddleware(logger, "/"),
		middleware.AuthenticationMiddleware(jwtService, repo.User()),
	)

	return &Server{
		engine:      engine,
		config:      conf,
		logger:      logger,
		repository:  repo,
		controllers: controllers,
	}, nil
}

type Server struct {
	engine *gin.Engine
	config *config.Config
	logger *logrus.Logger

	repository repository.Repository

	controllers []controller.Controller
}

func (s *Server) Run() error {
	defer s.Close()

	s.initRouter()

	serverHost := s.config.Server.Address
	if os.Getenv("SERVER_HOST") != "" {
		serverHost = os.Getenv("SERVER_HOST")
	}

	addr := fmt.Sprintf("%s:%d", serverHost, s.config.Server.Port)
	s.logger.Infof("Start server on: %s", addr)

	server := &http.Server{
		Addr:              addr,
		Handler:           s.engine,
		ReadHeaderTimeout: 2 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
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

	// register non-resource routers
	root.GET("/", common.WrapFunc(s.getRoutes))
	root.GET("/healthz", common.WrapFunc(s.Ping))
	root.GET("/version", common.WrapFunc(version.Get))

	api := root.Group("/api/v1")

	controllers := make([]string, 0, len(s.controllers))
	for _, router := range s.controllers {
		router.RegisterRoute(api)
		controllers = append(controllers, router.Name())
	}

	logrus.Infof("server enabled controllers: %v", controllers)
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

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := s.repository.Ping(ctx); err == nil {
		status.DBRepository = true
	}

	return status
}
