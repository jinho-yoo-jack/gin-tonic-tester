package repository

import (
	"ginTonicProject/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateNewUser(u *model.User) (*model.User, error) {
	err := r.DB.Create(&u).Error
	return u, err
}
