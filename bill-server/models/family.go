package models

import (
	"time"

	"gorm.io/gorm"
)

// Family 家庭组模型
type Family struct {
	FamilyID    uint           `gorm:"primaryKey" json:"family_id"`
	Name        string         `gorm:"size:100;not null" json:"name"`
	CreatorID   uint           `gorm:"not null" json:"creator_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Avatar      string         `gorm:"size:255" json:"avatar"`
	Description string         `gorm:"size:255" json:"description"`

	// 关联关系
	Creator User           `gorm:"foreignKey:CreatorID" json:"-"`
	Members []FamilyMember `gorm:"foreignKey:FamilyID" json:"members"`
}

// FamilyMember 家庭组成员模型
type FamilyMember struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	FamilyID  uint           `gorm:"not null;index" json:"family_id"`
	UserID    uint           `gorm:"not null;index" json:"user_id"`
	Role      string         `gorm:"size:20;default:member" json:"role"`
	JoinedAt  time.Time      `json:"joined_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Family Family `gorm:"foreignKey:FamilyID" json:"-"`
	User   User   `gorm:"foreignKey:UserID" json:"user"`
}

// TableName 指定表名
func (Family) TableName() string {
	return "families"
}

// TableName 指定表名
func (FamilyMember) TableName() string {
	return "family_members"
}
