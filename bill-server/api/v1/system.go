package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOperationLogs 获取操作审计日志
func GetOperationLogs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取操作审计日志成功",
		"data": gin.H{
			"logs": []gin.H{
				{
					"log_id":      1,
					"user_id":     1,
					"username":    "张三",
					"action":      "login",
					"description": "用户登录系统",
					"ip_address":  "192.168.1.100",
					"user_agent":  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
					"created_at":  "2023-01-01T12:00:00Z",
				},
				{
					"log_id":      2,
					"user_id":     1,
					"username":    "张三",
					"action":      "create_family",
					"description": "创建家庭组",
					"ip_address":  "192.168.1.100",
					"user_agent":  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
					"created_at":  "2023-01-01T12:05:00Z",
				},
			},
			"total": 2,
			"page":  1,
			"size":  10,
		},
	})
}

// SetupTwoFactor 设置两步验证
func SetupTwoFactor(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "设置两步验证成功",
		"data": gin.H{
			"enabled":     true,
			"secret_key":  "JBSWY3DPEHPK3PXP",
			"qr_code_url": "https://example.com/qrcode.png",
			"recovery_codes": []string{
				"123456",
				"234567",
				"345678",
				"456789",
				"567890",
			},
		},
	})
}

// GetLoginDevices 获取登录设备列表
func GetLoginDevices(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取登录设备列表成功",
		"data": gin.H{
			"devices": []gin.H{
				{
					"device_id":   1,
					"device_name": "Windows PC",
					"ip_address":  "192.168.1.100",
					"user_agent":  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36",
					"last_login":  "2023-01-01T12:00:00Z",
					"is_current":  true,
				},
				{
					"device_id":   2,
					"device_name": "iPhone",
					"ip_address":  "192.168.1.101",
					"user_agent":  "Mozilla/5.0 (iPhone; CPU iPhone OS 16_0 like Mac OS X) AppleWebKit/605.1.15",
					"last_login":  "2023-01-01T11:30:00Z",
					"is_current":  false,
				},
			},
			"total": 2,
			"page":  1,
			"size":  10,
		},
	})
}

// LogoutDevice 下线异常设备
func LogoutDevice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "下线设备成功",
		"data":    nil,
	})
}
