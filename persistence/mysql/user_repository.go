package mysql

import (
	"errors"
	"github.com/jinho-yoo-jack/gin-tonic-tester/model/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

var ErrDuplicatedUserName = errors.New("duplicated user name")

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
