package stock

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Stock struct {
	Id           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	ProductId    int
	QuantityHold int


}

func (u *Stock) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}
