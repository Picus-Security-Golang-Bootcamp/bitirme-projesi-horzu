package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	Id          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	UserId      string
	IsActive    bool
	OrderNumber string
	TotalPrice  float64
	Status      string

	Products   *[]CartProduct
	User       *User
}

func (u *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}