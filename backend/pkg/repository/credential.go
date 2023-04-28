package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/anaxaim/tui/backend/pkg/database"
	"github.com/anaxaim/tui/backend/pkg/model"
)

type credentialRepository struct {
	collection *mongo.Collection
}

func newCredentialRepository(db *database.MongoDB) CredentialRepository {
	credentialsCollection := db.Client.Database(db.DBName).Collection("credentials")

	return &credentialRepository{
		collection: credentialsCollection,
	}
}

func (c *credentialRepository) GetCredentialByID(id string) (*model.Credential, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	credential := new(model.Credential)

	result := c.collection.FindOne(context.Background(), bson.M{"_id": oid})
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return credential, mongo.ErrNoDocuments
		}

		return nil, result.Err()
	}

	if err := result.Decode(&credential); err != nil {
		return nil, err
	}

	return credential, nil
}

func (c *credentialRepository) GetCredentialByName(name string) (*model.Credential, error) {
	credential := new(model.Credential)

	result := c.collection.FindOne(context.Background(), bson.M{"name": name})
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return credential, mongo.ErrNoDocuments
		}

		return nil, result.Err()
	}

	if err := result.Decode(&credential); err != nil {
		return nil, err
	}

	return credential, nil
}

func (c *credentialRepository) List() (model.Credentials, error) {
	credentials := make(model.Credentials, 0)

	cursor, _ := c.collection.Find(context.Background(), bson.M{})
	if cursor.Err() != nil {
		return nil, cursor.Err()
	}

	if err := cursor.All(context.Background(), &credentials); err != nil {
		return nil, err
	}

	return credentials, nil
}

func (c *credentialRepository) Create(credential *model.Credential) (*model.Credential, error) {
	result, err := c.collection.InsertOne(context.Background(), credential)
	if err != nil {
		return nil, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		credential.ID = oid

		return credential, nil
	}

	return nil, ErrConvertToHex
}

func (c *credentialRepository) Delete(credential *model.Credential) error {
	result, err := c.collection.DeleteOne(context.Background(), bson.M{"_id": credential.ID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (c *credentialRepository) Migrate() error {
	return nil
}
