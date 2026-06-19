package routes

import (
	"individual-project/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "Wonderlog API running",
		})
	})
	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)
}
