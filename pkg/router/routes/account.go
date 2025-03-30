package routes

import (
	"Fank/internal/model/account"
	"github.com/labstack/echo/v4"
)

func RegisterAccountRoutes(r ...*echo.Group) {
	// api v1 group
	apiV1 := r[0]
	accountGroupV1 := apiV1.Group("/account")
	accountGroupV1.POST("/registerAccount", account.RegisterAcc)
}
