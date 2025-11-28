package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
)

// QuickTransactionRequest 快速记账请求结构体
type QuickTransactionRequest struct {
	BookID     uint    `json:"book_id" binding:"required"`
	Type       string  `json:"type" binding:"required,oneof=expense income"`
	Amount     float64 `json:"amount" binding:"required,gt=0"`
	AccountID  uint    `json:"account_id" binding:"required"`
	CategoryID uint    `json:"category_id" binding:"required"`
	Memo       string  `json:"memo" binding:"omitempty,max=200"`
	Date       string  `json:"date" binding:"omitempty"`
	Time       string  `json:"time" binding:"omitempty"`
	Location   string  `json:"location" binding:"omitempty,max=255"`
	Latitude   float64 `json:"latitude" binding:"omitempty"`
	Longitude  float64 `json:"longitude" binding:"omitempty"`
}

// TransactionTemplateRequest 记账模板请求结构体
type TransactionTemplateRequest struct {
	BookID     uint    `json:"book_id" binding:"required"`
	Name       string  `json:"name" binding:"required,max=100"`
	Type       string  `json:"type" binding:"required,oneof=expense income"`
	Amount     float64 `json:"amount" binding:"omitempty,gt=0"`
	AccountID  uint    `json:"account_id" binding:"required"`
	CategoryID uint    `json:"category_id" binding:"required"`
	Memo       string  `json:"memo" binding:"omitempty,max=200"`
	Tags       []uint  `json:"tags" binding:"omitempty"`
}

// CreateQuickTransaction 创建快速记账记录
func CreateQuickTransaction(c *gin.Context) {
	var req QuickTransactionRequest
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
				"error_details": "您没有权限创建收支记录",
			},
		})
		return
	}

	// 检查账户是否存在
	var account models.Account
	result = database.DB.First(&account, req.AccountID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "账户不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 检查分类是否存在
	var category models.Category
	result = database.DB.First(&category, req.CategoryID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "分类不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 解析日期时间
	var transactionDate time.Time
	var transactionTime time.Time
	var err error

	if req.Date == "" {
		transactionDate = time.Now()
	} else {
		transactionDate, err = time.Parse("2006-01-02", req.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "请求参数错误",
				"data": gin.H{
					"error_details": "无效的日期格式",
				},
			})
			return
		}
	}

	if req.Time == "" {
		transactionTime = time.Now()
	} else {
		transactionTime, err = time.Parse("15:04", req.Time)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "请求参数错误",
				"data": gin.H{
					"error_details": "无效的时间格式",
				},
			})
			return
		}
	}

	// 合并日期和时间
	transactionDateTime := time.Date(
		transactionDate.Year(),
		transactionDate.Month(),
		transactionDate.Day(),
		transactionTime.Hour(),
		transactionTime.Minute(),
		0,
		0,
		transactionDate.Location(),
	)

	// 开始事务
	tx := database.DB.Begin()

	// 创建收支记录
	transaction := models.Transaction{
		BookID:      req.BookID,
		UserID:      userID,
		AccountID:   req.AccountID,
		CategoryID:  req.CategoryID,
		Type:        req.Type,
		Amount:      req.Amount,
		Description: req.Memo,
		Date:        transactionDateTime,
		Location:    req.Location,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Status:      "active",
		Locked:      false,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建收支记录失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	// 更新账户余额
	if req.Type == "income" {
		if err := tx.Model(&account).Update("balance", account.Balance+req.Amount).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "更新账户余额失败",
				"data": gin.H{
					"error_details": err.Error(),
				},
			})
			return
		}
	} else if req.Type == "expense" {
		if err := tx.Model(&account).Update("balance", account.Balance-req.Amount).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "更新账户余额失败",
				"data": gin.H{
					"error_details": err.Error(),
				},
			})
			return
		}
	}

	// 提交事务
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "快速记账成功",
		"data": gin.H{
			"transaction_id": transaction.TransactionID,
			"type":           transaction.Type,
			"amount":         transaction.Amount,
			"currency":       account.Currency,
			"account_name":   account.Name,
			"category_name":  category.Name,
			"date":           transaction.Date.Format("2006-01-02"),
			"time":           transaction.Date.Format("15:04"),
			"memo":           transaction.Description,
			"created_at":     transaction.CreatedAt,
		},
	})
}

// GetTransactionTemplates 获取记账模板
func GetTransactionTemplates(c *gin.Context) {
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

	// 查询记账模板列表
	var templates []models.TransactionTemplate
	result = database.DB.Where("book_id = ?", bookID).Find(&templates)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取记账模板失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 准备响应数据
	var response []gin.H
	for _, template := range templates {
		response = append(response, gin.H{
			"template_id": template.TransactionTemplateID,
			"book_id":     template.BookID,
			"name":        template.Name,
			"type":        template.Type,
			"amount":      template.Amount,
			"account_id":  template.AccountID,
			"category_id": template.CategoryID,
			"tags":        []uint{}, // 标签功能待实现
			"memo":        template.Description,
			"created_at":  template.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取记账模板成功",
		"data":    response,
	})
}

// CreateTransactionTemplate 创建记账模板
func CreateTransactionTemplate(c *gin.Context) {
	var req TransactionTemplateRequest
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
				"error_details": "您没有权限创建记账模板",
			},
		})
		return
	}

	// 检查账户是否存在
	var account models.Account
	result = database.DB.First(&account, req.AccountID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "账户不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 检查分类是否存在
	var category models.Category
	result = database.DB.First(&category, req.CategoryID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "分类不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 创建记账模板
	template := models.TransactionTemplate{
		BookID:      req.BookID,
		UserID:      userID,
		Name:        req.Name,
		AccountID:   req.AccountID,
		CategoryID:  req.CategoryID,
		Type:        req.Type,
		Amount:      req.Amount,
		Description: req.Memo,
		// Tags:        strings.Join(tagStrings, ","), // 标签功能待实现
	}

	result = database.DB.Create(&template)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建记账模板失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建记账模板成功",
		"data": gin.H{
			"template_id": template.TransactionTemplateID,
			"book_id":     template.BookID,
			"name":        template.Name,
			"type":        template.Type,
			"amount":      template.Amount,
			"account_id":  template.AccountID,
			"category_id": template.CategoryID,
			"tags":        []uint{}, // 标签功能待实现
			"memo":        template.Description,
			"created_at":  template.CreatedAt,
		},
	})
}

// UpdateTransactionTemplate 更新记账模板
func UpdateTransactionTemplate(c *gin.Context) {
	// 获取模板ID
	templateIDStr := c.Param("id")
	templateID, err := strconv.ParseUint(templateIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "模板ID格式错误",
			},
		})
		return
	}

	var req TransactionTemplateRequest
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
				"error_details": "您没有权限更新记账模板",
			},
		})
		return
	}

	// 查询模板
	var template models.TransactionTemplate
	result = database.DB.First(&template, templateID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "记账模板不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 更新模板
	template.Name = req.Name
	template.Type = req.Type
	template.Amount = req.Amount
	template.AccountID = req.AccountID
	template.CategoryID = req.CategoryID
	template.Description = req.Memo
	// template.Tags = strings.Join(tagStrings, ",") // 标签功能待实现

	result = database.DB.Save(&template)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新记账模板失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新记账模板成功",
		"data": gin.H{
			"template_id": template.TransactionTemplateID,
			"book_id":     template.BookID,
			"name":        template.Name,
			"type":        template.Type,
			"amount":      template.Amount,
			"account_id":  template.AccountID,
			"category_id": template.CategoryID,
			"tags":        []uint{}, // 标签功能待实现
			"memo":        template.Description,
			"updated_at":  template.UpdatedAt,
		},
	})
}

// DeleteTransactionTemplate 删除记账模板
func DeleteTransactionTemplate(c *gin.Context) {
	// 获取模板ID
	templateIDStr := c.Param("id")
	templateID, err := strconv.ParseUint(templateIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "模板ID格式错误",
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
				"error_details": "您没有权限删除记账模板",
			},
		})
		return
	}

	// 查询模板
	var template models.TransactionTemplate
	result = database.DB.First(&template, templateID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "记账模板不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 删除模板
	result = database.DB.Delete(&template)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除记账模板失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除记账模板成功",
		"data":    nil,
	})
}
