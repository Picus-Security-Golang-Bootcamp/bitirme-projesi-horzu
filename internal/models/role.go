package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Role      string
}
