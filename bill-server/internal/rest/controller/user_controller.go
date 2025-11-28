package controller

import (
	"strings"

	"github.com/family-bill/bill-server/internal/rest/middleware"
	"github.com/family-bill/bill-server/internal/service"
	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct {
	userService *service.UserService
}

// NewUserController 创建用户控制器实例
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// parseValidationError 解析验证错误，只返回具体的错误信息
func parseValidationError(err error) string {
	// 从原始错误信息中提取Error:后面的内容
	errStr := err.Error()
	if idx := strings.Index(errStr, "Error:"); idx != -1 {
		return strings.TrimSpace(errStr[idx+6:])
	}

	return errStr
}

// Register 用户注册
// @Summary 用户注册
// @Description 创建新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "注册请求"
// @Success 200 {object} utils.Response "成功响应"
// @Failure 400 {object} utils.Response "请求参数错误"
// @Router /users/register [post]
func (c *UserController) Register(ctx *gin.Context) {

	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 解析验证错误，只返回具体的错误信息
		errMsg := parseValidationError(err)
		middleware.ErrorResponse(ctx, 400, errMsg)
		return
	}

	user, err := c.userService.Register(req.Username, req.Email, req.Password)
	if err != nil {
		middleware.ErrorResponse(ctx, 400, "注册失败")
		return
	}

	// 构造响应数据
	responseData := gin.H{
		"id":            user.ID,
		"username":      user.Username,
		"email":         user.Email,
		"avatar":        user.Avatar,
		"role":          user.HaveRole,
		"access_token":  "",
		"refresh_token": "",
	}
	middleware.SuccessResponse(ctx, responseData)
}
