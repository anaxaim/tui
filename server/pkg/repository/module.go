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

type moduleRepository struct {
	collection *mongo.Collection
}

func newModuleRepository(db *database.MongoDB) ModuleRepository {
	modulesCollection := db.Client.Database(db.DBName).Collection("modules")

	return &moduleRepository{
		collection: modulesCollection,
	}
}

func (r *moduleRepository) GetModuleByID(id string) (*model.TerraformModule, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	module := new(model.TerraformModule)

	result := r.collection.FindOne(context.Background(), bson.M{"_id": oid})
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return module, mongo.ErrNoDocuments
		}

		return nil, result.Err()
	}

	if err := result.Decode(&module); err != nil {
		return nil, err
	}

	return module, nil
}

func (r *moduleRepository) List() (model.TerraformModules, error) {
	modules := make(model.TerraformModules, 0)

	cursor, _ := r.collection.Find(context.Background(), bson.M{})
	if cursor.Err() != nil {
		return nil, cursor.Err()
	}

	if err := cursor.All(context.Background(), &modules); err != nil {
		return nil, err
	}

	return modules, nil
}

func (r *moduleRepository) Create(module *model.TerraformModule) (*model.TerraformModule, error) {
	result, err := r.collection.InsertOne(context.Background(), module)
	if err != nil {
		return nil, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		module.ID = oid

		return module, nil
	}

	return nil, ErrConvertToHex
}

func (r *moduleRepository) Update(module *model.TerraformModule) (*model.TerraformModule, error) {
	moduleBytes, err := bson.Marshal(module)
	if err != nil {
		return nil, err
	}

	var updateModuleObj bson.M

	err = bson.Unmarshal(moduleBytes, &updateModuleObj)
	if err != nil {
		return nil, err
	}

	delete(updateModuleObj, "_id")

	update := bson.M{
		"$set": updateModuleObj,
	}

	result, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": module.ID}, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return module, mongo.ErrNoDocuments
		}

		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return module, nil
}

func (r *moduleRepository) Delete(module *model.TerraformModule) error {
	result, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": module.ID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (r *moduleRepository) Migrate() error {
	return nil
}
