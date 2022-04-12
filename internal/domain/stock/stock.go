package stock

import (
	"time"

	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/domain/product"
	"gorm.io/gorm"
)

type Stock struct {
	Id           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	ProductId    int
	QuantityHold int

	Product *product.Product
}

func (u *Stock) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}

func (s *Stock) SetProduct(product *product.Product) {
	s.Product = product
}

func (s *Stock) GetProduct() *product.Product {
	return s.Product
}