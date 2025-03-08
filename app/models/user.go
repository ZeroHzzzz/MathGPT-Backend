package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`                     // 用户唯一标识
	Username  string    `gorm:"unique;not null;default:'momo'"` // 用户名
	Email     string    `gorm:"type:varchar(100)"`              // 邮箱
	Phone     string    `gorm:"unique;not null"`                // 手机号
	Password  string    `gorm:"not null"`                       // 密码
	AvatarUrl string    `gorm:"type:varchar(100)"`              // 头像地址
	CreatedAt time.Time // 创建时间
}

type Claims struct {
	UserID uint `json:"userID"`
	jwt.RegisteredClaims
}
