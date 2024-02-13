package usecase

import (
	"context"
	"fmt"
	"test_app/internal/adapter/db"
	"test_app/internal/entity"
	jwt_service "test_app/package/jwt"
	"test_app/package/methods"
	"time"
)

type authUseCase struct {
	repo           db.RefreshTokenRepo
	ts             jwt_service.TokenService
	accessTokenTTL time.Duration
}

func NewAuthUseCase(repo db.RefreshTokenRepo, token *jwt_service.JWTokenService, accessTokenTTL time.Duration) *authUseCase {
	return &authUseCase{repo, token, accessTokenTTL}
}

func (uc *authUseCase) Create(ctx context.Context, userId string) (*entity.SignTokens, error) {
	refresh, err := uc.ts.NewRefreshToken()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	hashedRefresh, err := methods.HashToken(refresh)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	dto := entity.RefreshToken{
		UserId:       userId,
		RefreshToken: hashedRefresh,
		CreatedAt:    time.Now().UTC(),
	}
	refId, err := uc.repo.Create(ctx, &dto)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	access, err := uc.ts.NewAccessToken(userId, refId, uc.accessTokenTTL)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	res := entity.SignTokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}
	return &res, nil
}

func (uc *authUseCase) FindById(ctx context.Context, id string) (*entity.RefreshToken, error) {
	token, err := uc.repo.FindById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return token, nil
}

func (uc *authUseCase) Refresh(ctx context.Context, state *entity.RefreshToken, input *entity.ResetTokens) (*entity.SignTokens, error) {
	err := methods.CompareRefresh(state.RefreshToken, input.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("refersh not match")
	}

	tokens, err := uc.Create(ctx, state.UserId)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	err = uc.repo.Delete(ctx, state.Id)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	return tokens, nil
}
