package user

import (
	"github.com/horzu/golang/cart-api/internal/models"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Migration() {
	u.db.AutoMigrate(&models.User{})
}

func (u *UserRepository) GetUserList() (*[]models.User, error) {
	zap.L().Debug("user.repo.getAll")

	var ul = &[]models.User{}
	if err := u.db.Preload("User").Find(&ul).Error; err != nil {
		zap.L().Error("user.repo.getAll failed to get users", zap.Error(err))
		return nil, err
	}

	return ul, nil
}

func (u *UserRepository) GetUser(email string) (*models.User, error) {
	zap.L().Debug("user.repo.getuser", zap.Reflect("email", email))

	var user = &models.User{}
	if result := u.db.First(&user, "email = ?", email); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (p *UserRepository) LoginCheck(email, password string) (*models.User, error) {
	zap.L().Debug("user.repo.LoginCheck", zap.Reflect("email", email))

	var user = &models.User{}
	if result := p.db.Preload("Role").First(&user,  "email = ?", email); result.Error != nil {
		return nil, result.Error
	}

	err := VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}

	return user, nil
}

func (u *UserRepository) SaveUser(b *models.User) (*models.User, error) {
	zap.L().Debug("user.repo.create", zap.Reflect("user", b))

	if err := u.db.Preload("Role").Create(b).Error; err != nil {
		zap.L().Error("user.repo.Create failed to create user", zap.Error(err))
		return nil, err
	}

	return b, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
