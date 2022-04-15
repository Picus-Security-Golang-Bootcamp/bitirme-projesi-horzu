package order

import (
	"time"

	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/domain/cart"
	"github.com/horzu/golang/cart-api/internal/domain/order/orderItem"
	"gorm.io/gorm"
)

type Order struct {
	Id         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	TotalPrice float64
	IsCanceled bool

	Items  []*orderItem.OrderItem
	UserId string

}

func (u *Order) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}

func NewOrder(cart cart.Cart) *Order {

	order := &Order{
		TotalPrice: cart.TotalPrice,
		UserId:     cart.UserID,
		Items:      []*orderItem.OrderItem{},
	}

	for _, item := range *cart.Items {
		orderItem := orderItem.NewOrderItem(item)
		order.Items = append(order.Items, *&orderItem)
	}

	return order

}

func (o *Order) isCancelable() bool {
	orderDate := o.CreatedAt
	now := time.Now()
	if now.Sub(orderDate) < time.Hour*24*14 {
		return true
	}
	return false
}