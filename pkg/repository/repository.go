package repository

import (
	"context"

	"github.com/anaxaim/tui/pkg/database"
)

func NewRepository(db *database.MongoDB) Repository {
	r := &repository{
		db: db,
	}

	return r
}

type repository struct {
	db *database.MongoDB
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
