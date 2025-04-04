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

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) SignUp(u UserInfo) *model.User {

	user := model.User{UserId: u.UserId, Password: u.Password, NickName: u.NickName, Role: 1}

	savedUser, err := user.Save()
	if err != nil {
		panic(err)
	}

	return savedUser
}
