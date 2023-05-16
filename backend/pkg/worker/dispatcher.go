package worker

import (
	"github.com/hibiken/asynq"

	"github.com/anaxaim/tui/backend/pkg/config"
)

type Dispatcher struct {
	Client *asynq.Client
}

func NewDispatcher(conf *config.RedisConfig) (*Dispatcher, error) {
	redisOpt := asynq.RedisClientOpt{
		Addr: conf.String(),
	}

	client := asynq.NewClient(redisOpt)

	return &Dispatcher{Client: client}, nil
}
