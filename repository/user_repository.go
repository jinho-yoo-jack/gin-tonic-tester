package repository

import (
	"gorm.io/gorm"
	"mark99/model/entity"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateNewUser(m *entity.Member) (*entity.Member, error) {
	err := r.DB.Create(&m).Error
	return m, err
}

func (r *UserRepository) FindById(userId string) (*entity.Member, error) {
	var m *entity.Member
	err := r.DB.Find(&m, "user_id = ?", userId).Error
	//err := r.DB.First(&m, userId).Error
	return m, err
}
