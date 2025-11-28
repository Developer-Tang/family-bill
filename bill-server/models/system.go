package models

import (
	"time"

	"gorm.io/gorm"
)

// OperationLog 操作日志模型
type OperationLog struct {
	LogID       uint           `gorm:"primaryKey" json:"log_id"`
	UserID      uint           `gorm:"not null;index" json:"user_id"`
	Username    string         `gorm:"size:50" json:"username"`
	Action      string         `gorm:"size:50;not null" json:"action"`
	Description string         `gorm:"size:255" json:"description"`
	IPAddress   string         `gorm:"size:50" json:"ip_address"`
	UserAgent   string         `gorm:"size:255" json:"user_agent"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	User User `gorm:"foreignKey:UserID" json:"-"`
}

// LoginDevice 登录设备模型
type LoginDevice struct {
	DeviceID   uint           `gorm:"primaryKey" json:"device_id"`
	UserID     uint           `gorm:"not null;index" json:"user_id"`
	DeviceName string         `gorm:"size:100" json:"device_name"`
	IPAddress  string         `gorm:"size:50" json:"ip_address"`
	UserAgent  string         `gorm:"size:255" json:"user_agent"`
	LastLogin  time.Time      `json:"last_login"`
	IsCurrent  bool           `gorm:"default:false" json:"is_current"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	User User `gorm:"foreignKey:UserID" json:"-"`
}

// TwoFactor 两步验证模型
type TwoFactor struct {
	TwoFactorID   uint           `gorm:"primaryKey" json:"two_factor_id"`
	UserID        uint           `gorm:"not null;uniqueIndex" json:"user_id"`
	Enabled       bool           `gorm:"default:false" json:"enabled"`
	SecretKey     string         `gorm:"size:100" json:"secret_key"`
	RecoveryCodes string         `gorm:"type:text" json:"recovery_codes"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	User User `gorm:"foreignKey:UserID" json:"-"`
}

// TableName 指定表名
func (OperationLog) TableName() string {
	return "operation_logs"
}

// TableName 指定表名
func (LoginDevice) TableName() string {
	return "login_devices"
}

// TableName 指定表名
func (TwoFactor) TableName() string {
	return "two_factors"
}
