package service

import (
	"ginTonicProject/model"
	"ginTonicProject/repository"
)

type UserInfo struct {
	UserId   string
	Password string
	NickName string
	Role     int
}

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{userRepository: ur}
}

func (s *UserService) SignUp(u UserInfo) *model.User {
	user := model.User{UserId: u.UserId, Password: u.Password, NickName: u.NickName, Role: 1}
	savedUser, err := s.userRepository.CreateNewUser(&user)
	if err != nil {
		panic(err)
	}
	return savedUser
}
