package model

import (
	"time"
)

// Operation 操作记录模型
type Operation struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	IP        string    `gorm:"size:128;not null" json:"ip"`
	UA        string    `gorm:"size:128;not null" json:"ua"`
	Path      string    `gorm:"size:128;not null" json:"path"`
	Method    string    `gorm:"size:128;not null" json:"method"`
	Params    string    `gorm:"type:text" json:"params"`
	Body      string    `gorm:"type:text" json:"body"`
	Result    string    `gorm:"type:text" json:"result"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (Operation) TableName() string {
	return "tb_operations"
}
