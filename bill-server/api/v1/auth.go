package v1

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/family-bill/bill-server/internal/config"
	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
	"github.com/family-bill/bill-server/utils"
)

// LoginRequest 登录请求结构体
type LoginRequest struct {
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=6,max=20"`
	CaptchaToken string `json:"captcha_token" binding:"required"`
	CaptchaCode  string `json:"captcha_code" binding:"required,len=4"`
	RememberMe   bool   `json:"remember_me"`
}

// RegisterRequest 注册请求结构体
type RegisterRequest struct {
	Username         string `json:"username" binding:"required,min=2,max=20"`
	Email            string `json:"email" binding:"required,email"`
	Password         string `json:"password" binding:"required,min=6,max=20"`
	ConfirmPassword  string `json:"confirm_password" binding:"required,eqfield=Password"`
	Phone            string `json:"phone" binding:"omitempty,e164"`
	VerificationCode string `json:"verification_code" binding:"required,len=6"`
	CaptchaToken     string `json:"captcha_token" binding:"required"`
	CaptchaCode      string `json:"captcha_code" binding:"required,len=4"`
}

// SendCodeRequest 发送验证码请求结构体
type SendCodeRequest struct {
	Contact string `json:"contact" binding:"required"`
	Type    string `json:"type" binding:"required,oneof=email sms"`
	Purpose string `json:"purpose" binding:"required,oneof=register login reset_password verify_email"`
}

// VerifyEmailRequest 邮箱验证请求结构体
type VerifyEmailRequest struct {
	Email            string `json:"email" binding:"required,email"`
	VerificationCode string `json:"verification_code" binding:"required,len=6"`
}

// ResetPasswordRequest 重置密码请求结构体
type ResetPasswordRequest struct {
	Email            string `json:"email" binding:"required,email"`
	VerificationCode string `json:"verification_code" binding:"required,len=6"`
	NewPassword      string `json:"new_password" binding:"required,min=6,max=20"`
	ConfirmPassword  string `json:"confirm_password" binding:"required,eqfield=NewPassword"`
}

// RefreshTokenRequest 刷新令牌请求结构体
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// Login 用户登录
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 验证验证码 - 这里简化处理，实际项目中需要验证

	// 查询用户
	var user models.User
	result := database.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		utils.ErrorResponseWithData(c, 4000, "登录失败", gin.H{
			"error_details": "邮箱或密码错误",
		})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.ErrorResponseWithData(c, 4000, "登录失败", gin.H{
			"error_details": "邮箱或密码错误",
		})
		return
	}

	// 更新最后登录时间
	user.LastLogin = time.Now()
	database.DB.Save(&user)

	// 生成JWT令牌
	accessToken, refreshToken, err := generateTokens(user.UserID)
	if err != nil {
		utils.ErrorResponseWithData(c, 500, "生成令牌失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	utils.SuccessResponseWithMessage(c, "登录成功", gin.H{
		"user": gin.H{
			"user_id":    user.UserID,
			"username":   user.Username,
			"email":      user.Email,
			"avatar":     user.Avatar,
			"created_at": user.CreatedAt,
		},
		"tokens": gin.H{
			"access_token":             accessToken,
			"refresh_token":            refreshToken,
			"access_token_expires_in":  3600,
			"refresh_token_expires_in": 86400,
		},
		"default_book_id": user.DefaultBookID,
	})
}

// Register 用户注册
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 验证验证码
	isValid, message := verifyVerificationCode(req.Email, req.VerificationCode, req.CaptchaToken, "register")
	if !isValid {
		utils.ErrorResponseWithData(c, 4003, "注册失败", gin.H{
			"error_details": message,
		})
		return
	}

	// 检查邮箱是否已注册
	var existingUser models.User
	result := database.DB.Where("email = ?", req.Email).First(&existingUser)
	if result.Error == nil {
		utils.ErrorResponseWithData(c, 4001, "注册失败", gin.H{
			"error_details": "邮箱已被注册",
		})
		return
	}

	// 哈希密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.ErrorResponseWithData(c, 500, "密码哈希失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 创建用户
	user := models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  string(hashedPassword),
		Phone:     req.Phone,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		LastLogin: time.Now(),
	}

	// 开始事务
	tx := database.DB.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponseWithData(c, 500, "创建用户失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 创建默认账本
	defaultBook := models.Book{
		Name:        fmt.Sprintf("%s的账本", req.Username),
		CreatorID:   user.UserID,
		Currency:    "CNY",
		Description: "默认账本",
		Status:      "active",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := tx.Create(&defaultBook).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponseWithData(c, 500, "创建默认账本失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 更新用户默认账本ID
	user.DefaultBookID = defaultBook.BookID
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponseWithData(c, 500, "更新用户默认账本失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 创建账本访问权限
	bookAccess := models.BookAccess{
		BookID:    defaultBook.BookID,
		UserID:    user.UserID,
		Role:      "owner",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := tx.Create(&bookAccess).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponseWithData(c, 500, "创建账本访问权限失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 提交事务
	tx.Commit()

	utils.SuccessResponseWithMessage(c, "注册成功", gin.H{
		"user_id":              user.UserID,
		"username":             user.Username,
		"email":                user.Email,
		"created_at":           user.CreatedAt,
		"default_book_created": true,
		"verification_status":  "pending",
	})
}

// Logout 用户登出
func Logout(c *gin.Context) {
	// 实际项目中需要将令牌加入黑名单
	utils.SuccessResponseWithMessage(c, "登出成功", gin.H{
		"logout_time": time.Now().Format(time.RFC3339),
	})
}

// RefreshToken 刷新令牌
func RefreshToken(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 解析刷新令牌
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GlobalConfig.JWT.Secret), nil // 从配置中获取密钥
	})

	if err != nil || !token.Valid {
		utils.ErrorResponseWithData(c, 401, "刷新令牌无效或已过期", gin.H{
			"error_details": "请重新登录",
		})
		return
	}

	// 从令牌中获取用户ID
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		utils.ErrorResponseWithData(c, 401, "刷新令牌无效", gin.H{
			"error_details": "请重新登录",
		})
		return
	}

	userID := uint(claims["user_id"].(float64))

	// 生成新的访问令牌和刷新令牌
	accessToken, refreshToken, err := generateTokens(userID)
	if err != nil {
		utils.ErrorResponseWithData(c, 500, "生成令牌失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	utils.SuccessResponseWithMessage(c, "令牌刷新成功", gin.H{
		"access_token":             accessToken,
		"access_token_expires_in":  3600,
		"refresh_token":            refreshToken,
		"refresh_token_expires_in": 86400,
	})
}

// VerifyEmail 邮箱验证
func VerifyEmail(c *gin.Context) {
	var req VerifyEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 这里需要从请求中获取token，暂时简化处理，假设token在请求头中
	token := c.GetHeader("X-Verification-Token")
	if token == "" {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": "缺少验证码令牌",
		})
		return
	}

	// 验证验证码
	isValid, message := verifyVerificationCode(req.Email, req.VerificationCode, token, "verify_email")
	if !isValid {
		utils.ErrorResponseWithData(c, 4003, "邮箱验证失败", gin.H{
			"error_details": message,
		})
		return
	}

	// 更新用户邮箱验证状态
	result := database.DB.Model(&models.User{}).Where("email = ?", req.Email).Update("email_verified", true)
	if result.Error != nil {
		utils.ErrorResponseWithData(c, 500, "邮箱验证失败", gin.H{
			"error_details": result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		utils.ErrorResponseWithData(c, 4001, "邮箱验证失败", gin.H{
			"error_details": "用户不存在",
		})
		return
	}

	utils.SuccessResponseWithMessage(c, "邮箱验证成功", gin.H{
		"email":       req.Email,
		"verified_at": time.Now().Format(time.RFC3339),
	})
}

// ResetPassword 重置密码
func ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 这里需要从请求中获取token，暂时简化处理，假设token在请求头中
	token := c.GetHeader("X-Verification-Token")
	if token == "" {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": "缺少验证码令牌",
		})
		return
	}

	// 验证验证码
	isValid, message := verifyVerificationCode(req.Email, req.VerificationCode, token, "reset_password")
	if !isValid {
		utils.ErrorResponseWithData(c, 4003, "密码重置失败", gin.H{
			"error_details": message,
		})
		return
	}

	// 哈希新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.ErrorResponseWithData(c, 500, "密码哈希失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 更新用户密码
	result := database.DB.Model(&models.User{}).Where("email = ?", req.Email).Update("password", string(hashedPassword))
	if result.Error != nil {
		utils.ErrorResponseWithData(c, 500, "密码重置失败", gin.H{
			"error_details": result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		utils.ErrorResponseWithData(c, 4001, "密码重置失败", gin.H{
			"error_details": "用户不存在",
		})
		return
	}

	utils.SuccessResponseWithMessage(c, "密码重置成功", gin.H{
		"email":    req.Email,
		"reset_at": time.Now().Format(time.RFC3339),
	})
}

// SendCode 发送验证码
func SendCode(c *gin.Context) {
	var req SendCodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 生成6位随机验证码
	verificationCode, err := utils.GenerateVerificationCode(6)
	if err != nil {
		utils.ErrorResponseWithData(c, 500, "生成验证码失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 生成令牌
	token, err := utils.GenerateToken(32)
	if err != nil {
		utils.ErrorResponseWithData(c, 500, "生成令牌失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 计算过期时间（5分钟后）
	expiresAt := time.Now().Add(5 * time.Minute)

	// 保存验证码到数据库
	code := models.VerificationCode{
		Code:        verificationCode,
		Contact:     req.Contact,
		Type:        req.Type,
		Purpose:     req.Purpose,
		Token:       token,
		ExpiresAt:   expiresAt,
		Attempts:    0,
		MaxAttempts: 5,
		IsUsed:      false,
	}

	if err := database.DB.Create(&code).Error; err != nil {
		utils.ErrorResponseWithData(c, 500, "保存验证码失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 发送验证码 - 这里简化处理，实际项目中需要调用邮件或短信服务
	fmt.Printf("发送验证码 %s 到 %s，用途：%s\n", verificationCode, req.Contact, req.Purpose)

	utils.SuccessResponseWithMessage(c, "验证码发送成功", gin.H{
		"contact":    req.Contact,
		"type":       req.Type,
		"expires_in": 300,
		"token":      token,
	})
}

// generateTokens 生成JWT令牌
func generateTokens(userID uint) (string, string, error) {
	// 从配置中获取JWT密钥
	jwtSecret := []byte(config.GlobalConfig.JWT.Secret)

	// 生成访问令牌
	accessTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour).Unix(),
		"iat":     time.Now().Unix(),
		"type":    "access",
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// 生成刷新令牌
	refreshTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
		"type":    "refresh",
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

// verifyVerificationCode 验证验证码
func verifyVerificationCode(contact, code, token, purpose string) (bool, string) {
	var verificationCode models.VerificationCode

	// 根据token和contact查找验证码
	result := database.DB.Where("token = ? AND contact = ? AND purpose = ?", token, contact, purpose).First(&verificationCode)
	if result.Error != nil {
		return false, "验证码不存在或已过期"
	}

	// 检查验证码是否已使用
	if verificationCode.IsUsed {
		return false, "验证码已使用"
	}

	// 检查验证码是否已过期
	if utils.IsExpired(verificationCode.ExpiresAt) {
		return false, "验证码已过期"
	}

	// 检查验证码尝试次数是否超过限制
	if verificationCode.Attempts >= verificationCode.MaxAttempts {
		return false, "验证码尝试次数已超过限制"
	}

	// 检查验证码是否正确
	if verificationCode.Code != code {
		// 增加尝试次数
		database.DB.Model(&verificationCode).Update("attempts", verificationCode.Attempts+1)
		return false, "验证码错误"
	}

	// 验证成功，标记验证码为已使用
	database.DB.Model(&verificationCode).Update("is_used", true)

	return true, "验证码验证成功"
}
