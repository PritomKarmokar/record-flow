package main

import (
	"github.com/PritomKarmokar/record-flow/cmd/config"
	"github.com/PritomKarmokar/record-flow/cmd/route"
	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

	config.LoadEnv()
	config.LoggerConfig()
	config.EchoConfig(e)
	config.ConnectDB()
	route.RegisterRoutes(e)
	config.StartServer(e)
}
