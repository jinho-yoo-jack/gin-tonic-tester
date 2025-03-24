package service

import (
	"ginTonicProject/model"
)

type UserInfo struct {
	UserId   string
	Password string
	NickName string
	Role     int
}

func SignUp(u UserInfo) *model.User {

	user := model.User{}
	user.UserId = u.UserId
	user.Password = u.Password
	user.NickName = u.NickName
	user.Role = 1

	savedUser, err := user.Save()
	if err != nil {
		panic(err)
	}

	return savedUser
}
