package orderItem

import (
	"time"

	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/domain/cart/cartItem"
	"github.com/horzu/golang/cart-api/internal/domain/product"
	"gorm.io/gorm"
)

type OrderItem struct {
	Id         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	OrderId    string
	Quantity   int
	IsCanceled bool
	
	ProductId  string
	Product *product.Product
}

func (u *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}

func NewOrderItem(cartItem cartItem.CartItem) *OrderItem {
	item := &OrderItem{
		ProductId: cartItem.ProductId,
		Quantity:  int(cartItem.Quantity),
	}
	return item
}