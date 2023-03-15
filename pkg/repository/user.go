package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/anaxaim/tui/pkg/database"
	"github.com/anaxaim/tui/pkg/model"
)

type userRepository struct {
	collection *mongo.Collection
}

func newUserRepository(db *database.MongoDB) UserRepository {
	usersCollection := db.Client.Database(db.DBName).Collection("users")

	return &userRepository{
		collection: usersCollection,
	}
}

func (r *userRepository) GetUserByUsername(username string) (*model.User, error) {
	user := new(model.User)

	result := r.collection.FindOne(context.Background(), bson.M{"username": username})
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return user, mongo.ErrNoDocuments
		}
		return nil, fmt.Errorf("failed to FindOne user by username: %s due to error: %w", username, result.Err())
	}

	if err := result.Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode user (username:%s) from DB due to error: %w", username, err)
	}

	return user, nil
}

func (r *userRepository) List() (model.Users, error) {
	users := make(model.Users, 0)

	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if cursor.Err() != nil {
		return nil, fmt.Errorf("failed to find all users due to error: %w", err)
	}

	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, fmt.Errorf("failed to read all users from cursor. error: %w", err)
	}

	return users, nil
}

func (r *userRepository) Delete(user *model.User) error {
	result, err := r.collection.DeleteOne(context.Background(), bson.M{"username": user.Username})
	if err != nil {
		return fmt.Errorf("failed to execute query: error: %w", err)
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (r *userRepository) Create(user *model.User) (*model.User, error) {
	if _, err := r.collection.InsertOne(context.Background(), user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(user *model.User) (*model.User, error) {
	userBytes, err := bson.Marshal(user)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user. error: %w", err)
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user bytes. error: %w", err)
	}
	delete(updateUserObj, "_id")

	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": updateUserObj}

	result, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, mongo.ErrNoDocuments
		}
		return nil, fmt.Errorf("failed to execute update user query. error: %w", err)
	}

	if result.MatchedCount == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return user, nil
}

func (r *userRepository) Migrate() error {
	return nil
}
