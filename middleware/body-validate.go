package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func BodyValidate[T any](next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input T
		if err := c.Bind(&input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		if err := c.Validate(&input); err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err)
		}
		c.Set("body", &input)
		return next(c)
	}
}
