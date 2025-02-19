package service

import (
	"errors"

	"github.com/mutinho/src/model"
	"github.com/mutinho/src/repository"
	"github.com/mutinho/util"
)

func RegisterUser(name, email, password string) error {
	user := model.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := user.HashPassword(); err != nil {
		return err
	}

	return repository.CreateUser(&user)
}

func LoginUser(email, password string) (string, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := user.CheckPassword(password); err != nil {
		return "", errors.New("invalid credentials")
	}

	return util.GenerateToken(user.ID)
}
