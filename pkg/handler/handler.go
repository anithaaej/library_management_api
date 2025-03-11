package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitHello(c echo.Context) {
	c.String(http.StatusOK, "Hello, Welcome to Library")
}
