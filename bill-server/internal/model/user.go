package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:128;not null" json:"username"`
	Password  string    `gorm:"size:128;not null" json:"-"`
	Email     string    `gorm:"uniqueIndex;size:128;not null" json:"email"`
	Avatar    string    `gorm:"type:text" json:"avatar"`
	HaveRole  string    `gorm:"size:128;default:''" json:"have_role"`
	Status    int8      `gorm:"default:0" json:"status"`
	LastIP    string    `gorm:"size:128;default:''" json:"last_ip"`
	LastLogin time.Time `gorm:"type:timestamp" json:"last_login"`
	CreatedAt time.Time `gorm:"type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null" json:"updated_at"`
}

// TableName 指定表名
func (User) TableName() string {
	return "tb_users"
}

// BeforeCreate 每次创建前设置CreatedAt和UpdatedAt
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

// BeforeUpdate 每次更新前设置UpdatedAt
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
