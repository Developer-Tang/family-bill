package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAccountTypes 获取账户类型列表
func GetAccountTypes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": []gin.H{
			{
				"type_id":   1,
				"name":      "银行卡",
				"icon":      "bank_icon",
				"color":     "#0080FF",
				"is_system": true,
			},
			{
				"type_id":   2,
				"name":      "第三方支付",
				"icon":      "payment_icon",
				"color":     "#4CAF50",
				"is_system": true,
			},
		},
	})
}

// CreateAccountType 创建自定义账户类型
func CreateAccountType(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data": gin.H{
			"type_id":   10,
			"name":      "投资账户",
			"icon":      "investment_icon",
			"color":     "#FF9800",
			"is_system": false,
		},
	})
}

// UpdateAccountType 更新账户类型
func UpdateAccountType(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
		"data": gin.H{
			"type_id":   10,
			"name":      "投资账户(更新)",
			"icon":      "new_investment_icon",
			"color":     "#FF5722",
			"is_system": false,
		},
	})
}

// DeleteAccountType 删除账户类型
func DeleteAccountType(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
		"data":    nil,
	})
}
