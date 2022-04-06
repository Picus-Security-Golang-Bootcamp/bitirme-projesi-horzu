package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model   `bson:"id"`
	Product_Name string `json:"product_name"`
	Price        int64   `json:"price"`
	Rating       int64   `json:"rating"`
	Image        string `json:"image"`
	// Category     Category
}

func (Product) TableName() string{
	// default table name
	return "products"
}