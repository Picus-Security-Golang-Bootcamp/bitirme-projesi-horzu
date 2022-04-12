package orderProduct

import (
	"time"

	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/domain/product"
	"gorm.io/gorm"
)

type OrderProduct struct {
	Id         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	OrderId    int
	ProductId  int
	Quantity   int
	BasePrice  float64
	TotalPrice float64

	Product *product.Product
}

func (u *OrderProduct) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}

func (op *OrderProduct) SetProduct(product *product.Product) {
	op.Product = product
}

func (op *OrderProduct) GetProduct() *product.Product {
	return op.Product
}
