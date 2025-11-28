package models

import (
	"time"

	"gorm.io/gorm"
)

// Transaction 交易记录模型
type Transaction struct {
	TransactionID uint           `gorm:"primaryKey" json:"transaction_id"`
	BookID        uint           `gorm:"not null;index" json:"book_id"`
	UserID        uint           `gorm:"not null" json:"user_id"`
	AccountID     uint           `gorm:"not null" json:"account_id"`
	CategoryID    uint           `json:"category_id"`
	Type          string         `gorm:"size:20;not null" json:"type"` // income, expense, transfer
	Amount        float64        `gorm:"type:decimal(15,2);not null" json:"amount"`
	Description   string         `gorm:"size:255" json:"description"`
	Date          time.Time      `gorm:"index" json:"date"`
	Status        string         `gorm:"size:20;default:active" json:"status"`
	Locked        bool           `gorm:"default:false" json:"locked"`
	Location      string         `gorm:"size:255" json:"location"`
	Latitude      float64        `json:"latitude"`
	Longitude     float64        `json:"longitude"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Book     Book     `gorm:"foreignKey:BookID" json:"-"`
	User     User     `gorm:"foreignKey:UserID" json:"-"`
	Account  Account  `gorm:"foreignKey:AccountID" json:"account"`
	Category Category `gorm:"foreignKey:CategoryID" json:"category"`
	Tags     []Tag    `gorm:"many2many:transaction_tags;" json:"tags"`
	Files    []File   `gorm:"foreignKey:EntityID;constraint:OnDelete:CASCADE;" json:"files"`
}

// Category 分类模型
type Category struct {
	CategoryID  uint           `gorm:"primaryKey" json:"category_id"`
	Name        string         `gorm:"size:50;not null" json:"name"`
	Icon        string         `gorm:"size:50" json:"icon"`
	Color       string         `gorm:"size:20" json:"color"`
	Type        string         `gorm:"size:20;not null" json:"type"` // income, expense
	ParentID    uint           `json:"parent_id"`
	Description string         `gorm:"size:255" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Parent       *Category     `gorm:"foreignKey:ParentID" json:"parent"`
	Children     []Category    `gorm:"foreignKey:ParentID" json:"children"`
	Transactions []Transaction `gorm:"foreignKey:CategoryID" json:"transactions"`
}

// Tag 标签模型
type Tag struct {
	TagID       uint           `gorm:"primaryKey" json:"tag_id"`
	Name        string         `gorm:"size:50;not null;uniqueIndex" json:"name"`
	Color       string         `gorm:"size:20" json:"color"`
	Description string         `gorm:"size:255" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Transactions []Transaction `gorm:"many2many:transaction_tags;" json:"transactions"`
}

// TransactionTag 交易标签关联模型
type TransactionTag struct {
	TransactionID uint      `gorm:"primaryKey" json:"transaction_id"`
	TagID         uint      `gorm:"primaryKey" json:"tag_id"`
	CreatedAt     time.Time `json:"created_at"`
}

// TableName 指定表名
func (Transaction) TableName() string {
	return "transactions"
}

// TableName 指定表名
func (Category) TableName() string {
	return "categories"
}

// TableName 指定表名
func (Tag) TableName() string {
	return "tags"
}

// TableName 指定表名
func (TransactionTag) TableName() string {
	return "transaction_tags"
}
