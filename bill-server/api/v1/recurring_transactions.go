package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRecurringTransactions 获取周期记账列表
func GetRecurringTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取周期记账列表成功",
		"data": []gin.H{
			{
				"recurring_id":       1,
				"book_id":            1,
				"name":               "房租",
				"type":               "expense",
				"amount":             3000.00,
				"currency":           "CNY",
				"account_id":         1,
				"category_id":        10,
				"frequency":          "monthly",
				"interval":           1,
				"start_date":         "2023-01-01",
				"end_date":           "2023-12-31",
				"next_date":          "2023-02-01",
				"memo":               "每月房租",
				"tags":               []int{3},
				"auto_create":        true,
				"notify_before_days": 3,
				"status":             "active",
				"created_at":         "2023-01-01T12:00:00Z",
			},
			{
				"recurring_id":       2,
				"book_id":            1,
				"name":               "工资",
				"type":               "income",
				"amount":             8000.00,
				"currency":           "CNY",
				"account_id":         1,
				"category_id":        20,
				"frequency":          "monthly",
				"interval":           1,
				"start_date":         "2023-01-01",
				"end_date":           nil,
				"next_date":          "2023-02-10",
				"memo":               "每月工资",
				"tags":               []int{5},
				"auto_create":        true,
				"notify_before_days": 0,
				"status":             "active",
				"created_at":         "2023-01-01T12:00:00Z",
			},
		},
	})
}

// CreateRecurringTransaction 创建周期记账
func CreateRecurringTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建周期记账成功",
		"data": gin.H{
			"recurring_id": 3,
			"name":         "水电费",
			"type":         "expense",
			"amount":       500.00,
			"frequency":    "monthly",
			"interval":     1,
			"next_date":    "2023-02-05",
			"status":       "active",
		},
	})
}

// UpdateRecurringTransaction 更新周期记账
func UpdateRecurringTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新周期记账成功",
		"data": gin.H{
			"recurring_id": 1,
			"name":         "房租（更新）",
			"type":         "expense",
			"amount":       3200.00,
			"frequency":    "monthly",
			"interval":     1,
			"next_date":    "2023-02-01",
			"status":       "active",
		},
	})
}

// DeleteRecurringTransaction 删除周期记账
func DeleteRecurringTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除周期记账成功",
		"data":    nil,
	})
}

// TriggerRecurringTransaction 手动触发周期记账
func TriggerRecurringTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "手动触发周期记账成功",
		"data": gin.H{
			"recurring_id":    1,
			"transaction_ids": []int{10, 11}, // 生成的交易记录ID
			"next_date":       "2023-03-01",
		},
	})
}
