package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
)

// BookRequest 账本请求结构体
type BookRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=50"`
	Currency    string `json:"currency" binding:"required,len=3"`
	Description string `json:"description" binding:"omitempty,max=255"`
}

// AddBookMemberRequest 添加账本成员请求结构体
type AddBookMemberRequest struct {
	UserID uint   `json:"user_id" binding:"required"`
	Role   string `json:"role" binding:"required,oneof=owner editor viewer"`
}

// UpdateBookMemberRoleRequest 更新账本成员角色请求结构体
type UpdateBookMemberRoleRequest struct {
	Role string `json:"role" binding:"required,oneof=owner editor viewer"`
}

// GetBooks 获取账本列表
func GetBooks(c *gin.Context) {
	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 查询用户有权访问的所有账本
	var bookAccess []models.BookAccess
	result := database.DB.Where("user_id = ?", userID).Find(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取账本列表失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 提取账本ID列表
	var bookIDs []uint
	for _, access := range bookAccess {
		bookIDs = append(bookIDs, access.BookID)
	}

	// 查询账本详细信息
	var books []models.Book
	if len(bookIDs) > 0 {
		result = database.DB.Where("book_id IN ?", bookIDs).Find(&books)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "获取账本详情失败",
				"data": gin.H{
					"error_details": result.Error.Error(),
				},
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    books,
	})
}

// CreateBook 创建账本
func CreateBook(c *gin.Context) {
	var req BookRequest
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

	// 创建账本
	book := models.Book{
		Name:        req.Name,
		CreatorID:   userID,
		Currency:    req.Currency,
		Description: req.Description,
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// 开始事务
	tx := database.DB.Begin()
	if err := tx.Create(&book).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "创建账本失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	// 添加创建者为账本成员，角色为owner
	bookAccess := models.BookAccess{
		BookID:    book.BookID,
		UserID:    userID,
		Role:      "owner",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := tx.Create(&bookAccess).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "添加账本成员失败",
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
		"message": "创建成功",
		"data":    book,
	})
}

// GetBook 获取账本详情
func GetBook(c *gin.Context) {
	// 获取账本ID
	bookID := c.Param("id")

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

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

	// 查询账本详情
	var book models.Book
	result = database.DB.First(&book, bookID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "账本不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    book,
	})
}

// UpdateBook 更新账本信息
func UpdateBook(c *gin.Context) {
	// 获取账本ID
	bookID := c.Param("id")

	var req BookRequest
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

	// 检查用户是否是账本所有者或编辑者
	var bookAccess models.BookAccess
	result := database.DB.Where("book_id = ? AND user_id = ? AND role IN ?", bookID, userID, []string{"owner", "editor"}).First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限修改该账本",
			},
		})
		return
	}

	// 更新账本信息
	var book models.Book
	result = database.DB.First(&book, bookID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "账本不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	book.Name = req.Name
	book.Currency = req.Currency
	book.Description = req.Description
	book.UpdatedAt = time.Now()

	if err := database.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
		"data":    book,
	})
}

// DeleteBook 删除账本
func DeleteBook(c *gin.Context) {
	// 获取账本ID
	bookID := c.Param("id")

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 检查用户是否是账本所有者
	var bookAccess models.BookAccess
	result := database.DB.Where("book_id = ? AND user_id = ? AND role = ?", bookID, userID, "owner").First(&bookAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "只有所有者可以删除账本",
			},
		})
		return
	}

	// 删除账本（软删除）
	if err := database.DB.Delete(&models.Book{}, bookID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除失败",
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

// GetBookMembers 获取账本成员列表
func GetBookMembers(c *gin.Context) {
	// 获取账本ID
	bookID := c.Param("id")

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

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

	// 查询账本成员列表
	var bookMembers []models.BookAccess
	result = database.DB.Where("book_id = ?", bookID).Preload("User").Find(&bookMembers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取账本成员列表失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 准备响应数据
	var response []gin.H
	for _, member := range bookMembers {
		response = append(response, gin.H{
			"user_id":   member.User.UserID,
			"username":  member.User.Username,
			"email":     member.User.Email,
			"avatar":    member.User.Avatar,
			"role":      member.Role,
			"joined_at": member.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    response,
	})
}

// AddBookMember 添加账本成员
func AddBookMember(c *gin.Context) {
	// 获取账本ID
	bookIDStr := c.Param("id")

	// 转换账本ID为uint
	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的账本ID",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	var req AddBookMemberRequest
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

	// 检查当前用户是否是账本所有者或编辑者
	var currentAccess models.BookAccess
	result := database.DB.Where("book_id = ? AND user_id = ? AND role IN ?", bookID, userID, []string{"owner", "editor"}).First(&currentAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限添加账本成员",
			},
		})
		return
	}

	// 检查要添加的用户是否存在
	var user models.User
	result = database.DB.First(&user, req.UserID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "用户不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 检查用户是否已加入该账本
	var existingAccess models.BookAccess
	result = database.DB.Where("book_id = ? AND user_id = ?", bookID, req.UserID).First(&existingAccess)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "添加失败",
			"data": gin.H{
				"error_details": "该用户已加入该账本",
			},
		})
		return
	}

	// 添加账本成员
	bookAccess := models.BookAccess{
		BookID:    uint(bookID),
		UserID:    req.UserID,
		Role:      req.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := database.DB.Create(&bookAccess).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "添加失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "添加成功",
		"data": gin.H{
			"book_id": bookID,
			"user_id": req.UserID,
			"role":    req.Role,
		},
	})
}

// RemoveBookMember 移除账本成员
func RemoveBookMember(c *gin.Context) {
	// 获取账本ID和用户ID
	bookID := c.Param("id")
	removeUserID := c.Param("userId")

	// 从上下文获取当前用户ID - 实际项目中应该从JWT令牌中解析
	currentUserID := uint(1)

	// 检查当前用户是否是账本所有者或编辑者
	var currentAccess models.BookAccess
	result := database.DB.Where("book_id = ? AND user_id = ? AND role IN ?", bookID, currentUserID, []string{"owner", "editor"}).First(&currentAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "您没有权限移除账本成员",
			},
		})
		return
	}

	// 检查要移除的用户是否是账本成员
	var removeAccess models.BookAccess
	result = database.DB.Where("book_id = ? AND user_id = ?", bookID, removeUserID).First(&removeAccess)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "成员不存在",
			"data": gin.H{
				"error_details": "要移除的用户不是该账本成员",
			},
		})
		return
	}

	// 不能移除所有者
	if removeAccess.Role == "owner" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "移除失败",
			"data": gin.H{
				"error_details": "不能移除账本所有者",
			},
		})
		return
	}

	// 移除账本成员
	if err := database.DB.Delete(&removeAccess).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "移除失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "移除成功",
		"data":    nil,
	})
}

// UpdateBookMemberRole 更新账本成员角色
func UpdateBookMemberRole(c *gin.Context) {
	// 获取账本ID和用户ID
	bookID := c.Param("id")
	userID := c.Param("userId")

	var req UpdateBookMemberRoleRequest
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

	// 从上下文获取当前用户ID - 实际项目中应该从JWT令牌中解析
	currentUserID := uint(1)

	// 检查当前用户是否是账本所有者
	var currentAccess models.BookAccess
	result := database.DB.Where("book_id = ? AND user_id = ? AND role = ?", bookID, currentUserID, "owner").First(&currentAccess)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "权限不足",
			"data": gin.H{
				"error_details": "只有所有者可以更新成员角色",
			},
		})
		return
	}

	// 更新账本成员角色
	result = database.DB.Model(&models.BookAccess{}).Where("book_id = ? AND user_id = ?", bookID, userID).Update("role", req.Role)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "成员不存在",
			"data": gin.H{
				"error_details": "要更新的用户不是该账本成员",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
		"data": gin.H{
			"book_id": bookID,
			"user_id": userID,
			"role":    req.Role,
		},
	})
}

// SetDefaultBook 设置默认账本
func SetDefaultBook(c *gin.Context) {
	// 获取账本ID
	bookID := c.Param("id")

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

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

	// 更新用户默认账本
	result = database.DB.Model(&models.User{}).Where("user_id = ?", userID).Update("default_book_id", bookID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "设置失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "设置成功",
		"data": gin.H{
			"default_book_id": bookID,
		},
	})
}
