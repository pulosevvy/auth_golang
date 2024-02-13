package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"test_app/config"
	"test_app/internal/adapter/db/mongo-repo/collections"
	"test_app/package/database/mongodb/indexes"
)

type MongoClient struct {
	Client *mongo.Client
}

func New(cfg *config.Config) (*MongoClient, error) {
	const fn = "mongo.New"

	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		cfg.Mongo.Username, cfg.Mongo.Password, cfg.Mongo.Host, cfg.Mongo.Port,
	)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("%s:%s", fn, err)
	}
	dbName := client.Database(cfg.Mongo.DBName)
	collection := dbName.Collection(collections.RefreshTokenModelConst)
	err = indexes.CreateRefreshTTL(collection)
	if err != nil {
		return nil, err
	}
	return &MongoClient{Client: client}, nil
}
