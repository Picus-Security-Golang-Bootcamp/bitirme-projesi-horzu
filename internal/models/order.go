package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	Id          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	UserId		string
	OrderNumber string
	TotalPrice  float64
	Status      string

	Products []*OrderProduct
	User     *User
}

func (u *Order) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}


func (o *Order) AddProduct(product *OrderProduct) {
	o.Products = append(o.Products, product)
}

func (o *Order) RemoveProduct(product *OrderProduct) {
	for idx, op := range o.Products {
		if op.Id == product.Id {
			o.Products = append(o.Products[:idx], o.Products[idx+1:]...)
			break
		}
	}
}

func (o *Order) SetProducts(products []*OrderProduct) {
	o.Products = products
}

func (o *Order) GetProducts() []*OrderProduct {
	return o.Products
}

func (o *Order) SetUser(user *User) {
	o.User = user
}

func (o *Order) GetUser() *User {
	return o.User
}