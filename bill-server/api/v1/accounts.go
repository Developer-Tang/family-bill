package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
)

// AccountRequest 账户请求结构体
type AccountRequest struct {
	Name           string  `json:"name" binding:"required,min=2,max=50"`
	AccountTypeID  uint    `json:"account_type_id" binding:"required"`
	AccountGroupID uint    `json:"account_group_id" binding:"omitempty"`
	InitialBalance float64 `json:"initial_balance" binding:"omitempty,gte=0"`
	Currency       string  `json:"currency" binding:"required,len=3"`
	Description    string  `json:"description" binding:"omitempty,max=255"`
}

// AdjustBalanceRequest 调整余额请求结构体
type AdjustBalanceRequest struct {
	Amount float64 `json:"amount" binding:"required"`
	Reason string  `json:"reason" binding:"omitempty,max=255"`
}

// UpdateStatusRequest 更新状态请求结构体
type UpdateStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=active inactive"`
}

// GetAccounts 获取账户列表
func GetAccounts(c *gin.Context) {
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

	// 查询账户列表
	var accounts []models.Account
	var total int64

	// 构建查询
	query := database.DB.Model(&models.Account{}).Where("book_id = ?", bookID)

	// 计算总数
	query.Count(&total)

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	offset := (page - 1) * pageSize

	// 查询数据
	result = query.Offset(offset).Limit(pageSize).Preload("AccountType").Preload("AccountGroup").Find(&accounts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取账户列表失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"total": total,
			"items": accounts,
		},
	})
}

// CreateAccount 创建账户
func CreateAccount(c *gin.Context) {
	var req AccountRequest
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
				"error_details": "您没有权限创建账户",
			},
		})
		return
	}

	// 检查账户类型是否存在
	var accountType models.AccountType
	result = database.DB.First(&accountType, req.AccountTypeID)
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

	// 如果指定了账户分组，检查分组是否存在
	if req.AccountGroupID > 0 {
		var accountGroup models.AccountGroup
		result = database.DB.First(&accountGroup, req.AccountGroupID)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "账户分组不存在",
				"data": gin.H{
					"error_details": result.Error.Error(),
				},
			})
			return
		}
	}

	// 创建账户
	account := models.Account{
		BookID:         uint(bookID),
		AccountTypeID:  req.AccountTypeID,
		AccountGroupID: req.AccountGroupID,
		Name:           req.Name,
		InitialBalance: req.InitialBalance,
		Balance:        req.InitialBalance, // 初始余额等于当前余额
		Currency:       req.Currency,
		Description:    req.Description,
		Status:         "active",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := database.DB.Create(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建账户失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data":    account,
	})
}

// GetAccount 获取账户详情
func GetAccount(c *gin.Context) {
	// 获取账户ID
	accountIDStr := c.Param("id")
	accountID, err := strconv.ParseUint(accountIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "账户ID格式错误",
			},
		})
		return
	}

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 查询账户详情
	var account models.Account
	result := database.DB.Preload("AccountType").Preload("AccountGroup").First(&account, accountID)
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

	// 检查用户是否有权访问该账户所属的账本
	var bookAccess models.BookAccess
	result = database.DB.Where("book_id = ? AND user_id = ?", account.BookID, userID).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限访问该账户",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    account,
	})
}

// UpdateAccount 更新账户信息
func UpdateAccount(c *gin.Context) {
	// 获取账户ID
	accountIDStr := c.Param("id")
	accountID, err := strconv.ParseUint(accountIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "账户ID格式错误",
			},
		})
		return
	}

	var req AccountRequest
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

	// 查询账户详情
	var account models.Account
	result := database.DB.First(&account, accountID)
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

	// 检查用户是否有权访问该账户所属的账本
	var bookAccess models.BookAccess
	result = database.DB.Where("book_id = ? AND user_id = ? AND role IN ?", account.BookID, userID, []string{"owner", "editor"}).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限修改该账户",
			},
		})
		return
	}

	// 检查账户类型是否存在
	var accountType models.AccountType
	result = database.DB.First(&accountType, req.AccountTypeID)
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

	// 如果指定了账户分组，检查分组是否存在
	if req.AccountGroupID > 0 {
		var accountGroup models.AccountGroup
		result = database.DB.First(&accountGroup, req.AccountGroupID)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"code":    404,
				"message": "账户分组不存在",
				"data": gin.H{
					"error_details": result.Error.Error(),
				},
			})
			return
		}
	}

	// 更新账户信息
	account.Name = req.Name
	account.AccountTypeID = req.AccountTypeID
	account.AccountGroupID = req.AccountGroupID
	account.Currency = req.Currency
	account.Description = req.Description
	account.UpdatedAt = time.Now()

	// 如果初始余额发生变化，调整当前余额
	if account.InitialBalance != req.InitialBalance {
		balanceDiff := req.InitialBalance - account.InitialBalance
		account.InitialBalance = req.InitialBalance
		account.Balance += balanceDiff
	}

	if err := database.DB.Save(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新账户失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
		"data":    account,
	})
}

// DeleteAccount 删除账户
func DeleteAccount(c *gin.Context) {
	// 获取账户ID
	accountIDStr := c.Param("id")
	accountID, err := strconv.ParseUint(accountIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "账户ID格式错误",
			},
		})
		return
	}

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 查询账户详情
	var account models.Account
	result := database.DB.First(&account, accountID)
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

	// 检查用户是否有权访问该账户所属的账本
	var bookAccess models.BookAccess
	result = database.DB.Where("book_id = ? AND user_id = ? AND role IN ?", account.BookID, userID, []string{"owner", "editor"}).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限删除该账户",
			},
		})
		return
	}

	// 检查账户是否有交易记录
	var transactionCount int64
	database.DB.Model(&models.Transaction{}).Where("account_id = ?", accountID).Count(&transactionCount)
	if transactionCount > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "删除失败",
			"data": gin.H{
				"error_details": "该账户已有交易记录，无法删除",
			},
		})
		return
	}

	// 删除账户
	if err := database.DB.Delete(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除账户失败",
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

// AdjustAccountBalance 调整账户余额
func AdjustAccountBalance(c *gin.Context) {
	// 获取账户ID
	accountIDStr := c.Param("id")
	accountID, err := strconv.ParseUint(accountIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "账户ID格式错误",
			},
		})
		return
	}

	var req AdjustBalanceRequest
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

	// 查询账户详情
	var account models.Account
	result := database.DB.First(&account, accountID)
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

	// 检查用户是否有权访问该账户所属的账本
	var bookAccess models.BookAccess
	result = database.DB.Where("book_id = ? AND user_id = ? AND role IN ?", account.BookID, userID, []string{"owner", "editor"}).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限调整该账户余额",
			},
		})
		return
	}

	// 调整账户余额
	account.Balance += req.Amount
	account.UpdatedAt = time.Now()

	if err := database.DB.Save(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "调整余额失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	// 这里应该记录余额调整日志，实际项目中需要实现

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "余额调整成功",
		"data": gin.H{
			"account_id":  account.AccountID,
			"old_balance": account.Balance - req.Amount,
			"new_balance": account.Balance,
			"amount":      req.Amount,
			"reason":      req.Reason,
		},
	})
}

// UpdateAccountStatus 修改账户状态
func UpdateAccountStatus(c *gin.Context) {
	// 获取账户ID
	accountIDStr := c.Param("id")
	accountID, err := strconv.ParseUint(accountIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "账户ID格式错误",
			},
		})
		return
	}

	var req UpdateStatusRequest
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

	// 查询账户详情
	var account models.Account
	result := database.DB.First(&account, accountID)
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

	// 检查用户是否有权访问该账户所属的账本
	var bookAccess models.BookAccess
	result = database.DB.Where("book_id = ? AND user_id = ? AND role IN ?", account.BookID, userID, []string{"owner", "editor"}).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限修改该账户状态",
			},
		})
		return
	}

	// 更新账户状态
	account.Status = req.Status
	account.UpdatedAt = time.Now()

	if err := database.DB.Save(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新状态失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "状态更新成功",
		"data": gin.H{
			"account_id": account.AccountID,
			"status":     account.Status,
		},
	})
}
