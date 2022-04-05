package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	// Order_Cart []Product `json:"order_list" gorm:"foreignkey:id;references:AuthorID"`
	Ordered_At time.Time `json:"ordered_at"`
	Price int `json:"total_price"`
	Discount int64 `json:"discount"`
}

func (Order) TableName() string{
	// default table name
	return "orders"
}