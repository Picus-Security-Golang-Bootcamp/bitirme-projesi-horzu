package cartItem

import (
	"time"

	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/domain/product"
	"gorm.io/gorm"
)

type CartItem struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CartId    string
	ProductId string
	Quantity  uint

	Product *product.Product `gorm:"foreignKey:ProductId"`
}

func (u *CartItem) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}

func NewCartItem(productID string, cartID string, quantity uint) *CartItem{
	return &CartItem{
		ProductId: productID,
		Quantity: quantity,
		CartId: cartID,
	}
}