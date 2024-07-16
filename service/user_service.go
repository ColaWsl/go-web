package service

import (
	"go-web/model"
	"go-web/repository"
)

type User = model.User

func GetAllUsers() ([]User, error) {
	return repository.FindAllUsers()
}

func GetUserByID(id string) (User, error) {
	return repository.FindUserByID(id)
}

func CreateUser(user *User) error {
	return repository.SaveUser(user)
}
