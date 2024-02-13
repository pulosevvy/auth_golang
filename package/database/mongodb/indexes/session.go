package indexes

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateRefreshTTL(collection *mongo.Collection) error {
	index := mongo.IndexModel{Keys: bson.M{
		"created_at": 1,
	},
		Options: options.Index().SetExpireAfterSeconds(15 * 24 * 60 * 60).SetName("created_at"),
	}

	_, err := collection.Indexes().CreateOne(context.Background(), index)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}
