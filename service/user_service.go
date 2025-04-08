package service

import (
	"database/sql"
	"fmt"
	"mark99/model/entity"
	"mark99/repository"
	"mark99/utils"
)

type UserInfo struct {
	UserId   string
	Password string
	NickName string
	Role     int
}

type SignInDto struct {
	UserId   string
	Password string
}

type UserService struct {
	userRepository *repository.UserRepository
	jwtUtils       *utils.JwtUtils
}

func NewUserService(ur *repository.UserRepository, ju *utils.JwtUtils) *UserService {
	return &UserService{ur, ju}
}

func (s *UserService) SignUp(u UserInfo) *entity.Member {
	encryptedPassword, err := utils.EncryptPassword(u.Password)
	user := entity.Member{
		UserId:   u.UserId,
		Password: encryptedPassword,
		NickName: u.NickName,
		CartId:   sql.NullInt32{Valid: false},
		Role:     sql.NullInt32{Int32: 1, Valid: true}}
	savedUser, err := s.userRepository.CreateNewUser(&user)
	if err != nil {
		panic(err)
	}
	return savedUser
}

func (s *UserService) SignIn(signDto SignInDto) string {
	m, err := s.userRepository.FindById(signDto.UserId)
	if err != nil {
		panic(err)
	}
	if isValid := utils.VerifyPassword(m.Password, signDto.Password); !isValid {
		token, err := s.jwtUtils.GenerateToken(m.UserId)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		return token
	}
	return "NOT_VALID"
}
