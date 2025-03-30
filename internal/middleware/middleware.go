package middleware

import "github.com/labstack/echo/v4"
import requestMiddleware "github.com/labstack/echo/v4/middleware"

func InitMiddleware(app *echo.Echo) {

	// 全局请求 ID 中间件
	app.Use(requestMiddleware.RequestID())
}
