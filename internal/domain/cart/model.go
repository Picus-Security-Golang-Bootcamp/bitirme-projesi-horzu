package cart

import (
	"time"

	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/domain/cart/cartItem"
	"gorm.io/gorm"
)

type Cart struct {
	Id         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	UserID     string
	TotalPrice float64
	Status     string

	Items *[]cartItem.CartItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

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
