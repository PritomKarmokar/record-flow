package route

import (
	"net/http"

	"github.com/PritomKarmokar/record-flow/cmd/config"
	"github.com/labstack/echo/v5"
)

func RegisterServiceRoutes(route *echo.Group) {
	// Liveness probe - always returns OK unless process crashed
	route.GET("/live", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "alive",
		})
	})

	// Readiness probe - checks dependencies
	route.GET("ready", func(c *echo.Context) error {
		logger := config.GetRequestLogger(c)
		status := "ready"

		// Check database
		db := config.GetDatabase()
		if err := db.Exec("SELECT 1").Error; err != nil {
			logger.Error().Err(err).Msg("Database check failed")
			// Database is critical
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"status": "unavailable",
				"reason": "database_unavailable",
			})
		}

		response := map[string]interface{}{
			"status": status,
		}
		return c.JSON(http.StatusOK, response)
	})

}
