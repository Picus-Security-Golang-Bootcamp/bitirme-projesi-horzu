package category

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	Id          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string         `gorm:"unique"`
	Description string
	IsActive    bool
}

func (u *Category) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}

func NewCategory(name string, desc string) *Category {
	return &Category{
		Name:        name,
		Description: desc,
		IsActive:    true,
	}
}
