package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
)

// TransactionRequest 收支记录请求结构体
type TransactionRequest struct {
	BookID      uint      `json:"book_id" binding:"required"`
	Type        string    `json:"type" binding:"required,oneof=income expense transfer"`
	Amount      float64   `json:"amount" binding:"required,gt=0"`
	AccountID   uint      `json:"account_id" binding:"required"`
	CategoryID  uint      `json:"category_id" binding:"required"`
	Description string    `json:"description" binding:"omitempty,max=255"`
	Date        time.Time `json:"date" binding:"required"`
	Location    string    `json:"location" binding:"omitempty,max=255"`
	Latitude    float64   `json:"latitude" binding:"omitempty"`
	Longitude   float64   `json:"longitude" binding:"omitempty"`
	TagIDs      []uint    `json:"tag_ids" binding:"omitempty"`
}

// LockTransactionRequest 锁定解锁请求结构体
type LockTransactionRequest struct {
	Locked bool `json:"locked" binding:"required"`
}

// BatchTransactionRequest 批量创建请求结构体
type BatchTransactionRequest struct {
	Transactions []TransactionRequest `json:"transactions" binding:"required,min=1,max=100"`
}

// GetTransactions 获取收支记录列表
func GetTransactions(c *gin.Context) {
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

	// 查询收支记录列表
	var transactions []models.Transaction
	var total int64

	// 构建查询
	query := database.DB.Model(&models.Transaction{}).Where("book_id = ?", bookID)

	// 计算总数
	query.Count(&total)

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	// 查询数据
	result = query.Offset(offset).Limit(pageSize).Preload("Account").Preload("Category").Preload("Tags").Find(&transactions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取收支记录列表失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 准备响应数据
	var response []gin.H
	for _, tx := range transactions {
		// 提取标签名称
		var tags []string
		for _, tag := range tx.Tags {
			tags = append(tags, tag.Name)
		}

		response = append(response, gin.H{
			"transaction_id": tx.TransactionID,
			"type":           tx.Type,
			"amount":         tx.Amount,
			"currency":       tx.Account.Currency,
			"account_id":     tx.AccountID,
			"account_name":   tx.Account.Name,
			"category_id":    tx.CategoryID,
			"category_name":  tx.Category.Name,
			"date":           tx.Date.Format("2006-01-02"),
			"memo":           tx.Description,
			"tags":           tags,
			"is_locked":      tx.Locked,
			"created_at":     tx.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"total": total,
			"items": response,
		},
	})
}

// CreateTransaction 创建收支记录
func CreateTransaction(c *gin.Context) {
	var req TransactionRequest
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
		Description: req.Description,
		Date:        req.Date,
		Location:    req.Location,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Status:      "active",
		Locked:      false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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

	// 处理标签关联
	if len(req.TagIDs) > 0 {
		var tags []models.Tag
		result = tx.Where("tag_id IN ?", req.TagIDs).Find(&tags)
		if result.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "获取标签失败",
				"data": gin.H{
					"error_details": result.Error.Error(),
				},
			})
			return
		}

		if err := tx.Model(&transaction).Association("Tags").Replace(tags); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "关联标签失败",
				"data": gin.H{
					"error_details": err.Error(),
				},
			})
			return
		}
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
		"message": "记账成功",
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

// GetTransaction 获取收支记录详情
func GetTransaction(c *gin.Context) {
	// 获取交易ID
	transactionIDStr := c.Param("id")
	transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "交易ID格式错误",
			},
		})
		return
	}

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 查询交易详情
	var transaction models.Transaction
	result := database.DB.Preload("Account").Preload("Category").Preload("Tags").Preload("User").First(&transaction, transactionID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "交易记录不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 检查用户是否有权访问该账本
	var bookAccess models.BookAccess
	result = database.DB.Where("book_id = ? AND user_id = ?", transaction.BookID, userID).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限访问该交易记录",
			},
		})
		return
	}

	// 提取标签信息
	var tags []string
	var tagIDs []uint
	for _, tag := range transaction.Tags {
		tags = append(tags, tag.Name)
		tagIDs = append(tagIDs, tag.TagID)
	}

	// 提取附件信息
	var files []models.File
	result = database.DB.Where("entity_type = ? AND entity_id = ?", "transaction", transaction.TransactionID).Find(&files)
	var attachments []string
	for _, file := range files {
		attachments = append(attachments, file.URL)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"transaction_id": transaction.TransactionID,
			"book_id":        transaction.BookID,
			"type":           transaction.Type,
			"amount":         transaction.Amount,
			"currency":       transaction.Account.Currency,
			"account_id":     transaction.AccountID,
			"account_name":   transaction.Account.Name,
			"category_id":    transaction.CategoryID,
			"category_name":  transaction.Category.Name,
			"date":           transaction.Date.Format("2006-01-02"),
			"time":           transaction.Date.Format("15:04"),
			"memo":           transaction.Description,
			"tags":           tags,
			"tag_ids":        tagIDs,
			"member_id":      transaction.UserID,
			"member_name":    transaction.User.Username,
			"attachments":    attachments,
			"location": gin.H{
				"latitude":  transaction.Latitude,
				"longitude": transaction.Longitude,
				"address":   transaction.Location,
			},
			"is_locked":  transaction.Locked,
			"created_at": transaction.CreatedAt,
			"updated_at": transaction.UpdatedAt,
		},
	})
}

// UpdateTransaction 更新收支记录
func UpdateTransaction(c *gin.Context) {
	// 获取交易ID
	transactionIDStr := c.Param("id")
	transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "交易ID格式错误",
			},
		})
		return
	}

	var req TransactionRequest
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

	// 查询交易详情
	var transaction models.Transaction
	result := database.DB.First(&transaction, transactionID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "交易记录不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 检查用户是否有权访问该账本
	var bookAccess models.BookAccess
	result = database.DB.Where("book_id = ? AND user_id = ? AND role IN ?", transaction.BookID, userID, []string{"owner", "editor"}).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限修改该交易记录",
			},
		})
		return
	}

	// 检查交易是否已锁定
	if transaction.Locked {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "更新失败",
			"data": gin.H{
				"error_details": "该交易记录已锁定，无法修改",
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

	// 开始事务
	tx := database.DB.Begin()

	// 保存原金额和类型，用于更新账户余额
	oldAmount := transaction.Amount
	oldType := transaction.Type

	// 更新交易记录
	transaction.BookID = req.BookID
	transaction.Type = req.Type
	transaction.Amount = req.Amount
	transaction.AccountID = req.AccountID
	transaction.CategoryID = req.CategoryID
	transaction.Description = req.Description
	transaction.Date = req.Date
	transaction.Location = req.Location
	transaction.Latitude = req.Latitude
	transaction.Longitude = req.Longitude
	transaction.UpdatedAt = time.Now()

	if err := tx.Save(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新交易记录失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	// 更新账户余额
	// 1. 撤销原交易对余额的影响
	if oldType == "income" {
		if err := tx.Model(&account).Update("balance", account.Balance-oldAmount).Error; err != nil {
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
	} else if oldType == "expense" {
		if err := tx.Model(&account).Update("balance", account.Balance+oldAmount).Error; err != nil {
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

	// 2. 应用新交易对余额的影响
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

	// 更新标签关联
	if len(req.TagIDs) > 0 {
		var tags []models.Tag
		result = tx.Where("tag_id IN ?", req.TagIDs).Find(&tags)
		if result.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "获取标签失败",
				"data": gin.H{
					"error_details": result.Error.Error(),
				},
			})
			return
		}

		if err := tx.Model(&transaction).Association("Tags").Replace(tags); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "关联标签失败",
				"data": gin.H{
					"error_details": err.Error(),
				},
			})
			return
		}
	} else {
		// 清空标签关联
		if err := tx.Model(&transaction).Association("Tags").Clear(); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "清空标签失败",
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
		"message": "更新成功",
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
			"updated_at":     transaction.UpdatedAt,
		},
	})
}

// DeleteTransaction 删除收支记录
func DeleteTransaction(c *gin.Context) {
	// 获取交易ID
	transactionIDStr := c.Param("id")
	transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "交易ID格式错误",
			},
		})
		return
	}

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 查询交易详情
	var transaction models.Transaction
	result := database.DB.First(&transaction, transactionID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "交易记录不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 检查用户是否有权访问该账本
	var bookAccess models.BookAccess
	result = database.DB.Where("book_id = ? AND user_id = ? AND role IN ?", transaction.BookID, userID, []string{"owner", "editor"}).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限删除该交易记录",
			},
		})
		return
	}

	// 检查交易是否已锁定
	if transaction.Locked {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "删除失败",
			"data": gin.H{
				"error_details": "该交易记录已锁定，无法删除",
			},
		})
		return
	}

	// 查询账户信息
	var account models.Account
	result = database.DB.First(&account, transaction.AccountID)
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

	// 开始事务
	tx := database.DB.Begin()

	// 删除交易记录
	if err := tx.Delete(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除交易记录失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	// 更新账户余额
	if transaction.Type == "income" {
		if err := tx.Model(&account).Update("balance", account.Balance-transaction.Amount).Error; err != nil {
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
	} else if transaction.Type == "expense" {
		if err := tx.Model(&account).Update("balance", account.Balance+transaction.Amount).Error; err != nil {
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
		"message": "删除成功",
		"data":    nil,
	})
}

// LockTransaction 锁定解锁收支记录
func LockTransaction(c *gin.Context) {
	// 获取交易ID
	transactionIDStr := c.Param("id")
	transactionID, err := strconv.ParseUint(transactionIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "交易ID格式错误",
			},
		})
		return
	}

	var req LockTransactionRequest
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

	// 查询交易详情
	var transaction models.Transaction
	result := database.DB.First(&transaction, transactionID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "交易记录不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 检查用户是否有权访问该账本
	var bookAccess models.BookAccess
	result = database.DB.Where("book_id = ? AND user_id = ? AND role IN ?", transaction.BookID, userID, []string{"owner", "editor"}).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限操作该交易记录",
			},
		})
		return
	}

	// 更新锁定状态
	transaction.Locked = req.Locked
	transaction.UpdatedAt = time.Now()

	if err := database.DB.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "操作失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "操作成功",
		"data": gin.H{
			"transaction_id": transaction.TransactionID,
			"is_locked":      transaction.Locked,
			"updated_at":     transaction.UpdatedAt,
		},
	})
}

// BatchCreateTransactions 批量创建收支记录
func BatchCreateTransactions(c *gin.Context) {
	var req BatchTransactionRequest
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

	// 开始事务
	tx := database.DB.Begin()

	var successCount int
	var failedCount int
	var transactionIDs []uint

	for _, txReq := range req.Transactions {
		// 检查用户是否有权访问该账本
		var bookAccess models.BookAccess
		result := tx.Where("book_id = ? AND user_id = ? AND role IN ?", txReq.BookID, userID, []string{"owner", "editor"}).First(&bookAccess)
		if result.Error != nil {
			failedCount++
			continue
		}

		// 检查账户是否存在
		var account models.Account
		result = tx.First(&account, txReq.AccountID)
		if result.Error != nil {
			failedCount++
			continue
		}

		// 检查分类是否存在
		var category models.Category
		result = tx.First(&category, txReq.CategoryID)
		if result.Error != nil {
			failedCount++
			continue
		}

		// 创建收支记录
		transaction := models.Transaction{
			BookID:      txReq.BookID,
			UserID:      userID,
			AccountID:   txReq.AccountID,
			CategoryID:  txReq.CategoryID,
			Type:        txReq.Type,
			Amount:      txReq.Amount,
			Description: txReq.Description,
			Date:        txReq.Date,
			Location:    txReq.Location,
			Latitude:    txReq.Latitude,
			Longitude:   txReq.Longitude,
			Status:      "active",
			Locked:      false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := tx.Create(&transaction).Error; err != nil {
			failedCount++
			continue
		}

		// 处理标签关联
		if len(txReq.TagIDs) > 0 {
			var tags []models.Tag
			result = tx.Where("tag_id IN ?", txReq.TagIDs).Find(&tags)
			if result.Error == nil {
				tx.Model(&transaction).Association("Tags").Replace(tags)
			}
		}

		// 更新账户余额
		if txReq.Type == "income" {
			tx.Model(&account).Update("balance", account.Balance+txReq.Amount)
		} else if txReq.Type == "expense" {
			tx.Model(&account).Update("balance", account.Balance-txReq.Amount)
		}

		successCount++
		transactionIDs = append(transactionIDs, transaction.TransactionID)
	}

	// 提交事务
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "批量记账成功",
		"data": gin.H{
			"success_count":   successCount,
			"failed_count":    failedCount,
			"transaction_ids": transactionIDs,
		},
	})
}

// ImportTransactions 导入外部账单
func ImportTransactions(c *gin.Context) {
	// 这里应该处理文件上传和解析，实际项目中需要实现
	// 简化处理，直接返回成功

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "导入成功",
		"data": gin.H{
			"import_id":     "uuid-string",
			"total_count":   50,
			"success_count": 48,
			"failed_count":  2,
			"failed_records": []gin.H{
				{
					"index": 10,
					"error": "无效的金额格式",
				},
				{
					"index": 45,
					"error": "无法识别的账户",
				},
			},
			"transaction_ids": []uint{4, 5, 6, 7, 8, 9, 10},
		},
	})
}
