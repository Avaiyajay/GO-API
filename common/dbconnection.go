package common

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	godotenv.Load("../.env")
	USERNAME := os.Getenv("USERNAME");
	PASSWORD := os.Getenv("PASSWORD");
	HOST := os.Getenv("HOST");
	DBPORT := os.Getenv("DBPORT");
	DBNAME := os.Getenv("DBNAME");
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, HOST, DBPORT, DBNAME) 
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to Establish Database Connection!")
	}

	log.Println("Database Conection Successful!")
	DB = gormDB;
}