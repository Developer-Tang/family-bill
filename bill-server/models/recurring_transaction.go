package models

import (
	"time"

	"gorm.io/gorm"
)

// RecurringTransaction 周期记账模型
type RecurringTransaction struct {
	RecurringTransactionID uint           `gorm:"primaryKey" json:"recurring_transaction_id"`
	BookID                 uint           `gorm:"not null;index" json:"book_id"`
	UserID                 uint           `gorm:"not null" json:"user_id"`
	AccountID              uint           `gorm:"not null" json:"account_id"`
	CategoryID             uint           `json:"category_id"`
	Type                   string         `gorm:"size:20;not null" json:"type"` // income, expense
	Amount                 float64        `gorm:"type:decimal(15,2);not null" json:"amount"`
	Description            string         `gorm:"size:255" json:"description"`
	StartDate              time.Time      `json:"start_date"`
	EndDate                time.Time      `json:"end_date"`
	Frequency              string         `gorm:"size:20;not null" json:"frequency"` // daily, weekly, monthly, yearly
	Interval               int            `gorm:"default:1" json:"interval"`
	DayOfWeek              int            `json:"day_of_week"`
	DayOfMonth             int            `json:"day_of_month"`
	Status                 string         `gorm:"size:20;default:active" json:"status"`
	LastExecutedAt         time.Time      `json:"last_executed_at"`
	NextExecutedAt         time.Time      `json:"next_executed_at"`
	CreatedAt              time.Time      `json:"created_at"`
	UpdatedAt              time.Time      `json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Book     Book     `gorm:"foreignKey:BookID" json:"-"`
	User     User     `gorm:"foreignKey:UserID" json:"-"`
	Account  Account  `gorm:"foreignKey:AccountID" json:"account"`
	Category Category `gorm:"foreignKey:CategoryID" json:"category"`
}

// TransactionTemplate 记账模板模型
type TransactionTemplate struct {
	TransactionTemplateID uint           `gorm:"primaryKey" json:"transaction_template_id"`
	BookID                uint           `gorm:"not null;index" json:"book_id"`
	UserID                uint           `gorm:"not null" json:"user_id"`
	Name                  string         `gorm:"size:100;not null" json:"name"`
	AccountID             uint           `gorm:"not null" json:"account_id"`
	CategoryID            uint           `json:"category_id"`
	Type                  string         `gorm:"size:20;not null" json:"type"` // income, expense
	Amount                float64        `gorm:"type:decimal(15,2)" json:"amount"`
	Description           string         `gorm:"size:255" json:"description"`
	Tags                  string         `gorm:"size:255" json:"tags"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Book     Book     `gorm:"foreignKey:BookID" json:"-"`
	User     User     `gorm:"foreignKey:UserID" json:"-"`
	Account  Account  `gorm:"foreignKey:AccountID" json:"account"`
	Category Category `gorm:"foreignKey:CategoryID" json:"category"`
}

// TableName 指定表名
func (RecurringTransaction) TableName() string {
	return "recurring_transactions"
}

// TableName 指定表名
func (TransactionTemplate) TableName() string {
	return "transaction_templates"
}
