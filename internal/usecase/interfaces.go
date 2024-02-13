package usecase

import (
	"context"
	"test_app/internal/entity"
)

type (
	RefreshToken interface {
		Create(ctx context.Context, userId string) (*entity.SignTokens, error)
		Refresh(ctx context.Context, refreshToken *entity.RefreshToken, input *entity.ResetTokens) (*entity.SignTokens, error)
		FindById(ctx context.Context, id string) (*entity.RefreshToken, error)
	}
)
