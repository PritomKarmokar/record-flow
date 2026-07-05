package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func CorrelationID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {

			correlationID := c.Request().Header.Get("X-Correlation-ID")
			if correlationID == "" {
				correlationID = uuid.New().String()
			}

			requestID := c.Request().Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = uuid.New().String()
			}

			c.Response().Header().Set("X-Correlation-ID", correlationID)
			c.Response().Header().Set("X-Request-ID", requestID)

			c.Set("correlation_id", correlationID)
			c.Set("request_id", requestID)

			return next(c)
		}
	}
}
