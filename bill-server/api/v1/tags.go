package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
)

// TagRequest 标签请求结构体
type TagRequest struct {
	BookID    uint   `json:"book_id" binding:"required"`
	Name      string `json:"name" binding:"required,max=50"`
	Color     string `json:"color" binding:"omitempty,max=20"`
	Count     int    `json:"count" binding:"omitempty,min=0"`
	CreatedAt string `json:"created_at" binding:"omitempty"`
}

// GetTags 获取标签列表
func GetTags(c *gin.Context) {
	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 获取账本ID
	bookIDStr := c.Query("book_id")
	if bookIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "账本ID不能为空",
			},
		})
		return
	}

	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "账本ID格式错误",
			},
		})
		return
	}

	// 检查用户是否有权访问该账本
	var bookAccess models.BookAccess
	result := database.DB.Where("book_id = ? AND user_id = ?", bookID, userID).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限访问该账本",
			},
		})
		return
	}

	// 查询标签列表
	var tags []models.Tag
	result = database.DB.Find(&tags)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取标签列表失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 准备响应数据
	var response []gin.H
	for _, tag := range tags {
		// 计算标签使用次数
		var count int64
		database.DB.Model(&models.Transaction{}).Where("? = ANY(tags)", tag.TagID).Count(&count)

		response = append(response, gin.H{
			"tag_id":     tag.TagID,
			"book_id":    bookID,
			"name":       tag.Name,
			"color":      tag.Color,
			"count":      count,
			"created_at": tag.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取标签列表成功",
		"data":    response,
	})
}

// CreateTag 创建标签
func CreateTag(c *gin.Context) {
	var req TagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 检查用户是否有权访问该账本
	var bookAccess models.BookAccess
	result := database.DB.Where("book_id = ? AND user_id = ? AND role IN ?", req.BookID, userID, []string{"owner", "editor"}).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限创建标签",
			},
		})
		return
	}

	// 创建标签
	tag := models.Tag{
		Name:        req.Name,
		Color:       req.Color,
		Description: "",
	}

	result = database.DB.Create(&tag)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建标签失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建标签成功",
		"data": gin.H{
			"tag_id":     tag.TagID,
			"book_id":    req.BookID,
			"name":       tag.Name,
			"color":      tag.Color,
			"count":      0,
			"created_at": tag.CreatedAt,
		},
	})
}

// UpdateTag 更新标签
func UpdateTag(c *gin.Context) {
	// 获取标签ID
	tagIDStr := c.Param("id")
	tagID, err := strconv.ParseUint(tagIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "标签ID格式错误",
			},
		})
		return
	}

	var req TagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 检查用户是否有权访问该账本
	var bookAccess models.BookAccess
	result := database.DB.Where("book_id = ? AND user_id = ? AND role IN ?", req.BookID, userID, []string{"owner", "editor"}).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限更新标签",
			},
		})
		return
	}

	// 查询标签
	var tag models.Tag
	result = database.DB.First(&tag, tagID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "标签不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 更新标签
	tag.Name = req.Name
	tag.Color = req.Color

	result = database.DB.Save(&tag)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新标签失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 计算标签使用次数
	var count int64
	database.DB.Model(&models.Transaction{}).Where("? = ANY(tags)", tag.TagID).Count(&count)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新标签成功",
		"data": gin.H{
			"tag_id":     tag.TagID,
			"book_id":    req.BookID,
			"name":       tag.Name,
			"color":      tag.Color,
			"count":      count,
			"updated_at": tag.UpdatedAt,
		},
	})
}

// DeleteTag 删除标签
func DeleteTag(c *gin.Context) {
	// 获取标签ID
	tagIDStr := c.Param("id")
	tagID, err := strconv.ParseUint(tagIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "标签ID格式错误",
			},
		})
		return
	}

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 获取账本ID
	bookIDStr := c.Query("book_id")
	if bookIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "账本ID不能为空",
			},
		})
		return
	}

	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "账本ID格式错误",
			},
		})
		return
	}

	// 检查用户是否有权访问该账本
	var bookAccess models.BookAccess
	result := database.DB.Where("book_id = ? AND user_id = ? AND role IN ?", bookID, userID, []string{"owner", "editor"}).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限删除标签",
			},
		})
		return
	}

	// 查询标签
	var tag models.Tag
	result = database.DB.First(&tag, tagID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "标签不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 删除标签
	result = database.DB.Delete(&tag)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除标签失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除标签成功",
		"data":    nil,
	})
}
