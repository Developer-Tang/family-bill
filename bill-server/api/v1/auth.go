package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 用户登录
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"data": gin.H{
			"user": gin.H{
				"user_id":    1,
				"username":   "张三",
				"email":      "user@example.com",
				"avatar":     "https://example.com/avatar.jpg",
				"created_at": "2023-01-01T00:00:00Z",
			},
			"tokens": gin.H{
				"access_token":             "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
				"refresh_token":            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
				"access_token_expires_in":  3600,
				"refresh_token_expires_in": 86400,
			},
			"default_book_id": 1,
		},
	})
}

// Logout 用户登出
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登出成功",
		"data": gin.H{
			"logout_time": "2023-01-05T12:30:00Z",
		},
	})
}

// Register 用户注册
func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
		"data": gin.H{
			"user_id":              1,
			"username":             "张三",
			"email":                "user@example.com",
			"created_at":           "2023-01-05T12:30:00Z",
			"default_book_created": true,
			"verification_status":  "pending",
		},
	})
}

// RefreshToken 刷新令牌
func RefreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "令牌刷新成功",
		"data": gin.H{
			"access_token":             "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
			"access_token_expires_in":  3600,
			"refresh_token":            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
			"refresh_token_expires_in": 86400,
		},
	})
}

// VerifyEmail 邮箱验证
func VerifyEmail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "邮箱验证成功",
		"data": gin.H{
			"email":       "user@example.com",
			"verified_at": "2023-01-05T12:30:00Z",
		},
	})
}

// ResetPassword 重置密码
func ResetPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "密码重置成功",
		"data": gin.H{
			"email":    "user@example.com",
			"reset_at": "2023-01-05T12:30:00Z",
		},
	})
}

// SendCode 发送验证码
func SendCode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "验证码发送成功",
		"data": gin.H{
			"contact":    "user@example.com",
			"type":       "email",
			"expires_in": 300,
		},
	})
}
