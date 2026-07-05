package main

import (
	"github.com/PritomKarmokar/record-flow/cmd/config"
	"github.com/labstack/echo/v5"
	"net/http"
)

func main() {
	e := echo.New()

	config.LoadEnv()
	config.LoggerConfig()
	config.EchoConfig(e)
	config.ConnectDB()

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
		})
	})

	config.StartServer(e)
}
