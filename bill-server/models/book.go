package models

import (
	"time"

	"gorm.io/gorm"
)

// Book 账本模型
type Book struct {
	BookID      uint           `gorm:"primaryKey" json:"book_id"`
	Name        string         `gorm:"size:100;not null" json:"name"`
	CreatorID   uint           `gorm:"not null" json:"creator_id"`
	Currency    string         `gorm:"size:10;default:CNY" json:"currency"`
	Description string         `gorm:"size:255" json:"description"`
	Status      string         `gorm:"size:20;default:active" json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Creator User         `gorm:"foreignKey:CreatorID" json:"-"`
	Access  []BookAccess `gorm:"foreignKey:BookID" json:"access"`
}

// BookAccess 账本访问权限模型
type BookAccess struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	BookID    uint           `gorm:"not null;index" json:"book_id"`
	UserID    uint           `gorm:"not null;index" json:"user_id"`
	Role      string         `gorm:"size:20;default:viewer" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Book Book `gorm:"foreignKey:BookID" json:"-"`
	User User `gorm:"foreignKey:UserID" json:"user"`
}

// TableName 指定表名
func (Book) TableName() string {
	return "books"
}

// TableName 指定表名
func (BookAccess) TableName() string {
	return "book_access"
}
