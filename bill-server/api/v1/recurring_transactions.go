package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
)

// RecurringTransactionRequest 周期记账请求结构体
type RecurringTransactionRequest struct {
	BookID           uint      `json:"book_id" binding:"required"`
	Name             string    `json:"name" binding:"required,max=100"`
	Type             string    `json:"type" binding:"required,oneof=expense income"`
	Amount           float64   `json:"amount" binding:"required,gt=0"`
	Currency         string    `json:"currency" binding:"omitempty,max=10"`
	AccountID        uint      `json:"account_id" binding:"required"`
	CategoryID       uint      `json:"category_id" binding:"required"`
	Frequency        string    `json:"frequency" binding:"required,oneof=daily weekly monthly yearly"`
	Interval         int       `json:"interval" binding:"omitempty,min=1"`
	StartDate        time.Time `json:"start_date" binding:"required"`
	EndDate          time.Time `json:"end_date" binding:"omitempty"`
	NextDate         time.Time `json:"next_date" binding:"required"`
	Memo             string    `json:"memo" binding:"omitempty,max=200"`
	Tags             []uint    `json:"tags" binding:"omitempty"`
	AutoCreate       bool      `json:"auto_create" binding:"omitempty"`
	NotifyBeforeDays int       `json:"notify_before_days" binding:"omitempty,min=0"`
}

// GetRecurringTransactions 获取周期记账列表
func GetRecurringTransactions(c *gin.Context) {
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

	// 查询周期记账列表
	var recurringTransactions []models.RecurringTransaction
	result = database.DB.Where("book_id = ?", bookID).Find(&recurringTransactions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取周期记账列表失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 准备响应数据
	var response []gin.H
	for _, rt := range recurringTransactions {
		response = append(response, gin.H{
			"recurring_id":       rt.RecurringTransactionID,
			"book_id":            rt.BookID,
			"name":               rt.Description,
			"type":               rt.Type,
			"amount":             rt.Amount,
			"currency":           "CNY", // 从账户获取货币类型
			"account_id":         rt.AccountID,
			"category_id":        rt.CategoryID,
			"frequency":          rt.Frequency,
			"interval":           rt.Interval,
			"start_date":         rt.StartDate.Format("2006-01-02"),
			"end_date":           rt.EndDate.Format("2006-01-02"),
			"next_date":          rt.NextExecutedAt.Format("2006-01-02"),
			"memo":               rt.Description,
			"tags":               []uint{}, // 标签功能待实现
			"auto_create":        true,     // 默认值
			"notify_before_days": 0,        // 默认值
			"status":             rt.Status,
			"created_at":         rt.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取周期记账列表成功",
		"data":    response,
	})
}

// CreateRecurringTransaction 创建周期记账
func CreateRecurringTransaction(c *gin.Context) {
	var req RecurringTransactionRequest
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
				"error_details": "您没有权限创建周期记账",
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

	// 创建周期记账
	recurringTransaction := models.RecurringTransaction{
		BookID:         req.BookID,
		UserID:         userID,
		AccountID:      req.AccountID,
		CategoryID:     req.CategoryID,
		Type:           req.Type,
		Amount:         req.Amount,
		Description:    req.Name,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		Frequency:      req.Frequency,
		Interval:       req.Interval,
		NextExecutedAt: req.NextDate,
		Status:         "active",
	}

	result = database.DB.Create(&recurringTransaction)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建周期记账失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建周期记账成功",
		"data": gin.H{
			"recurring_id": recurringTransaction.RecurringTransactionID,
			"name":         recurringTransaction.Description,
			"type":         recurringTransaction.Type,
			"amount":       recurringTransaction.Amount,
			"frequency":    recurringTransaction.Frequency,
			"interval":     recurringTransaction.Interval,
			"next_date":    recurringTransaction.NextExecutedAt.Format("2006-01-02"),
			"status":       recurringTransaction.Status,
		},
	})
}

// UpdateRecurringTransaction 更新周期记账
func UpdateRecurringTransaction(c *gin.Context) {
	// 获取周期记账ID
	recurringIDStr := c.Param("id")
	recurringID, err := strconv.ParseUint(recurringIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "周期记账ID格式错误",
			},
		})
		return
	}

	var req RecurringTransactionRequest
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
				"error_details": "您没有权限更新周期记账",
			},
		})
		return
	}

	// 查询周期记账
	var recurringTransaction models.RecurringTransaction
	result = database.DB.First(&recurringTransaction, recurringID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "周期记账不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 更新周期记账
	recurringTransaction.Description = req.Name
	recurringTransaction.Type = req.Type
	recurringTransaction.Amount = req.Amount
	recurringTransaction.AccountID = req.AccountID
	recurringTransaction.CategoryID = req.CategoryID
	recurringTransaction.Frequency = req.Frequency
	recurringTransaction.Interval = req.Interval
	recurringTransaction.StartDate = req.StartDate
	recurringTransaction.EndDate = req.EndDate
	recurringTransaction.NextExecutedAt = req.NextDate

	result = database.DB.Save(&recurringTransaction)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新周期记账失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新周期记账成功",
		"data": gin.H{
			"recurring_id": recurringTransaction.RecurringTransactionID,
			"name":         recurringTransaction.Description,
			"type":         recurringTransaction.Type,
			"amount":       recurringTransaction.Amount,
			"frequency":    recurringTransaction.Frequency,
			"interval":     recurringTransaction.Interval,
			"next_date":    recurringTransaction.NextExecutedAt.Format("2006-01-02"),
			"status":       recurringTransaction.Status,
		},
	})
}

// DeleteRecurringTransaction 删除周期记账
func DeleteRecurringTransaction(c *gin.Context) {
	// 获取周期记账ID
	recurringIDStr := c.Param("id")
	recurringID, err := strconv.ParseUint(recurringIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "周期记账ID格式错误",
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
				"error_details": "您没有权限删除周期记账",
			},
		})
		return
	}

	// 查询周期记账
	var recurringTransaction models.RecurringTransaction
	result = database.DB.First(&recurringTransaction, recurringID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "周期记账不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 删除周期记账
	result = database.DB.Delete(&recurringTransaction)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除周期记账失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除周期记账成功",
		"data":    nil,
	})
}

// TriggerRecurringTransaction 手动触发周期记账
func TriggerRecurringTransaction(c *gin.Context) {
	// 获取周期记账ID
	recurringIDStr := c.Param("id")
	recurringID, err := strconv.ParseUint(recurringIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "周期记账ID格式错误",
			},
		})
		return
	}

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 查询周期记账
	var recurringTransaction models.RecurringTransaction
	result := database.DB.First(&recurringTransaction, recurringID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "周期记账不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 检查用户是否有权访问该账本
	var bookAccess models.BookAccess
	result = database.DB.Where("book_id = ? AND user_id = ? AND role IN ?", recurringTransaction.BookID, userID, []string{"owner", "editor"}).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限触发该周期记账",
			},
		})
		return
	}

	// 开始事务
	tx := database.DB.Begin()

	// 创建交易记录
	transaction := models.Transaction{
		BookID:      recurringTransaction.BookID,
		UserID:      userID,
		AccountID:   recurringTransaction.AccountID,
		CategoryID:  recurringTransaction.CategoryID,
		Type:        recurringTransaction.Type,
		Amount:      recurringTransaction.Amount,
		Description: recurringTransaction.Description,
		Date:        time.Now(),
		Status:      "active",
		Locked:      false,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建交易记录失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	// 更新周期记账的下次执行时间
	// 这里简化处理，实际项目中需要根据频率和间隔计算下次执行时间
	nextDate := time.Now().AddDate(0, 1, 0) // 假设每月执行一次
	recurringTransaction.NextExecutedAt = nextDate
	recurringTransaction.LastExecutedAt = time.Now()

	if err := tx.Save(&recurringTransaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新周期记账失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	// 提交事务
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "手动触发周期记账成功",
		"data": gin.H{
			"recurring_id":    recurringTransaction.RecurringTransactionID,
			"transaction_ids": []uint{transaction.TransactionID}, // 生成的交易记录ID
			"next_date":       nextDate.Format("2006-01-02"),
		},
	})
}
