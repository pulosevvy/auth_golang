package db

import (
	"context"
	"test_app/internal/entity"
)

type (
	RefreshTokenRepo interface {
		Create(ctx context.Context, dto *entity.RefreshToken) (string, error)
		FindById(ctx context.Context, id string) (*entity.RefreshToken, error)
		Delete(ctx context.Context, id string) error
	}
)
