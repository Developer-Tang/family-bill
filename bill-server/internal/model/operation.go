package model

import (
	"time"

	"gorm.io/gorm"
)

// Operation 操作记录模型
type Operation struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint      `gorm:"not null" json:"user_id"`
	IP         string    `gorm:"size:128;not null" json:"ip"`
	ClientType string    `gorm:"size:128;not null" json:"client_type"`
	UA         string    `gorm:"size:128;not null" json:"ua"`
	Path       string    `gorm:"size:128;not null" json:"path"`
	Method     string    `gorm:"size:128;not null" json:"method"`
	Params     string    `gorm:"type:text" json:"params"`
	Body       string    `gorm:"type:text" json:"body"`
	Result     string    `gorm:"type:text" json:"result"`
	CreatedAt  time.Time `gorm:"type:timestamp;not null" json:"created_at"`
}

// TableName 指定表名
func (Operation) TableName() string {
	return "tb_operations"
}

// BeforeCreate 每次创建前设置CreatedAt
func (o *Operation) BeforeCreate(tx *gorm.DB) (err error) {
	o.CreatedAt = time.Now()
	return
}
