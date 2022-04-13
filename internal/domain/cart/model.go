package cart

import (
	"os/user"
	"time"

	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/domain/cartItem"
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
	User  *user.User
}

func NewCart(uid string) *Cart {
	return &Cart{
		UserID: uid,
	}
}

func (u *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}
