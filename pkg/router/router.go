package router

import (
	"Fank/pkg/router/routes"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(app *echo.Echo) {
	api1 := app.Group("/api/v1")

	routes.RegisterAccountRoutes(api1)
}
