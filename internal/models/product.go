package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model   `bson:"id"`
	Product_Name *string `json:"product_name"`
	Price        *uint   `json:"price"`
	Rating       *uint   `json:"rating"`
	Image        *string `json:"image"`
	Category     Category
}

func (Product) TableName() string{
	// default table name
	return "products"
}