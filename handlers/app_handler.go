package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllProducts(c echo.Context) error {
	return c.String(http.StatusOK, "Here is a list of all the users!")
}