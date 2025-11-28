package model

import (
	"time"
)

// VerifyCode 验证码模型
type VerifyCode struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Code      string    `gorm:"size:128;not null" json:"code"`
	Type      string    `gorm:"size:128;not null" json:"type"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (VerifyCode) TableName() string {
	return "tb_verify_codes"
}
