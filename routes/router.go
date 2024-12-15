package router

import (
	handler "github.com/Avaiyajay/Go-api/handlers"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo)  {
	e.GET("/", handler.GetAllProducts)
}

