package repository

import (
	"github.com/mutinho/config"
	"github.com/mutinho/src/model"
)

func CreateUser(user *model.User) (*model.User, error) {
	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
