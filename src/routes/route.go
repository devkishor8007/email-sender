package routes

import (
	"net/http"

	"github.com/devkishor8007/email-sender/src/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Secure())

	api := e.Group("/api/v1")

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	api.POST("/email-generates", controllers.CreateEmailTemplates)
	api.GET("/email-generates", controllers.GetEmailTemplates)

	return e
}
