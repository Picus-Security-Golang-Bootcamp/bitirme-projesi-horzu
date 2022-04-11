package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Stock struct {
	Id           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	ProductId    int
	QuantityHold int

	Product *Product
}

func (u *Stock) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}

func (s *Stock) SetProduct(product *Product) {
	s.Product = product
}

func (s *Stock) GetProduct() *Product {
	return s.Product
}