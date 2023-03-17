package repository

import (
	"context"

	"github.com/anaxaim/tui/server/pkg/database"
)

func NewRepository(db *database.MongoDB) Repository {
	r := &repository{
		db:   db,
		user: newUserRepository(db),
	}

	return r
}

type repository struct {
	db   *database.MongoDB
	user UserRepository
}

func (r *repository) User() UserRepository {
	return r.user
}

func (r *repository) Close() error {
	if r.db.Client != nil {
		if err := r.db.Disconnect(context.Background()); err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) Ping(ctx context.Context) error {
	if r.db.Client != nil {
		if err := r.db.Ping(ctx, nil); err != nil {
			return err
		}
	}

	return nil
}
