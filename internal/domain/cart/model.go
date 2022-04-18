package cart

import (
	"time"

	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/domain/cart/cartItem"
	"gorm.io/gorm"
)

var (
	MaxAllowedForBasket             = 20
	MaxAllowedQtyPerProduct         = 9
	MinCartAmountForOrder   float64 = 10
)


type Cart struct {
	Id            string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	UserID        string

	Items *[]cartItem.CartItem `gorm:"foreignKey:CartId"`
}

func NewCart(userId string) *Cart {
	return &Cart{
		UserID: userId,
	}
}

func (u *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}
