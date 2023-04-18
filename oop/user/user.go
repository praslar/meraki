package user


type User struct {
	Email string
	Password string
	Birthday string
	InitBalance float64 // 100
}

func (u User) CreateUser(email string, birthday string) User {
	newUser := User{
		Email:       email,
		Birthday:    birthday,
		Password:    "admin123",
		InitBalance: 1000,
	}
	return newUser
}

func (u User) Login(email string, password string) error {
	return nil
}