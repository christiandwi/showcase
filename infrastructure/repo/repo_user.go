package repo

import (
	"github.com/christiandwi/showcase/database"
	"github.com/christiandwi/showcase/domain/users"
	"github.com/christiandwi/showcase/entity"
)

type userRepo struct {
	db *database.Database
}

func SetupUserRepo(db *database.Database) users.UsersRepository {
	return &userRepo{db: db}
}

func (u userRepo) GetUser(userIdentifier string) (user entity.Users, err error) {
	err = u.db.
		Where("phone_number = ?", userIdentifier).
		Or("email = ?", userIdentifier).
		First(&user).Error

	return
}

func (u userRepo) CreateUser(email string, password string, phoneNumber string) error {
	user := entity.Users{
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
	}

	err := u.db.Create(&user).Error
	return err
}
