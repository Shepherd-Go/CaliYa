package main

import (
	"CaliYa/cmd/api/router"
	"CaliYa/cmd/providers"
	"CaliYa/config"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {

	container := providers.BuildContainer()

	if err := container.Invoke(func(router *router.Router, server *echo.Echo, config config.Config) {

		router.Init()

		server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", config.Server.Port)))

	}); err != nil {

		log.Fatal(err)
	}

}
