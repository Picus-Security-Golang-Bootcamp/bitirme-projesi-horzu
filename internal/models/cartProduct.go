package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartProduct struct {
	Id         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	CartId     int
	ProductId  int
	Quantity   int
	TotalPrice float64

	Product *Product
	Cart    *Cart
}

func (u *CartProduct) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}
