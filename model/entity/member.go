package entity

import (
	"database/sql"
	"github.com/jinho-yoo-jack/gin-tonic-tester/service"
	"time"
)

type Member struct {
	UserId    string        `gorm:"column:user_id;primary_key;auto_increment"`
	Password  string        `gorm:"column:password;not null"`
	NickName  string        `gorm:"column:nick_name"`
	Role      sql.NullInt32 `gorm:"column:role;null"`
	CartId    sql.NullInt32 `gorm:"column:cart_id;null"`
	CreatedAt time.Time     `gorm:"column:created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at"`
}

func of(u service.UserInfo) *Member {
	return &Member{}
}
