package users

import (
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/domain/users/role"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string         `gorm:"size:255;not null;unique" json:"email"`
	Password  string         `gorm:"size:255;not null;" json:"password"`
	RoleId    int64          `gorm:"not null; default:2" json:"role_id"`
	Role      role.Role
}

// BeforeSave hashes password before saving to database
func (u *User) BeforeSave(tx *gorm.DB) error {
	u.Id = uuid.New().String()
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))

	return nil

}
