package models

import (
	"time"

	"gorm.io/gorm"
)

// HelpCategory 帮助分类模型
type HelpCategory struct {
	CategoryID  uint           `gorm:"primaryKey" json:"category_id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Description string         `gorm:"size:255" json:"description"`
	Language    string         `gorm:"size:10;default:zh-CN" json:"language"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Articles []HelpArticle `gorm:"foreignKey:CategoryID" json:"articles"`
}

// HelpArticle 帮助文章模型
type HelpArticle struct {
	ArticleID    uint           `gorm:"primaryKey" json:"article_id"`
	CategoryID   uint           `gorm:"not null" json:"category_id"`
	Title        string         `gorm:"size:100;not null" json:"title"`
	Subtitle     string         `gorm:"size:255" json:"subtitle"`
	Summary      string         `gorm:"size:500" json:"summary"`
	Content      string         `gorm:"type:text;not null" json:"content"`
	CoverImage   string         `gorm:"size:255" json:"cover_image"`
	Language     string         `gorm:"size:10;default:zh-CN" json:"language"`
	ViewCount    int            `gorm:"default:0" json:"view_count"`
	LikeCount    int            `gorm:"default:0" json:"like_count"`
	CommentCount int            `gorm:"default:0" json:"comment_count"`
	Status       string         `gorm:"size:20;default:published" json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Category HelpCategory `gorm:"foreignKey:CategoryID" json:"category"`
}

// Feedback 反馈模型
type Feedback struct {
	FeedbackID uint           `gorm:"primaryKey" json:"feedback_id"`
	UserID     uint           `json:"user_id"`
	Type       string         `gorm:"size:20;not null" json:"type"` // question, suggestion, bug, other
	Content    string         `gorm:"size:500;not null" json:"content"`
	Contact    string         `gorm:"size:100" json:"contact"`
	Status     string         `gorm:"size:20;default:pending" json:"status"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	User User `gorm:"foreignKey:UserID" json:"user"`
}

// TableName 指定表名
func (HelpCategory) TableName() string {
	return "help_categories"
}

// TableName 指定表名
func (HelpArticle) TableName() string {
	return "help_articles"
}

// TableName 指定表名
func (Feedback) TableName() string {
	return "feedback"
}
