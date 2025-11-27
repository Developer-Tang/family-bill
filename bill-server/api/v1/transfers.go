package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTransfer 创建转账记录
func CreateTransfer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "转账成功",
		"data": gin.H{
			"transfer_id": 1,
			"from_account": gin.H{
				"id":      1,
				"name":    "我的工资卡",
				"balance": 5000.00,
			},
			"to_account": gin.H{
				"id":      2,
				"name":    "支付宝",
				"balance": 3000.00,
			},
			"amount":        1000.00,
			"currency":      "CNY",
			"transfer_date": "2023-01-05T00:00:00Z",
			"memo":          "转账到支付宝",
		},
	})
}

// GetTransfers 获取转账记录列表
func GetTransfers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取转账记录列表成功",
		"data": gin.H{
			"total": 5,
			"items": []gin.H{
				{
					"transfer_id":       1,
					"from_account_id":   1,
					"from_account_name": "我的工资卡",
					"to_account_id":     2,
					"to_account_name":   "支付宝",
					"amount":            1000.00,
					"currency":          "CNY",
					"transfer_date":     "2023-01-05T00:00:00Z",
					"memo":              "转账到支付宝",
				},
				{
					"transfer_id":       2,
					"from_account_id":   2,
					"from_account_name": "支付宝",
					"to_account_id":     3,
					"to_account_name":   "微信钱包",
					"amount":            500.00,
					"currency":          "CNY",
					"transfer_date":     "2023-01-10T00:00:00Z",
					"memo":              "转账到微信",
				},
			},
		},
	})
}

// GetTransfer 获取转账记录详情
func GetTransfer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取转账记录详情成功",
		"data": gin.H{
			"transfer_id":       1,
			"book_id":           1,
			"from_account_id":   1,
			"from_account_name": "我的工资卡",
			"from_account_type": "银行卡",
			"to_account_id":     2,
			"to_account_name":   "支付宝",
			"to_account_type":   "第三方支付",
			"amount":            1000.00,
			"currency":          "CNY",
			"transfer_date":     "2023-01-05T00:00:00Z",
			"memo":              "转账到支付宝",
			"created_at":        "2023-01-05T12:00:00Z",
		},
	})
}
