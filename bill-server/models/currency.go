package models

import (
	"time"

	"gorm.io/gorm"
)

// Currency 货币模型
type Currency struct {
	CurrencyID uint           `gorm:"primaryKey" json:"currency_id"`
	Code       string         `gorm:"size:10;uniqueIndex;not null" json:"code"`
	Name       string         `gorm:"size:50;not null" json:"name"`
	Symbol     string         `gorm:"size:10" json:"symbol"`
	Rate       float64        `gorm:"type:decimal(10,6);default:1" json:"rate"`
	IsBase     bool           `gorm:"default:false" json:"is_base"`
	Status     string         `gorm:"size:20;default:active" json:"status"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

// ExchangeRate 汇率模型
type ExchangeRate struct {
	ExchangeRateID uint      `gorm:"primaryKey" json:"exchange_rate_id"`
	FromCurrency   string    `gorm:"size:10;not null;index:idx_currency_pair" json:"from_currency"`
	ToCurrency     string    `gorm:"size:10;not null;index:idx_currency_pair" json:"to_currency"`
	Rate           float64   `gorm:"type:decimal(15,6);not null" json:"rate"`
	Source         string    `gorm:"size:50" json:"source"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}

// TableName 指定表名
func (Currency) TableName() string {
	return "currencies"
}

// TableName 指定表名
func (ExchangeRate) TableName() string {
	return "exchange_rates"
}
