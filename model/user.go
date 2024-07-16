package model

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
}
