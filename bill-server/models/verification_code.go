package models

import (
	"time"

	"gorm.io/gorm"
)

// VerificationCode 验证码模型
type VerificationCode struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Code        string         `gorm:"size:6;not null" json:"code"`
	Contact     string         `gorm:"size:100;not null;index" json:"contact"`
	Type        string         `gorm:"size:20;not null" json:"type"`    // email, sms
	Purpose     string         `gorm:"size:50;not null" json:"purpose"` // register, login, reset_password, verify_email
	Token       string         `gorm:"size:100;not null;uniqueIndex" json:"token"`
	ExpiresAt   time.Time      `gorm:"not null;index" json:"expires_at"`
	Attempts    int            `gorm:"default:0" json:"attempts"`
	MaxAttempts int            `gorm:"default:5" json:"max_attempts"`
	IsUsed      bool           `gorm:"default:false" json:"is_used"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 指定表名
func (VerificationCode) TableName() string {
	return "verification_codes"
}
