package mongo_repo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"test_app/internal/adapter/db/mongo-repo/collections"
	"test_app/internal/entity"
)

type refreshTokenRepo struct {
	db *mongo.Collection
}

func NewRefreshTokenRepo(db *mongo.Database) *refreshTokenRepo {
	collection := db.Collection(collections.RefreshTokenModelConst)
	return &refreshTokenRepo{db: collection}
}

func (ur *refreshTokenRepo) Create(ctx context.Context, dto *entity.RefreshToken) (string, error) {
	created, err := ur.db.InsertOne(ctx, dto)
	if err != nil {
		return "", err
	}
	id, _ := created.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (ur *refreshTokenRepo) FindById(ctx context.Context, id string) (*entity.RefreshToken, error) {
	var e entity.RefreshToken
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	if err := ur.db.FindOne(ctx, bson.M{"_id": objectId}).Decode(&e); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("sessoin not found")
		}
		return nil, fmt.Errorf("%s", err)
	}
	return &e, nil
}

func (ur *refreshTokenRepo) Delete(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	_, err = ur.db.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	return nil
}
