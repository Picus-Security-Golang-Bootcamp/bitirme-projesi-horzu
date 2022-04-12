package cart

import (
	"time"

	"github.com/google/uuid"
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

	Items *[]CartItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User  *User
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
