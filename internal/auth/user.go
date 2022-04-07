package auth

type User struct {
	Id       int
	Email    string
	Password string
	Roles    []string
}

func GetUserList() []*User {
	return []*User{
		{
			Id:       1,
			Email:    "admin@mail.com",
			Password: "1234admin",
			Roles:    []string{"admin", "customer"},
		},
		{
			Id:       2,
			Email:    "customer1@mail.com",
			Password: "customer1",
			Roles:    []string{"customer"},
		},
		{
			Id:       3,
			Email:    "customer2@mail.com",
			Password: "customer2",
			Roles:    []string{"customer"},
		},
	}
}

func GetUser(email, password string) *User {
	for _, v := range GetUserList() {
		if v.Email == email && v.Password == password {
			return v
		}

	}
	return nil
}