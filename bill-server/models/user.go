package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	UserID        uint           `gorm:"primaryKey" json:"user_id"`
	Username      string         `gorm:"size:50;not null" json:"username"`
	Email         string         `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password      string         `gorm:"size:100;not null" json:"-"`
	Phone         string         `gorm:"size:20" json:"phone"`
	Avatar        string         `gorm:"size:255" json:"avatar"`
	EmailVerified bool           `gorm:"default:false" json:"email_verified"`
	PhoneVerified bool           `gorm:"default:false" json:"phone_verified"`
	DefaultBookID uint           `json:"default_book_id"`
	Status        string         `gorm:"size:20;default:active" json:"status"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	LastLogin     time.Time      `json:"last_login"`

	// 关联关系
	DefaultBook *Book `gorm:"foreignKey:DefaultBookID" json:"-"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}
