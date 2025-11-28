package model

import (
	"time"
)

// Login 登录记录模型
type Login struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	IP        string    `gorm:"size:128;not null" json:"ip"`
	UA        string    `gorm:"size:128;not null" json:"ua"`
	Token     string    `gorm:"size:128;not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (Login) TableName() string {
	return "tb_logins"
}
