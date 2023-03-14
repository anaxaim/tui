package repository

import (
	"context"
)

type Repository interface {
	Close() error
	Ping(ctx context.Context) error
}
