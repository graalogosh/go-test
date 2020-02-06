package configs

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-test/internals/handlers"
)

func SetupHttpRouter() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", handlers.RootHandler)

	e.Logger.Fatal(e.Start(":80"))
}
