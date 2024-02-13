package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"test_app/package/validator/validation"
)

func GuidValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		guid := c.Param("id")

		ok := validation.IsGuid(guid)
		if !ok {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid guid")
		}
		return next(c)
	}
}
