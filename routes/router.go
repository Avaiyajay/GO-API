package router

import (
	handler "github.com/Avaiyajay/Go-api/handlers"
	"github.com/Avaiyajay/Go-api/middleware"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo)  {
	e.GET("/", handler.GetAllProducts)
	e.POST("/add-product", middleware.ProductRequestValidation(handler.AddNewProduct))
	e.PUT("/update-product", handler.UpdateProduct)
	e.DELETE("/delete-product", handler.DeleteProduct)
	e.POST("/upload", handler.FileUpload)
}

