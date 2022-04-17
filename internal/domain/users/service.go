package users

import (
	"context"

	"github.com/horzu/golang/cart-api/internal/domain/cart"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo Repository
	cartService cart.Service
}

type Service interface {
	Create(ctx context.Context,user *User) error 
	LoginCheck(email ,password string) (*User, error) 
	GetUserId(email string) *User
}

func NewUserService(repo Repository, cartService cart.Service) Service {
	if repo == nil {
		return nil
	}

	return &userService{repo: repo, cartService: cartService}
}

func (u *userService) Create(ctx context.Context, user *User) error {
	user, err := u.repo.Create(user)
	if err != nil {
		return err
	}

	u.cartService.Create(ctx, user.Id)

	return nil
}

func (u *userService) GetUserId(email string) *User {
	_, err := u.repo.Get(email)
	if err != nil {
		return nil
	}

	return nil
}


func (u *userService) LoginCheck(email string,password string) (*User, error) {
	user, err := u.repo.Get(email)
	if err != nil {
		return nil, err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, err
	}

	return user, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
