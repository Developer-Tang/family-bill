package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCategoryReport 获取分类统计报表
func GetCategoryReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取分类统计报表成功",
		"data": gin.H{
			"type":         "expense",
			"period":       "month",
			"date":         "2023-01",
			"total_amount": 3500.00,
			"categories": []gin.H{
				{
					"category_id":   5,
					"category_name": "餐饮",
					"amount":        1200.00,
					"percentage":    34.3,
					"count":         30,
					"icon":          "food_icon",
					"color":         "#FF6B6B",
				},
				{
					"category_id":   10,
					"category_name": "住房",
					"amount":        1000.00,
					"percentage":    28.6,
					"count":         5,
					"icon":          "home_icon",
					"color":         "#4ECDC4",
				},
			},
		},
	})
}

// GetCategoryTrendReport 获取分类趋势报表
func GetCategoryTrendReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取分类趋势报表成功",
		"data": gin.H{
			"category_id":   5,
			"category_name": "餐饮",
			"type":          "expense",
			"period":        "month",
			"start_date":    "2023-01",
			"end_date":      "2023-12",
			"data": []gin.H{
				{
					"date":   "2023-01",
					"amount": 1200.00,
					"count":  30,
				},
				{
					"date":   "2023-02",
					"amount": 1300.00,
					"count":  32,
				},
			},
			"summary": gin.H{
				"total_amount": 15000.00,
				"total_count":  365,
				"avg_amount":   1250.00,
			},
		},
	})
}

// GetAccountReport 获取账户收支报表
func GetAccountReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取账户收支报表成功",
		"data": gin.H{
			"type":   "both",
			"period": "month",
			"date":   "2023-01",
			"accounts": []gin.H{
				{
					"account_id":     1,
					"account_name":   "我的工资卡",
					"account_type":   "银行卡",
					"income":         8000.00,
					"expense":        2000.00,
					"balance_change": 6000.00,
				},
				{
					"account_id":     2,
					"account_name":   "支付宝",
					"account_type":   "第三方支付",
					"income":         0.00,
					"expense":        1500.00,
					"balance_change": -1500.00,
				},
			},
		},
	})
}

// GetAccountBalanceReport 获取账户余额报表
func GetAccountBalanceReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取账户余额报表成功",
		"data": gin.H{
			"date":          "2023-01-31",
			"total_balance": 25000.00,
			"accounts": []gin.H{
				{
					"account_id":   1,
					"account_name": "我的工资卡",
					"account_type": "银行卡",
					"balance":      15000.00,
					"percentage":   60.0,
					"currency":     "CNY",
					"icon":         "bank_icon",
					"color":        "#0080FF",
				},
				{
					"account_id":   2,
					"account_name": "支付宝",
					"account_type": "第三方支付",
					"balance":      5000.00,
					"percentage":   20.0,
					"currency":     "CNY",
					"icon":         "alipay_icon",
					"color":        "#1677FF",
				},
			},
			"balance_by_type": []gin.H{
				{
					"type":       "银行卡",
					"amount":     18000.00,
					"percentage": 72.0,
				},
				{
					"type":       "第三方支付",
					"amount":     7000.00,
					"percentage": 28.0,
				},
			},
		},
	})
}

// GetMemberReport 获取成员收支报表
func GetMemberReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成员收支报表成功",
		"data": gin.H{
			"period": "month",
			"date":   "2023-01",
			"members": []gin.H{
				{
					"member_id":   1,
					"member_name": "张三",
					"income":      8000.00,
					"expense":     2000.00,
					"balance":     6000.00,
				},
				{
					"member_id":   2,
					"member_name": "李四",
					"income":      0.00,
					"expense":     1500.00,
					"balance":     -1500.00,
				},
			},
		},
	})
}

// GetMemberContributionReport 获取成员贡献度分析
func GetMemberContributionReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成员贡献度分析成功",
		"data": gin.H{
			"period":        "month",
			"date":          "2023-01",
			"total_income":  8000.00,
			"total_expense": 3500.00,
			"members": []gin.H{
				{
					"member_id":          1,
					"member_name":        "张三",
					"income":             8000.00,
					"income_percentage":  100.0,
					"expense":            2000.00,
					"expense_percentage": 57.1,
					"net_contribution":   6000.00,
				},
				{
					"member_id":          2,
					"member_name":        "李四",
					"income":             0.00,
					"income_percentage":  0.0,
					"expense":            1500.00,
					"expense_percentage": 42.9,
					"net_contribution":   -1500.00,
				},
			},
		},
	})
}

// GetBudgetReport 获取预算执行报表
func GetBudgetReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取预算执行报表成功",
		"data": gin.H{
			"period":             "month",
			"date":               "2023-01",
			"total_budget":       5000.00,
			"total_used":         3500.00,
			"total_remaining":    1500.00,
			"overall_usage_rate": 70.0,
			"categories": []gin.H{
				{
					"category_id":    5,
					"category_name":  "餐饮",
					"budget":         1500.00,
					"used":           1200.00,
					"remaining":      300.00,
					"usage_rate":     80.0,
					"is_over_budget": false,
				},
				{
					"category_id":    6,
					"category_name":  "交通",
					"budget":         500.00,
					"used":           600.00,
					"remaining":      -100.00,
					"usage_rate":     120.0,
					"is_over_budget": true,
				},
			},
		},
	})
}

// GetBudgetAlertReport 获取预算超支提醒
func GetBudgetAlertReport(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取预算超支提醒成功",
		"data": gin.H{
			"period":    "month",
			"date":      "2023-01",
			"threshold": 80,
			"alerts": []gin.H{
				{
					"category_id":   5,
					"category_name": "餐饮",
					"budget":        1500.00,
					"used":          1200.00,
					"usage_rate":    80.0,
					"over_amount":   0.00,
				},
				{
					"category_id":   6,
					"category_name": "交通",
					"budget":        500.00,
					"used":          600.00,
					"usage_rate":    120.0,
					"over_amount":   100.00,
				},
			},
		},
	})
}
