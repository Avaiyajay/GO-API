package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/Avaiyajay/Go-api/common"
	"github.com/Avaiyajay/Go-api/dto"
	"github.com/Avaiyajay/Go-api/models"
	"github.com/labstack/echo/v4"
)

func GetAllProducts(c echo.Context) error {
	var userList []*models.Product
	result := common.DB.Find(&userList)
	if result.Error != nil {
		return c.String(http.StatusNoContent, "No Data Found!")
	}
	response := map[string]interface{}{
		"data" : userList,
	}
	return c.JSON(http.StatusOK, response)
}

func AddNewProduct(c echo.Context) error {
	payload := new(dto.ProductDTO)
	var bodyBytes, _ = io.ReadAll(c.Request().Body)
	err := json.Unmarshal(bodyBytes, &payload)
	if err != nil {
		fmt.Println(err)
	}
	
	NewProduct := &models.Product{
		ProductName: payload.ProductName,
		Price: payload.Price,
		Quantity: payload.Quantity,
		ImagePath: "",
	}

	common.DB.Create(NewProduct)
	return c.String(http.StatusOK, "Record Created Successfully")
}

func UpdateProduct(c echo.Context) error {
	var id = c.QueryParam("id")
	intId, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.String(http.StatusBadRequest, "Failed to Parse QueryParam")
	}
	SelectedProduct := new(models.Product)
	reqData, _ := io.ReadAll(c.Request().Body)
	payload := new(models.Product)
	err := json.Unmarshal(reqData, &payload)
	if err != nil {
		return c.String(http.StatusBadRequest, "Please send currect data.")
	}

	result := common.DB.Find(&SelectedProduct, intId)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Unable to find product with provided id.")
	}


	SelectedProduct.ProductName = payload.ProductName
	SelectedProduct.Price = payload.Price
	SelectedProduct.Quantity = payload.Quantity

	result = common.DB.Save(&SelectedProduct)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Error saving data to database.")
	}
	return c.String(http.StatusOK, "Record Updated Successfully")
}

func DeleteProduct(c echo.Context) error {
	id := c.QueryParam("id");
	intId, parseErr := strconv.Atoi(id)
	if parseErr != nil {
		return c.String(http.StatusBadRequest, "Failed to parse QueryParam")
	}
	var product = new(models.Product)
	result := common.DB.Delete(&product, intId)
	if result.Error != nil {
		return c.String(http.StatusBadGateway, "Failed to find item in database!")
	}
	response := map[string]string{
		"message" : "Record Deleted Successfully!",
	}
	return c.JSON(http.StatusOK, response)
}

func FileUpload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "Please send file with proper field name");
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Unable to Open File")
	}
	defer src.Close()

	dest, err := os.Create(file.Filename)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Unable to Create File")
	}
	defer dest.Close()

	if _, err := io.Copy(dest, src); err != nil {
		return c.String(http.StatusInternalServerError, "Unable to copy file to destination")
	}

	return c.String(http.StatusOK, "File Upload Successful")

}