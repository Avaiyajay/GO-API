package dto

type ProductDTO struct {
	ProductName string `json:"product_name" validate:"required"`
	Price int `json:"price" validate:"required,gt=0"`
	Quantity int `json:"quantity" validate:"required,gt=0"`
	ImagePath string `json:"image_path"`
}