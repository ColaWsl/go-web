package repository

import (
	"errors"
	"go-web/model"
)

var users = []model.User{
	{ID: "1", Name: "Alice", NickName: "爱丽丝", Email: "alice@example.com"},
	{ID: "2", Name: "Bob", NickName: "鲍勃", Email: "bob@example.com"},
	{ID: "3", Name: "Charlie", NickName: "克瑞儿", Email: "charlie@example.com"},
}

func FindAllUsers() ([]model.User, error) {
	return users, nil
}

func FindUserByID(id string) (model.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return model.User{}, errors.New("user not found")
}

func SaveUser(user *model.User) error {
	users = append(users, *user)
	return nil
}
