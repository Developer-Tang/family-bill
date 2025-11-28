package models

import (
	"time"

	"gorm.io/gorm"
)

// AccountType 账户类型模型
type AccountType struct {
	AccountTypeID uint           `gorm:"primaryKey" json:"account_type_id"`
	Name          string         `gorm:"size:50;not null" json:"name"`
	Icon          string         `gorm:"size:50" json:"icon"`
	Color         string         `gorm:"size:20" json:"color"`
	Type          string         `gorm:"size:20;not null" json:"type"` // asset, liability, income, expense
	Description   string         `gorm:"size:255" json:"description"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Accounts []Account `gorm:"foreignKey:AccountTypeID" json:"accounts"`
}

// AccountGroup 账户分组模型
type AccountGroup struct {
	AccountGroupID uint           `gorm:"primaryKey" json:"account_group_id"`
	Name           string         `gorm:"size:50;not null" json:"name"`
	Description    string         `gorm:"size:255" json:"description"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Accounts []Account `gorm:"foreignKey:AccountGroupID" json:"accounts"`
}

// Account 账户模型
type Account struct {
	AccountID      uint           `gorm:"primaryKey" json:"account_id"`
	BookID         uint           `gorm:"not null;index" json:"book_id"`
	AccountTypeID  uint           `gorm:"not null" json:"account_type_id"`
	AccountGroupID uint           `json:"account_group_id"`
	Name           string         `gorm:"size:100;not null" json:"name"`
	Balance        float64        `gorm:"type:decimal(15,2);default:0" json:"balance"`
	InitialBalance float64        `gorm:"type:decimal(15,2);default:0" json:"initial_balance"`
	Currency       string         `gorm:"size:10;default:CNY" json:"currency"`
	Description    string         `gorm:"size:255" json:"description"`
	Status         string         `gorm:"size:20;default:active" json:"status"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Book         Book          `gorm:"foreignKey:BookID" json:"-"`
	AccountType  AccountType   `gorm:"foreignKey:AccountTypeID" json:"account_type"`
	AccountGroup AccountGroup  `gorm:"foreignKey:AccountGroupID" json:"account_group"`
	Transactions []Transaction `gorm:"foreignKey:AccountID" json:"transactions"`
}

// TableName 指定表名
func (AccountType) TableName() string {
	return "account_types"
}

// TableName 指定表名
func (AccountGroup) TableName() string {
	return "account_groups"
}

// TableName 指定表名
func (Account) TableName() string {
	return "accounts"
}
