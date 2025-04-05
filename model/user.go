package model

import (
	"ginTonicProject/config"
	"time"
)

type User struct {
	UserId    string    `gorm:"column:user_id;primary_key;auto_increment"`
	Password  string    `gorm:"column:password;not null"`
	NickName  string    `gorm:"column:nick_name"`
	Role      int32     `gorm:"column:role;null"`
	CartId    int32     `gorm:"column:cart_id;null"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u *User) Save() (*User, error) {
	var err error

	err = config.DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
