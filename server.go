package main

import (
	"fmt"
	"os"

	"github.com/Avaiyajay/Go-api/common"
	router "github.com/Avaiyajay/Go-api/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)



func main() {
	e := echo.New()
 	godotenv.Load()

	common.InitDB()
	
	router.Router(e);

	PORT := os.Getenv("PORT");
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%v", PORT)))
}