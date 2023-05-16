package worker

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/pkg/errors"

	"github.com/anaxaim/tui/backend/pkg/config"
	"github.com/anaxaim/tui/backend/pkg/database"
	"github.com/anaxaim/tui/backend/pkg/model"
	"github.com/anaxaim/tui/backend/pkg/repository"
	"github.com/anaxaim/tui/backend/pkg/service"
)

const (
	TaskExecute = "task:execute"
)

type moduleService interface {
	Execute(terraformVersion, command, id string) ([]byte, error)
}

type handler struct {
	moduleService moduleService
}

func newHandler(conf *config.Config) (*handler, error) {
	db, err := database.NewMongoClient(&conf.DB)
	if err != nil {
		return nil, errors.Wrap(err, "task: db init failed")
	}

	repo := repository.NewRepository(db)
	moduleService := service.NewModuleService(repo.Module(), repo.Credential())

	return &handler{moduleService: moduleService}, nil
}

func (h handler) HandleExecuteCommand(_ context.Context, t *asynq.Task) error {
	payload := new(model.ExecuteCommand)
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return err
	}

	_, err := h.moduleService.Execute("latest", payload.Command, payload.ModuleID)
	if err != nil {
		return err
	}

	return nil
}
