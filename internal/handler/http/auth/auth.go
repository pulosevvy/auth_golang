package auth

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
	"net/http"
	"test_app/internal/entity"
	"test_app/internal/usecase"
	"test_app/middleware"
	"test_app/package/logger/logger/sl"
)

type authRouter struct {
	uc usecase.RefreshToken
	l  *slog.Logger
}

func NewAuthRouter(e *echo.Echo, log *slog.Logger, u usecase.RefreshToken, am *middleware.AuthMiddleware) {
	r := authRouter{uc: u, l: log}
	router := e.Group("/auth")

	router.GET("/:id", r.auth, middleware.GuidValidation)
	router.POST("/refresh", r.refresh, middleware.BodyValidate[entity.ResetTokens], am.TokenMiddleware)
}

func (r *authRouter) auth(c echo.Context) error {
	userId := c.Param("id")
	info, err := r.uc.Create(c.Request().Context(), userId)
	if err != nil {
		r.l.Error("authRouter.auth", sl.Err(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	cookie := new(http.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = info.RefreshToken
	cookie.HttpOnly = true
	cookie.MaxAge = 15 * 86400
	c.SetCookie(cookie)

	return c.JSON(http.StatusCreated, info)
}

func (r *authRouter) refresh(c echo.Context) error {
	token := c.Get("token").(*entity.RefreshToken)
	body := c.Get("body").(*entity.ResetTokens)

	tokens, err := r.uc.Refresh(c.Request().Context(), token, body)
	if err != nil {
		r.l.Error("authRouter.refresh", sl.Err(err))
		return echo.NewHTTPError(http.StatusConflict, err)
	}

	return c.JSON(http.StatusOK, tokens)
}
