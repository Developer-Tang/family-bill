package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTags 获取标签列表
func GetTags(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取标签列表成功",
		"data": []gin.H{
			{
				"tag_id":     1,
				"book_id":    1,
				"name":       "日常",
				"color":      "#FF6B6B",
				"count":      50,
				"created_at": "2023-01-01T12:00:00Z",
			},
			{
				"tag_id":     2,
				"book_id":    1,
				"name":       "午餐",
				"color":      "#4ECDC4",
				"count":      25,
				"created_at": "2023-01-01T12:00:00Z",
			},
			{
				"tag_id":     3,
				"book_id":    1,
				"name":       "交通",
				"color":      "#45B7D1",
				"count":      15,
				"created_at": "2023-01-01T12:00:00Z",
			},
			{
				"tag_id":     4,
				"book_id":    1,
				"name":       "晚餐",
				"color":      "#FFA07A",
				"count":      20,
				"created_at": "2023-01-05T18:00:00Z",
			},
		},
	})
}

// CreateTag 创建标签
func CreateTag(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建标签成功",
		"data": gin.H{
			"tag_id":     5,
			"book_id":    1,
			"name":       "购物",
			"color":      "#98D8C8",
			"count":      0,
			"created_at": "2023-01-05T19:00:00Z",
		},
	})
}

// UpdateTag 更新标签
func UpdateTag(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新标签成功",
		"data": gin.H{
			"tag_id":     1,
			"book_id":    1,
			"name":       "日常（更新）",
			"color":      "#FF8C42",
			"count":      50,
			"updated_at": "2023-01-05T19:30:00Z",
		},
	})
}

// DeleteTag 删除标签
func DeleteTag(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除标签成功",
		"data":    nil,
	})
}
