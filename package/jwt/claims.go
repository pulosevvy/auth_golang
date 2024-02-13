package jwt_service

import (
	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	jwt.StandardClaims
}
