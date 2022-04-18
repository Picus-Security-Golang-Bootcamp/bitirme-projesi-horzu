package users

import (
	"github.com/horzu/golang/cart-api/internal/api"
)

func responseToUser(a *api.UserCreateUserRequest) *User {
	return &User{
		Username: *a.Username,
		Name:        *a.Name,
		Surname: *a.Surname,
		Email: *a.Email,
		Password: *a.Password,
		Phone: *&a.Phone,
		RoleId: a.RoleID,
	}
}
