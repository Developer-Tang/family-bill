package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
)

// GetHelpArticles 获取帮助文章列表
func GetHelpArticles(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	// 查询总数
	var total int64
	database.DB.Model(&models.HelpArticle{}).Count(&total)

	// 查询文章列表
	var articles []models.HelpArticle
	database.DB.Where("status = ?", "published").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&articles)

	// 准备响应数据
	var response []gin.H
	for _, article := range articles {
		response = append(response, gin.H{
			"article_id":    article.ArticleID,
			"category_id":   article.CategoryID,
			"category_name": article.Category.Name,
			"title":         article.Title,
			"subtitle":      article.Subtitle,
			"summary":       article.Summary,
			"cover_image":   article.CoverImage,
			"view_count":    article.ViewCount,
			"created_at":    article.CreatedAt,
			"updated_at":    article.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"total":     total,
			"page":      page,
			"page_size": pageSize,
			"articles":  response,
		},
	})
}

// GetHelpArticle 获取帮助文章详情
func GetHelpArticle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	// 查询文章详情
	var article models.HelpArticle
	if err := database.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "文章不存在",
		})
		return
	}

	// 增加浏览量
	database.DB.Model(&article).Update("view_count", gorm.Expr("view_count + ?", 1))

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"article_id":    article.ArticleID,
			"category_id":   article.CategoryID,
			"category_name": article.Category.Name,
			"title":         article.Title,
			"subtitle":      article.Subtitle,
			"summary":       article.Summary,
			"content":       article.Content,
			"cover_image":   article.CoverImage,
			"view_count":    article.ViewCount,
			"like_count":    article.LikeCount,
			"comment_count": article.CommentCount,
			"created_at":    article.CreatedAt,
			"updated_at":    article.UpdatedAt,
			"language":      article.Language,
		},
	})
}

// GetHelpCategories 获取帮助分类
func GetHelpCategories(c *gin.Context) {
	var categories []models.HelpCategory
	database.DB.Find(&categories)

	var response []gin.H
	for _, category := range categories {
		response = append(response, gin.H{
			"category_id":   category.CategoryID,
			"name":          category.Name,
			"description":   category.Description,
			"article_count": len(category.Articles),
			"created_at":    category.CreatedAt,
			"updated_at":    category.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    response,
	})
}

// SearchHelpArticles 搜索帮助文档
func SearchHelpArticles(c *gin.Context) {
	keyword := c.Query("keyword")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	// 查询总数
	var total int64
	database.DB.Model(&models.HelpArticle{}).Where("title LIKE ? OR summary LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Count(&total)

	// 查询文章列表
	var articles []models.HelpArticle
	database.DB.Where("title LIKE ? OR summary LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&articles)

	// 准备响应数据
	var response []gin.H
	for _, article := range articles {
		response = append(response, gin.H{
			"article_id":    article.ArticleID,
			"category_id":   article.CategoryID,
			"category_name": article.Category.Name,
			"title":         article.Title,
			"summary":       article.Summary,
			"cover_image":   article.CoverImage,
			"view_count":    article.ViewCount,
			"created_at":    article.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "搜索成功",
		"data": gin.H{
			"total":     total,
			"page":      page,
			"page_size": pageSize,
			"articles":  response,
		},
	})
}

// SubmitHelpFeedback 提交帮助反馈
func SubmitHelpFeedback(c *gin.Context) {
	var feedback models.Feedback
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data":    gin.H{"error_details": err.Error()},
		})
		return
	}

	// 保存反馈
	if err := database.DB.Create(&feedback).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "提交反馈失败",
			"data":    gin.H{"error_details": err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "反馈提交成功",
		"data": gin.H{
			"feedback_id": feedback.FeedbackID,
			"type":        feedback.Type,
			"created_at":  feedback.CreatedAt,
		},
	})
}
