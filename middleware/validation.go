package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Avaiyajay/Go-api/dto"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Field string `json:"field"`
}

func ProductRequestValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
		var reqData dto.ProductDTO;
		var bodyBytes, _ = io.ReadAll(c.Request().Body)
		bindErr := json.Unmarshal(bodyBytes, &reqData)
		// bindErr := c.Bind(&reqData)
		if bindErr != nil {
			log.Fatalf("Failed to bind request with DTO.")
		}

		c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		validate := validator.New(validator.WithRequiredStructEnabled())
		err := validate.Struct(&reqData)
		var response ErrorResponse
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				log.Fatalf(`Invalid Validation Error: %v`, err)
			}
			
			for _, err := range err.(validator.ValidationErrors) {
				switch err.ActualTag(){
					case "required": {
						response = ErrorResponse{
							Field: err.StructField(),
							Message: err.StructField() + " is required",
						}
					}
				}
			}

			return c.JSON(http.StatusBadRequest, &response)
		}
		return next(c)
	}
}