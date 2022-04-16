package users

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	// Get returns the User with the specified User Id.
	Get(email string) (*User, error)
	// Create creates a new User in the storage.
	Create(user *User) (*User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Migration() {
	u.db.AutoMigrate(&User{})
}

func (u *UserRepository) Create(user *User) (*User, error) {
	zap.L().Debug("user.repo.create", zap.Reflect("user", user))

	if err := u.db.Preload("Role").Create(user).Error; err != nil {
		zap.L().Error("user.repo.Create failed to create user", zap.Error(err))
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) Get(email string) (*User, error) {
	zap.L().Debug("user.repo.getuser", zap.Reflect("email", email))

	var user = &User{}
	if result := u.db.Preload("Role").First(&user,  "email = ?", email); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
