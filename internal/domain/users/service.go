package users

import (
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo Repository
}

type Service interface {
	Create(user *User) error 
	LoginCheck(email ,password string) (*User, error) 
}

func NewUserService(repo Repository) Service {
	if repo == nil {
		return nil
	}

	return &userService{repo: repo}
}

func (u *userService) Create(user *User) error {
	_, err := u.repo.Create(user)
	if err != nil {
		return err
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
