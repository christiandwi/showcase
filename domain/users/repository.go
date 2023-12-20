package users

import "github.com/christiandwi/showcase/entity"

type UsersRepository interface {
	GetUser(userIdentifier string) (user entity.Users, err error)
	CreateUser(email string, password string, phoneNumber string) error
}
