package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"test_app/internal/usecase"
	jwt_service "test_app/package/jwt"
)

type AuthMiddleware struct {
	uc     usecase.RefreshToken
	secret string
}

func NewAuthMiddleware(uc usecase.RefreshToken, secret string) *AuthMiddleware {
	return &AuthMiddleware{uc, secret}
}

func (au *AuthMiddleware) TokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tokenHeader string
		authorization := c.Request().Header.Get("Authorization")
		if strings.HasPrefix(authorization, "Bearer ") {
			tokenHeader = strings.TrimPrefix(authorization, "Bearer ")
		}
		if tokenHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		token, err := jwt.ParseWithClaims(tokenHeader, &jwt_service.JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("Unauthorized")
			}
			return []byte(au.secret), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		claims, ok := token.Claims.(*jwt_service.JwtClaims)
		if !ok || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}
		refresh, err := au.uc.FindById(c.Request().Context(), claims.Id)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		c.Set("token", refresh)

		return next(c)
	}
}
