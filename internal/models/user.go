package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"password"`
	IsAdmin bool `gorm:"not null;" json:"isAdmin"`
}

func (User) TableName() string{
	// default table name
	return "users"
}

// BeforeSave hashes password before saving to database
func (u *User) BeforeSave(tx *gorm.DB) error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username 
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil

}