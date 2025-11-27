package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetIncomeExpenseAnalysis 获取收支对比分析
func GetIncomeExpenseAnalysis(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取收支对比分析成功",
		"data": gin.H{
			"current": gin.H{
				"income":  8000.00,
				"expense": 3500.00,
				"balance": 4500.00,
			},
			"compare": gin.H{
				"income":  7200.00,
				"expense": 3800.00,
				"balance": 3400.00,
			},
			"change": gin.H{
				"income_rate":  11.1,
				"expense_rate": -7.9,
				"balance_rate": 32.4,
			},
		},
	})
}

// GetTrendAnalysis 获取收支趋势分析
func GetTrendAnalysis(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取收支趋势分析成功",
		"data": gin.H{
			"period":     "month",
			"start_date": "2023-01",
			"end_date":   "2023-12",
			"data": []gin.H{
				{
					"date":    "2023-01",
					"income":  8000.00,
					"expense": 3500.00,
					"balance": 4500.00,
				},
				{
					"date":    "2023-02",
					"income":  8200.00,
					"expense": 3800.00,
					"balance": 4400.00,
				},
			},
			"summary": gin.H{
				"total_income":  96000.00,
				"total_expense": 42000.00,
				"total_balance": 54000.00,
			},
		},
	})
}

// GetFlowAnalysis 获取资金流向分析
func GetFlowAnalysis(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取资金流向分析成功",
		"data": gin.H{
			"income_flow": []gin.H{
				{
					"category_id":   1,
					"category_name": "工资",
					"amount":        8000.00,
					"percentage":    100.0,
				},
			},
			"expense_flow": []gin.H{
				{
					"category_id":   5,
					"category_name": "餐饮",
					"amount":        1200.00,
					"percentage":    34.3,
				},
				{
					"category_id":   10,
					"category_name": "住房",
					"amount":        1000.00,
					"percentage":    28.6,
				},
			},
			"transfer_flow": []gin.H{
				{
					"from_account_id":   1,
					"from_account_name": "工资卡",
					"to_account_id":     2,
					"to_account_name":   "支付宝",
					"amount":            2000.00,
				},
			},
		},
	})
}
