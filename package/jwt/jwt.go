package jwt_service

import (
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt"
	"math/rand"
	"time"
)

type JWTokenService struct {
	privateKey string
}

func NewTokenService(privateKey string) (*JWTokenService, error) {
	if privateKey == "" {
		return nil, errors.New("signing key is not set")
	}
	return &JWTokenService{privateKey}, nil
}

func (j *JWTokenService) NewAccessToken(userId, refId string, ttl time.Duration) (string, error) {
	claims := JwtClaims{
		UserId: userId,
		Id:     refId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(j.privateKey))
}

func (j *JWTokenService) NewRefreshToken() (string, error) {
	b := make([]byte, 36)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	if _, err := r.Read(b); err != nil {
		return "", nil
	}
	return hex.EncodeToString(b), nil
}
