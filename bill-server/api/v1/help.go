package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHelpArticles 获取帮助文章列表
func GetHelpArticles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "获取成功",
		"data": gin.H{
			"total": 50,
			"page": 1,
			"page_size": 20,
			"articles": []gin.H{
				{
					"article_id": 1,
					"category_id": 1,
					"category_name": "入门指南",
					"title": "如何开始使用家庭记账系统",
					"subtitle": "快速上手，开始记录您的收支",
					"summary": "本文将指导您如何注册账号、创建账本并开始记录第一笔交易...",
					"cover_image": "https://example.com/help/cover1.jpg",
					"view_count": 1500,
					"created_at": "2023-01-01T00:00:00Z",
					"updated_at": "2023-01-02T00:00:00Z",
				},
			},
		},
	})
}

// GetHelpArticle 获取帮助文章详情
func GetHelpArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "获取成功",
		"data": gin.H{
			"article_id": 1,
			"category_id": 1,
			"category_name": "入门指南",
			"title": "如何开始使用家庭记账系统",
			"subtitle": "快速上手，开始记录您的收支",
			"summary": "本文将指导您如何注册账号、创建账本并开始记录第一笔交易...",
			"content": "<h1>欢迎使用家庭记账系统</h1><p>...详细内容...</p>",
			"cover_image": "https://example.com/help/cover1.jpg",
			"view_count": 1500,
			"like_count": 150,
			"comment_count": 20,
			"created_at": "2023-01-01T00:00:00Z",
			"updated_at": "2023-01-02T00:00:00Z",
			"language": "zh-CN",
		},
	})
}

// GetHelpCategories 获取帮助分类
func GetHelpCategories(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "获取成功",
		"data": []gin.H{
			{
				"category_id": 1,
				"name": "入门指南",
				"description": "系统使用入门教程",
				"article_count": 10,
				"created_at": "2023-01-01T00:00:00Z",
				"updated_at": "2023-01-02T00:00:00Z",
			},
		},
	})
}

// SearchHelpArticles 搜索帮助文档
func SearchHelpArticles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "搜索成功",
		"data": gin.H{
			"total": 5,
			"page": 1,
			"page_size": 20,
			"articles": []gin.H{
				{
					"article_id": 1,
					"category_id": 1,
					"category_name": "入门指南",
					"title": "如何开始使用家庭记账系统",
					"summary": "本文将指导您如何注册账号、创建账本并开始记录第一笔交易...",
					"cover_image": "https://example.com/help/cover1.jpg",
					"view_count": 1500,
					"created_at": "2023-01-01T00:00:00Z",
				},
			},
		},
	})
}

// SubmitHelpFeedback 提交帮助反馈
func SubmitHelpFeedback(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "反馈提交成功",
		"data": gin.H{
			"feedback_id": "feedback_123",
			"type": "suggestion",
			"created_at": "2023-01-05T12:30:00Z",
		},
	})
}