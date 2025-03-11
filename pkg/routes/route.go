package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func Handler() {
	echoHandler := echo.New()

	echoHandler.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Welcome to Library")
	})

	echoHandler.Logger.Fatal(echoHandler.Start(":1313"))
}
