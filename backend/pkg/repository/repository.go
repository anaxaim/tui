package repository

import (
	"context"
	"errors"

	"github.com/anaxaim/tui/backend/pkg/database"
)

var ErrConvertToHex = errors.New("failed to convert objectid to hex")

func NewRepository(db *database.MongoDB) Repository {
	r := &repository{
		db:         db,
		user:       newUserRepository(db),
		module:     newModuleRepository(db),
		credential: newCredentialRepository(db),
	}

	return r
}

type repository struct {
	db         *database.MongoDB
	user       UserRepository
	module     ModuleRepository
	credential CredentialRepository
}

func (r *repository) User() UserRepository {
	return r.user
}

func (r *repository) Module() ModuleRepository {
	return r.module
}

func (r *repository) Credential() CredentialRepository {
	return r.credential
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
