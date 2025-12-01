package rest

import (
	"log"
	"time"

	"github.com/family-bill/bill-server/internal/constant"
	"github.com/family-bill/bill-server/internal/model"
	"github.com/family-bill/bill-server/internal/model/bo"
	"github.com/family-bill/bill-server/internal/model/vo"
	"github.com/family-bill/bill-server/internal/service"
	"github.com/family-bill/bill-server/internal/util"
	"github.com/gin-gonic/gin"
)

// UserRest 用户控制器
type UserRest struct {
	userService *service.UserService
}

// UserRest 创建用户控制器实例
func InitUserRest() *UserRest {
	return &UserRest{
		userService: service.NewUserService(),
	}
}

// Register 用户注册
// @Summary 用户注册
// @Description 创建新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body bo.RegisterBo true "注册请求"
// @Success 200 {object} vo.R{data=vo.AuthVO} "成功响应"
// @Failure 400 {object} vo.R "请求参数错误"
// @Router /users/register [post]
func (c *UserRest) Register(ctx *gin.Context) {

	var req bo.RegisterBo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 解析验证错误，只返回具体的错误信息
		errMsg := util.ParseValidationError(err)
		util.RespFail(ctx, 400, errMsg)
		return
	}

	user, err := c.userService.Register(req.Username, req.Email, req.Password)
	if err != nil {
		util.RespFail(ctx, 400, err.Error())
		return
	}

	// 构建AuthVO
	authVO := BuildAuthVO(ctx, user)

	util.RespOk(ctx, authVO)
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录，返回token
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body bo.LoginBo true "登录请求"
// @Success 200 {object} vo.R{data=vo.AuthVO} "成功响应"
// @Failure 400 {object} vo.R "请求参数错误"
// @Failure 401 {object} vo.R "用户名或密码错误"
// @Router /users/login [post]
func (c *UserRest) Login(ctx *gin.Context) {
	var req bo.LoginBo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 解析验证错误，只返回具体的错误信息
		errMsg := util.ParseValidationError(err)
		util.RespFail(ctx, 400, errMsg)
		return
	}

	user, err := c.userService.Login(req.Username, req.Password)
	if err != nil {
		util.RespFail(ctx, 401, "用户名或密码错误")
		return
	}

	// 构建AuthVO
	authVO := BuildAuthVO(ctx, user)

	util.RespOk(ctx, authVO)
}

// RefreshToken 刷新token
// @Summary 刷新token
// @Description 使用refresh_token刷新access_token
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body bo.RefreshTokenBo true "刷新token请求"
// @Success 200 {object} vo.R{data=vo.AuthVO} "成功响应"
// @Failure 400 {object} vo.R "请求参数错误"
// @Failure 401 {object} vo.R "无效的refresh_token"
// @Router /users/refresh-token [post]
func (c *UserRest) RefreshToken(ctx *gin.Context) {
	var req bo.RefreshTokenBo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 解析验证错误，只返回具体的错误信息
		errMsg := util.ParseValidationError(err)
		util.RespFail(ctx, 400, errMsg)
		return
	}

	// 解析refresh_token
	claims, err := util.ParseToken(req.RefreshToken)
	if err != nil {
		util.RespFail(ctx, 401, "无效的refresh_token")
		return
	}

	// 获取用户信息
	user, err := c.userService.FindUserById(claims.UserID)
	if err != nil {
		util.RespFail(ctx, 401, "用户不存在")
		return
	}

	// 构建AuthVO
	authVO := BuildAuthVO(ctx, user)

	util.RespOk(ctx, authVO)
}

// GetUserInfo 获取用户信息
// @Summary 获取用户信息
// @Description 根据用户ID查询用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param userID path int true "用户ID"
// @Success 200 {object} vo.R{data=vo.UserVO} "成功响应"
// @Failure 401 {object} vo.R "用户不存在"
// @Router /users/{userID} [post]
func (c *UserRest) GetUserInfo(ctx *gin.Context) {
	// 从上下文获取用户ID
	userID := util.GetUserID(ctx)
	if userID == 0 {
		util.RespFail(ctx, 401, "用户不存在")
		log.Printf("用户不存在，userID: %d", userID)
		return
	}

	// 根据用户ID查询用户信息
	user, err := c.userService.FindUserById(userID)
	if err != nil {
		util.RespFail(ctx, 404, "用户不存在")
		return
	}

	// 构造响应数据
	responseData := vo.UserVO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Role:     user.HaveRole,
	}
	util.RespOk(ctx, responseData)
}

// BuildAuthVO 构建AuthVO
func BuildAuthVO(ctx *gin.Context, user *model.User) *vo.AuthVO {
	clientType := util.GetClientType(ctx)

	// 根据客户端类型设置过期时间
	var expirationTime time.Duration
	if clientType == constant.ClientTypeApp {
		expirationTime = constant.TokenExpireApp
	} else {
		expirationTime = constant.TokenExpireWeb
	}

	// 生成新的access_token
	accessToken, err := util.GenerateToken(user.ID, user.Username, expirationTime)
	if err != nil {
		util.RespFail(ctx, 500, "生成token失败")
		return nil
	}

	// 生成新的refresh_token
	refreshToken, err := util.GenerateToken(user.ID, user.Username, constant.TokenExpireRefresh)
	if err != nil {
		util.RespFail(ctx, 500, "生成refresh_token失败")
		return nil
	}

	return &vo.AuthVO{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		Avatar:       user.Avatar,
		Role:         user.HaveRole,
		Token:        accessToken,
		RefreshToken: refreshToken,
	}
}
