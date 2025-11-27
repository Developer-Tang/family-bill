package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateQuickTransaction 创建快速记账记录
func CreateQuickTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "快速记账成功",
		"data": gin.H{
			"transaction_id": 1,
			"type":           "expense",
			"amount":         30.00,
			"currency":       "CNY",
			"account_name":   "支付宝",
			"category_name":  "餐饮",
			"date":           "2023-01-05",
			"time":           "08:30",
			"memo":           "早餐",
			"created_at":     "2023-01-05T08:30:00Z",
		},
	})
}

// GetTransactionTemplates 获取记账模板
func GetTransactionTemplates(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取记账模板成功",
		"data": []gin.H{
			{
				"template_id": 1,
				"book_id":     1,
				"name":        "午餐模板",
				"type":        "expense",
				"amount":      30.00,
				"account_id":  2,
				"category_id": 5,
				"tags":        []int{1, 2},
				"memo":        "午餐",
				"created_at":  "2023-01-01T12:00:00Z",
			},
			{
				"template_id": 2,
				"book_id":     1,
				"name":        "打车模板",
				"type":        "expense",
				"amount":      50.00,
				"account_id":  2,
				"category_id": 6,
				"tags":        []int{3},
				"memo":        "打车",
				"created_at":  "2023-01-02T10:30:00Z",
			},
		},
	})
}

// CreateTransactionTemplate 创建记账模板
func CreateTransactionTemplate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建记账模板成功",
		"data": gin.H{
			"template_id": 3,
			"book_id":     1,
			"name":        "晚餐模板",
			"type":        "expense",
			"amount":      40.00,
			"account_id":  2,
			"category_id": 5,
			"tags":        []int{1, 4},
			"memo":        "晚餐",
			"created_at":  "2023-01-05T18:00:00Z",
		},
	})
}

// UpdateTransactionTemplate 更新记账模板
func UpdateTransactionTemplate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新记账模板成功",
		"data": gin.H{
			"template_id": 1,
			"book_id":     1,
			"name":        "午餐模板（更新）",
			"type":        "expense",
			"amount":      35.00,
			"account_id":  2,
			"category_id": 5,
			"tags":        []int{1, 2},
			"memo":        "午餐（更新）",
			"updated_at":  "2023-01-05T12:00:00Z",
		},
	})
}

// DeleteTransactionTemplate 删除记账模板
func DeleteTransactionTemplate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除记账模板成功",
		"data":    nil,
	})
}
