package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDashboardSummary 获取财务概览数据
func GetDashboardSummary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取财务概览数据成功",
		"data": gin.H{
			"total_income":  8000.00,
			"total_expense": 3500.00,
			"balance":       4500.00,
			"income_rate":   10.5,
			"expense_rate":  -5.2,
			"top_expense_category": gin.H{
				"id":         5,
				"name":       "餐饮",
				"amount":     1200.00,
				"percentage": 34.3,
			},
			"budget_status": gin.H{
				"total":      5000.00,
				"used":       3500.00,
				"percentage": 70.0,
				"alert":      false,
			},
			"account_summary": gin.H{
				"total_balance": 25000.00,
				"account_count": 5,
			},
		},
	})
}

// GetDashboardQuickStats 获取快速统计数据
func GetDashboardQuickStats(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取快速统计数据成功",
		"data": gin.H{
			"income_count":       10,
			"expense_count":      30,
			"avg_income":         800.00,
			"avg_expense":        116.67,
			"total_transactions": 40,
			"highest_income":     5000.00,
			"highest_expense":    800.00,
		},
	})
}

// GetDashboardRecentTransactions 获取最近收支记录
func GetDashboardRecentTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取最近收支记录成功",
		"data": gin.H{
			"list": []gin.H{
				{
					"id":            1001,
					"amount":        100.00,
					"type":          "expense",
					"category_name": "餐饮",
					"account_name":  "支付宝",
					"remark":        "午餐",
					"created_at":    "2023-01-31T12:00:00Z",
				},
				{
					"id":            1002,
					"amount":        5000.00,
					"type":          "income",
					"category_name": "工资",
					"account_name":  "我的工资卡",
					"remark":        "1月工资",
					"created_at":    "2023-01-30T10:00:00Z",
				},
			},
			"total":  40,
			"limit":  10,
			"offset": 0,
		},
	})
}

// GetDashboardBudgetProgress 获取预算执行进度
func GetDashboardBudgetProgress(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取预算执行进度成功",
		"data": gin.H{
			"total_budget":     5000.00,
			"used_budget":      3500.00,
			"remaining_budget": 1500.00,
			"usage_rate":       70.0,
			"categories": []gin.H{
				{
					"category_id":   5,
					"category_name": "餐饮",
					"budget":        1500.00,
					"used":          1200.00,
					"usage_rate":    80.0,
				},
				{
					"category_id":   10,
					"category_name": "住房",
					"budget":        1000.00,
					"used":          1000.00,
					"usage_rate":    100.0,
				},
			},
		},
	})
}
