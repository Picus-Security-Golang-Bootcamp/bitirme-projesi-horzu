package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderProduct struct {
	Id          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	OrderId    int
	ProductId  int
	Quantity   int
	BasePrice  float64
	TotalPrice float64

	Product *Product
}

func (u *OrderProduct) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}

func (op *OrderProduct) SetProduct(product *Product) {
	op.Product = product
}

func (op *OrderProduct) GetProduct() *Product {
	return op.Product
}