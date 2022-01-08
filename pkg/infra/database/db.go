package database

import (
	"context"
	"time"

	"github.com/Met4tron/go-rest-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB    *mongo.Database
	err   error
	DBErr error
)

type Database struct {
	*mongo.Database
}

func InitDB() error {
	var db = DB

	env := config.GetConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client_options := options.Client().ApplyURI(env.Db.Uri)

	defer cancel()

	client, err := mongo.Connect(ctx, client_options)

	if err != nil {
		DBErr = err
		return err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		DBErr = err
		return err
	}

	db = client.Database(env.Db.Name)

	DB = db

	return nil
}

// Helper to get connection from database
func GetDB() *mongo.Database {
	return DB
}

// Helper to get collection
func GetCollection(collection string) *mongo.Collection {
	return DB.Collection(collection)
}
