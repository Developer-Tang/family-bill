package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTransaction 创建收支记录
func CreateTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "记账成功",
		"data": gin.H{
			"transaction_id": 1,
			"type":           "expense",
			"amount":         100.50,
			"currency":       "CNY",
			"account_name":   "我的工资卡",
			"category_name":  "餐饮",
			"date":           "2023-01-05",
			"time":           "12:30",
			"memo":           "午餐",
			"created_at":     "2023-01-05T12:30:00Z",
		},
	})
}

// GetTransactions 获取收支记录列表
func GetTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"total": 150,
			"items": []gin.H{
				{
					"transaction_id": 1,
					"type":           "expense",
					"amount":         100.50,
					"currency":       "CNY",
					"account_id":     1,
					"account_name":   "我的工资卡",
					"category_id":    5,
					"category_name":  "餐饮",
					"date":           "2023-01-05",
					"memo":           "午餐",
					"tags":           []string{"日常", "午餐"},
					"is_locked":      false,
					"created_at":     "2023-01-05T12:30:00Z",
				},
				{
					"transaction_id": 2,
					"type":           "expense",
					"amount":         50.00,
					"currency":       "CNY",
					"account_id":     2,
					"account_name":   "支付宝",
					"category_id":    6,
					"category_name":  "交通",
					"date":           "2023-01-05",
					"memo":           "打车",
					"tags":           []string{"交通"},
					"is_locked":      false,
					"created_at":     "2023-01-05T15:45:00Z",
				},
			},
		},
	})
}

// GetTransaction 获取收支记录详情
func GetTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"transaction_id": 1,
			"book_id":        1,
			"type":           "expense",
			"amount":         100.50,
			"currency":       "CNY",
			"account_id":     1,
			"account_name":   "我的工资卡",
			"category_id":    5,
			"category_name":  "餐饮",
			"subcategory_id": nil,
			"date":           "2023-01-05",
			"time":           "12:30",
			"memo":           "午餐",
			"tags":           []string{"日常", "午餐"},
			"tag_ids":        []int{1, 2},
			"member_id":      1,
			"member_name":    "张三",
			"attachments":    []string{"url1", "url2"},
			"location": gin.H{
				"latitude":  39.9042,
				"longitude": 116.4074,
				"address":   "北京市朝阳区",
			},
			"is_locked":  false,
			"created_at": "2023-01-05T12:30:00Z",
			"updated_at": "2023-01-05T12:30:00Z",
		},
	})
}

// UpdateTransaction 更新收支记录
func UpdateTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
		"data": gin.H{
			"transaction_id": 1,
			"type":           "expense",
			"amount":         120.50,
			"currency":       "CNY",
			"account_name":   "我的工资卡",
			"category_name":  "餐饮",
			"date":           "2023-01-05",
			"time":           "12:30",
			"memo":           "午餐（更新）",
			"updated_at":     "2023-01-05T13:00:00Z",
		},
	})
}

// DeleteTransaction 删除收支记录
func DeleteTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
		"data":    nil,
	})
}

// LockTransaction 锁定解锁收支记录
func LockTransaction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "操作成功",
		"data": gin.H{
			"transaction_id": 1,
			"is_locked":      true,
			"updated_at":     "2023-01-05T13:30:00Z",
		},
	})
}

// BatchCreateTransactions 批量创建收支记录
func BatchCreateTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "批量记账成功",
		"data": gin.H{
			"success_count":   3,
			"failed_count":    0,
			"transaction_ids": []int{1, 2, 3},
		},
	})
}

// ImportTransactions 导入外部账单
func ImportTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "导入成功",
		"data": gin.H{
			"import_id":     "uuid-string",
			"total_count":   50,
			"success_count": 48,
			"failed_count":  2,
			"failed_records": []gin.H{
				{
					"index": 10,
					"error": "无效的金额格式",
				},
				{
					"index": 45,
					"error": "无法识别的账户",
				},
			},
			"transaction_ids": []int{4, 5, 6, 7, 8, 9, 10},
		},
	})
}
