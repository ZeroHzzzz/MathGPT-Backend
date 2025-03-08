package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID        string    `gorm:"primaryKey" json:"user_id"`                      // 用户唯一标识
	Username  string    `gorm:"unique;not null;default:'momo'" json:"username"` // 用户名
	Email     string    `gorm:"type:varchar(100)" json:"email"`                 // 邮箱
	Phone     string    `gorm:"unique;not null" json:"phone"`                   // 手机号
	Password  string    `gorm:"not null" json:"-"`                              // 密码
	AvatarUrl string    `gorm:"type:varchar(100)" json:"avatar_url"`            // 头像地址
	CreatedAt time.Time `json:"create_at"`                                      // 创建时间
}

type Claims struct {
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}
