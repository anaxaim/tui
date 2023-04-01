package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/anaxaim/tui/server/pkg/database"
	"github.com/anaxaim/tui/server/pkg/model"
)

type registryRepository struct {
	collection *mongo.Collection
}

func newRegistryRepository(db *database.MongoDB) RegistryRepository {
	registriesCollection := db.Client.Database(db.DBName).Collection("registries")

	return &registryRepository{
		collection: registriesCollection,
	}
}

func (r *registryRepository) Save(registry *model.RegistryContent) (*model.RegistryContent, error) {
	result, err := r.collection.InsertOne(context.Background(), registry)
	if err != nil {
		return nil, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		registry.ID = oid
		return registry, nil
	}

	return nil, ErrConvertToHex
}

func (r *registryRepository) Get(id string) (*model.RegistryContent, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	registry := new(model.RegistryContent)

	result := r.collection.FindOne(context.Background(), bson.M{"_id": oid})
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return registry, mongo.ErrNoDocuments
		}

		return nil, result.Err()
	}

	if err := result.Decode(&registry); err != nil {
		return nil, err
	}

	return registry, nil
}

func (r *registryRepository) Migrate() error {
	return nil
}
