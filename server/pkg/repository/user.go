package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/anaxaim/tui/server/pkg/database"
	"github.com/anaxaim/tui/server/pkg/model"
)

var ErrEmptyUserID = errors.New("empty user id")

type userRepository struct {
	collection *mongo.Collection
}

func newUserRepository(db *database.MongoDB) UserRepository {
	usersCollection := db.Client.Database(db.DBName).Collection("users")

	return &userRepository{
		collection: usersCollection,
	}
}

func (r *userRepository) GetUserByID(id string) (*model.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	user := new(model.User)

	result := r.collection.FindOne(context.Background(), bson.M{"_id": oid})
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return user, mongo.ErrNoDocuments
		}

		return nil, result.Err()
	}

	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByAuthID(authType, authID string) (*model.User, error) {
	authInfo := new(model.AuthInfo)

	result := r.collection.FindOne(context.Background(), bson.M{"authType": authType, "authId": authID})
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, mongo.ErrNoDocuments
		}

		return nil, result.Err()
	}

	if err := result.Decode(&authInfo); err != nil {
		return nil, err
	}

	return r.GetUserByID(authInfo.UserID)
}

func (r *userRepository) GetUserByName(name string) (*model.User, error) {
	user := new(model.User)

	result := r.collection.FindOne(context.Background(), bson.M{"name": name})
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return user, mongo.ErrNoDocuments
		}

		return nil, result.Err()
	}

	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) List() (model.Users, error) {
	users := make(model.Users, 0)

	cursor, _ := r.collection.Find(context.Background(), bson.M{})
	if cursor.Err() != nil {
		return nil, cursor.Err()
	}

	if err := cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Delete(user *model.User) error {
	result, err := r.collection.DeleteOne(context.Background(), bson.M{"name": user.Name})
	if err != nil {
		return err
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
			"name":     user.Name,
			"password": user.Password,
		},
	}

	result, err := r.collection.UpdateByID(context.Background(), user.ID, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user, mongo.ErrNoDocuments
		}

		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return user, nil
}

func (r *userRepository) AddAuthInfo(authInfo *model.AuthInfo) error {
	if authInfo == nil {
		return nil
	}

	if authInfo.UserID == "" {
		return fmt.Errorf("%w", ErrEmptyUserID)
	}

	result, err := r.collection.InsertOne(context.Background(), authInfo)
	if err != nil {
		return err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		authInfo.ID = oid

		return nil
	}

	return ErrConvertToHex
}

func (r *userRepository) DelAuthInfo(authInfo *model.AuthInfo) error {
	if authInfo == nil {
		return nil
	}

	result, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": authInfo.ID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (r *userRepository) Exists(name string) (bool, error) {
	count, err := r.collection.CountDocuments(context.Background(), bson.M{"name": name})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepository) Migrate() error {
	return nil
}
