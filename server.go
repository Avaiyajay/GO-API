package main

import (
	"fmt"
	"os"

	"github.com/Avaiyajay/Go-api/common"
	"github.com/Avaiyajay/Go-api/models"
	router "github.com/Avaiyajay/Go-api/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)



func main() {
	e := echo.New()
	// wd, _ := os.Getwd()
	// e.Static("/static", fmt.Sprintf(`%v/public`, wd))
	e.Static("/static", "./public")
 	godotenv.Load()

	common.InitDB()
	common.DB.AutoMigrate(&models.Product{})

	router.Router(e);

	PORT := os.Getenv("PORT");
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%v", PORT)))
}