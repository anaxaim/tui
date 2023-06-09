package database

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/anaxaim/tui/backend/pkg/config"
)

var ErrFailedCloseMigration = errors.New("failed during close migration")

type MongoDB struct {
	*mongo.Client
	DBName string
}

func NewMongoClient(conf *config.DBConfig) (*MongoDB, error) {
	mongoHost := conf.Host
	if os.Getenv("MONGO_HOST") != "" {
		mongoHost = os.Getenv("MONGO_HOST")
	}

	if mongoHost == "" {
		mongoHost = "127.0.0.1"
	}

	uri := fmt.Sprintf("mongodb://%s/%s", net.JoinHostPort(mongoHost, conf.Port), conf.Database)
	if err := CreateDBUser(conf.MigrationsPath, uri); err != nil {
		return &MongoDB{}, err
	}

	newUserURI := fmt.Sprintf("mongodb://%s:%s@%s/%s", conf.User, conf.Password, net.JoinHostPort(mongoHost, conf.Port), conf.Database)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(newUserURI))
	if err != nil {
		return &MongoDB{}, err
	}

	return &MongoDB{Client: client, DBName: conf.Database}, nil
}

func CreateDBUser(migrations string, uri string) error {
	m, err := migrate.New(fmt.Sprintf("file://%s", migrations), uri)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	sourceErr, databaseErr := m.Close()
	if sourceErr != nil || databaseErr != nil {
		return fmt.Errorf("%w", ErrFailedCloseMigration)
	}

	return nil
}
