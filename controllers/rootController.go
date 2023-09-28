package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RedirectToDocs(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
}
