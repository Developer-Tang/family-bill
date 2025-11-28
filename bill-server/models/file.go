package models

import (
	"time"

	"gorm.io/gorm"
)

// File 文件模型
type File struct {
	FileID       uint           `gorm:"primaryKey" json:"file_id"`
	BookID       uint           `gorm:"not null;index" json:"book_id"`
	UserID       uint           `gorm:"not null" json:"user_id"`
	Filename     string         `gorm:"size:255;not null" json:"filename"`
	OriginalName string         `gorm:"size:255;not null" json:"original_name"`
	MimeType     string         `gorm:"size:100;not null" json:"mime_type"`
	Size         int64          `json:"size"`
	Path         string         `gorm:"size:255;not null" json:"path"`
	URL          string         `gorm:"size:255;not null" json:"url"`
	ThumbnailURL string         `gorm:"size:255" json:"thumbnail_url"`
	EntityType   string         `gorm:"size:20;not null;index" json:"entity_type"` // transaction, user, book, category
	EntityID     uint           `gorm:"index" json:"entity_id"`
	Description  string         `gorm:"size:255" json:"description"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Book Book `gorm:"foreignKey:BookID" json:"-"`
	User User `gorm:"foreignKey:UserID" json:"-"`
}

// TableName 指定表名
func (File) TableName() string {
	return "files"
}
