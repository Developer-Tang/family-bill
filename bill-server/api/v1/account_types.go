package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
	"github.com/gin-gonic/gin"
)

// GetAccountTypes 获取账户类型列表
func GetAccountTypes(c *gin.Context) {
	var accountTypes []models.AccountType
	result := database.DB.Find(&accountTypes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取账户类型失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 转换为响应格式
	var response []gin.H
	for _, at := range accountTypes {
		response = append(response, gin.H{
			"account_type_id": at.AccountTypeID,
			"name":            at.Name,
			"icon":            at.Icon,
			"color":           at.Color,
			"type":            at.Type,
			"description":     at.Description,
			"created_at":      at.CreatedAt,
			"updated_at":      at.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    response,
	})
}

// AccountTypeRequest 账户类型请求结构体
type AccountTypeRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=50"`
	Icon        string `json:"icon" binding:"omitempty,max=50"`
	Color       string `json:"color" binding:"omitempty,max=20"`
	Type        string `json:"type" binding:"required,oneof=asset liability income expense"`
	Description string `json:"description" binding:"omitempty,max=255"`
}

// CreateAccountType 创建自定义账户类型
func CreateAccountType(c *gin.Context) {
	var req AccountTypeRequest
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

	accountType := models.AccountType{
		Name:        req.Name,
		Icon:        req.Icon,
		Color:       req.Color,
		Type:        req.Type,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := database.DB.Create(&accountType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建账户类型失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data": gin.H{
			"account_type_id": accountType.AccountTypeID,
			"name":            accountType.Name,
			"icon":            accountType.Icon,
			"color":           accountType.Color,
			"type":            accountType.Type,
			"description":     accountType.Description,
			"created_at":      accountType.CreatedAt,
			"updated_at":      accountType.UpdatedAt,
		},
	})
}

// UpdateAccountType 更新账户类型
func UpdateAccountType(c *gin.Context) {
	// 获取账户类型ID
	typeIDStr := c.Param("id")
	typeID, err := strconv.ParseUint(typeIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "账户类型ID格式错误",
			},
		})
		return
	}

	var req AccountTypeRequest
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

	// 查询账户类型
	var accountType models.AccountType
	result := database.DB.First(&accountType, uint(typeID))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "账户类型不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 更新账户类型
	accountType.Name = req.Name
	accountType.Icon = req.Icon
	accountType.Color = req.Color
	accountType.Type = req.Type
	accountType.Description = req.Description
	accountType.UpdatedAt = time.Now()

	if err := database.DB.Save(&accountType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新账户类型失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
		"data": gin.H{
			"account_type_id": accountType.AccountTypeID,
			"name":            accountType.Name,
			"icon":            accountType.Icon,
			"color":           accountType.Color,
			"type":            accountType.Type,
			"description":     accountType.Description,
			"updated_at":      accountType.UpdatedAt,
		},
	})
}

// DeleteAccountType 删除账户类型
func DeleteAccountType(c *gin.Context) {
	// 获取账户类型ID
	typeIDStr := c.Param("id")
	typeID, err := strconv.ParseUint(typeIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "账户类型ID格式错误",
			},
		})
		return
	}

	// 查询账户类型
	var accountType models.AccountType
	result := database.DB.First(&accountType, uint(typeID))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "账户类型不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 检查是否有账户使用该类型
	var accountCount int64
	database.DB.Model(&models.Account{}).Where("account_type_id = ?", uint(typeID)).Count(&accountCount)
	if accountCount > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "删除失败",
			"data": gin.H{
				"error_details": "该账户类型已被使用，无法删除",
			},
		})
		return
	}

	// 删除账户类型
	if err := database.DB.Delete(&accountType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除账户类型失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
		"data":    nil,
	})
}
