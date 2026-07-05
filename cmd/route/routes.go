package route

import (
	"github.com/PritomKarmokar/record-flow/cmd/middleware"
	"github.com/labstack/echo/v5"
)

func RegisterRoutes(e *echo.Echo) {
	// Apply correlation ID tracking globally (first middleware for distributed tracing)
	e.Use(middleware.CorrelationID())

	// Apply security headers globally
	e.Use(middleware.SecurityHeaders())
}
