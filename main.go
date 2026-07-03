package main

import (
	"github.com/PritomKarmokar/record-flow/cmd/config"
	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

	config.LoadEnv()
	config.LoggerConfig()
	
	config.StartServer(e)
}
