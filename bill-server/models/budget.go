package models

import (
	"time"

	"gorm.io/gorm"
)

// Budget 预算模型
type Budget struct {
	BudgetID    uint           `gorm:"primaryKey" json:"budget_id"`
	BookID      uint           `gorm:"not null;index" json:"book_id"`
	CategoryID  uint           `json:"category_id"`
	Amount      float64        `gorm:"type:decimal(15,2);not null" json:"amount"`
	Period      string         `gorm:"size:20;not null" json:"period"` // month, quarter, year
	Date        string         `gorm:"size:20;not null" json:"date"`   // YYYY-MM
	Description string         `gorm:"size:255" json:"description"`
	Status      string         `gorm:"size:20;default:active" json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Book     Book     `gorm:"foreignKey:BookID" json:"-"`
	Category Category `gorm:"foreignKey:CategoryID" json:"category"`
}

// TableName 指定表名
func (Budget) TableName() string {
	return "budgets"
}
