package jwt_service

import (
	"time"
)

type TokenService interface {
	NewAccessToken(userId, refId string, ttl time.Duration) (string, error)
	NewRefreshToken() (string, error)
	//Parse(accessToken string) (string, error)
}
