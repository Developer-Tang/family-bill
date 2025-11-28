package model

import (
	"time"
)

// User 用户模型
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:128;not null" json:"username"`
	Password  string    `gorm:"size:128;not null" json:"-"`
	Email     string    `gorm:"uniqueIndex;size:128;not null" json:"email"`
	Avatar    string    `gorm:"type:text" json:"avatar"`
	HaveRole  string    `gorm:"size:128;default:'[]'" json:"have_role"`
	Status    int8      `gorm:"default:0" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (User) TableName() string {
	return "tb_users"
}
