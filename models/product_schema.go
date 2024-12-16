package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string `json:"product_name" gorm:"column:product_name;not null;type:varchar(50)"`
    Price       int    `json:"price" gorm:"type:int;not null"`
    Quantity    int    `json:"quantity" gorm:"type:int;not null"`
    ImagePath   string `json:"image_path" gorm:"type:varchar(255)"`
}
