package mediaFile

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MediaFile struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Filename  string
	Url       string
}

func (u *MediaFile) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}
