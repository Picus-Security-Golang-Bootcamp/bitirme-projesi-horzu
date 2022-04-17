package order

import (
	"time"

	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/domain/order/orderItem"
	"gorm.io/gorm"
)

var days14ToHours float64 = 14 * 24

type Order struct {
	Id         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	TotalPrice float64
	IsCanceled bool

	Items  []orderItem.OrderItem `gorm:"foreignKey:OrderId"`
	UserId string
}

func (u *Order) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}

func NewOrder(userID string, items []orderItem.OrderItem) *Order {
	// var totalPrice float64 = 0.0
	// for _, item := range items {
	// 	totalPrice += item.Product.Price
	// }
	return &Order{
		UserId:     userID,
		Items:      items,
		// TotalPrice: totalPrice,
		IsCanceled: false,
	}
}
func NewOrderItem(quantity uint, productID string) *orderItem.OrderItem {
	return &orderItem.OrderItem{
		Quantity:   quantity,
		ProductId:  productID,
	}
}
