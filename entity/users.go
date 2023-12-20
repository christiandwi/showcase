package entity

import (
	"github.com/christiandwi/showcase/constant"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID          int64  `gorm:"column:id;primary_key"`
	SecureId    string `gorm:"column:secure_id"`
	Email       string
	PhoneNumber string
	Password    string
}

func (Users) TableName() string {
	return constant.EntityUsers
}

func (u *Users) BeforeCreate(scope *gorm.DB) (err error) {
	u.SecureId = uuid.NewString()

	return nil
}
