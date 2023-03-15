package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/anaxaim/tui/pkg/database"
	"github.com/anaxaim/tui/pkg/model"
)

var ErrConvertToHex = errors.New("failed to convert objectid to hex")

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
	result, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		user.ID = oid
		return user, nil
	}

	return nil, ErrConvertToHex
}

func (r *userRepository) Update(user *model.User) (*model.User, error) {
	update := bson.M{
		"$set": bson.M{
			"username": user.Username,
			"password": user.Password,
		},
	}

	result, err := r.collection.UpdateByID(context.Background(), user.ID, update)
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

func (r *userRepository) Exists(username string) (bool, error) {
	count, err := r.collection.CountDocuments(context.Background(), bson.M{"username": username})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *userRepository) Migrate() error {
	return nil
}
