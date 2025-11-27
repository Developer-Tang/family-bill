package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProfile 获取用户个人信息
func GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"user_id":    1,
			"username":  "张三",
			"email":     "user@example.com",
			"phone":     "138****8000",
			"avatar":    "https://example.com/avatar.jpg",
			"created_at": "2023-01-01T00:00:00Z",
			"last_login": "2023-01-05T12:30:00Z",
			"status":    "active",
			"email_verified": true,
			"phone_verified": false,
			"default_book_id": 1
		}
	})
}

// UpdateProfile 更新用户个人信息
func UpdateProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
		"data": gin.H{
			"user_id":    1,
			"username":  "张三",
			"email":     "user@example.com",
			"phone":     "138****8000",
			"avatar":    "https://example.com/avatar.jpg",
			"updated_at": "2023-01-05T12:30:00Z"
		}
	})
}

// UpdateAvatar 更新用户头像
func UpdateAvatar(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "头像更新成功",
		"data": gin.H{
			"user_id": 1,
			"avatar":  "https://example.com/new_avatar.jpg",
			"updated_at": "2023-01-05T12:30:00Z"
		}
	})
}