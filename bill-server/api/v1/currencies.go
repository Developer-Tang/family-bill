package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCurrencies 获取货币列表
func GetCurrencies(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取货币列表成功",
		"data": []gin.H{
			{"code": "CNY", "name": "人民币", "symbol": "¥"},
			{"code": "USD", "name": "美元", "symbol": "$"},
			{"code": "EUR", "name": "欧元", "symbol": "€"},
			{"code": "GBP", "name": "英镑", "symbol": "£"},
			{"code": "JPY", "name": "日元", "symbol": "¥"},
		},
	})
}

// GetExchangeRates 获取汇率
func GetExchangeRates(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取汇率成功",
		"data": gin.H{
			"base":       "CNY",
			"rates":      gin.H{"USD": 0.14, "EUR": 0.13, "GBP": 0.11},
			"updated_at": "2023-01-01T12:00:00Z",
		},
	})
}

// UpdateExchangeRates 更新汇率
func UpdateExchangeRates(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新汇率成功",
		"data": gin.H{
			"base":       "CNY",
			"updated_at": "2023-01-01T12:00:00Z",
		},
	})
}
