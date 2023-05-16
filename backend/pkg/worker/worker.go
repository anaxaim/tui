package worker

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"

	"github.com/anaxaim/tui/backend/pkg/config"
	"github.com/anaxaim/tui/backend/pkg/middleware"
)

type Worker struct {
	mux     *asynq.ServeMux
	server  *asynq.Server
	logger  *logrus.Logger
	handler *handler
}

func New(conf *config.Config, logger *logrus.Logger) (*Worker, error) {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: conf.Redis.String()},
		asynq.Config{Concurrency: 20},
	)

	mux := asynq.NewServeMux()
	mux.Use(middleware.LogTaskMiddleware(logger))

	h, err := newHandler(conf)
	if err != nil {
		return nil, err
	}

	w := &Worker{
		mux:     mux,
		server:  srv,
		handler: h,
	}
	w.routes()

	return w, nil
}

func (w *Worker) routes() {
	w.mux.HandleFunc(TaskExecute, w.handler.HandleExecuteCommand)
}

func (w *Worker) Run() error {
	go func() {
		if err := w.server.Run(w.mux); err != nil {
			w.logger.Fatalf("Failed to start worker server, %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ch := <-sig
	w.logger.Infof("Receive signal: %s", ch)

	w.server.Stop()
	w.server.Shutdown()

	w.logger.Infof("worker has been stopped")

	return nil
}
