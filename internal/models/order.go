package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Order_Cart *[]Product `json:"order_list" bson:"order_list"`
	Ordered_At time.Time `json:"ordered_at" bson:"ordered_at"`
	Price int `json:"total_price" bson:"total_price"`
	Discount *int `json:"discount" bson:"discount"`
}

func (Order) TableName() string{
	// default table name
	return "orders"
}