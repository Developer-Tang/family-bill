package v1

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/family-bill/bill-server/internal/config"
	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
)

// UpdateProfileRequest 更新个人信息请求结构体
type UpdateProfileRequest struct {
	Username string `json:"username" binding:"omitempty,min=2,max=20"`
	Phone    string `json:"phone" binding:"omitempty,e164"`
}

// GetProfile 获取用户个人信息
func GetProfile(c *gin.Context) {
	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	// 这里简化处理，假设用户ID为1
	userID := uint(1)

	// 查询用户信息
	var user models.User
	result := database.DB.First(&user, userID)
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

	// 格式化手机号，隐藏中间四位
	formattedPhone := user.Phone
	if len(formattedPhone) == 11 {
		formattedPhone = fmt.Sprintf("%s****%s", formattedPhone[:3], formattedPhone[7:])
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"user_id":         user.UserID,
			"username":        user.Username,
			"email":           user.Email,
			"phone":           formattedPhone,
			"avatar":          user.Avatar,
			"created_at":      user.CreatedAt,
			"last_login":      user.LastLogin,
			"status":          user.Status,
			"email_verified":  user.EmailVerified,
			"phone_verified":  user.PhoneVerified,
			"default_book_id": user.DefaultBookID,
		},
	})
}

// UpdateProfile 更新用户个人信息
func UpdateProfile(c *gin.Context) {
	var req UpdateProfileRequest
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
	// 这里简化处理，假设用户ID为1
	userID := uint(1)

	// 查询用户信息
	var user models.User
	result := database.DB.First(&user, userID)
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

	// 更新用户信息
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	user.UpdatedAt = time.Now()

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "更新失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	// 格式化手机号，隐藏中间四位
	formattedPhone := user.Phone
	if len(formattedPhone) == 11 {
		formattedPhone = fmt.Sprintf("%s****%s", formattedPhone[:3], formattedPhone[7:])
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
		"data": gin.H{
			"user_id":    user.UserID,
			"username":   user.Username,
			"email":      user.Email,
			"phone":      formattedPhone,
			"avatar":     user.Avatar,
			"updated_at": user.UpdatedAt,
		},
	})
}

// UpdateAvatar 更新用户头像
func UpdateAvatar(c *gin.Context) {
	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	// 这里简化处理，假设用户ID为1
	userID := uint(1)

	// 查询用户信息
	var user models.User
	result := database.DB.First(&user, userID)
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

	// 处理文件上传
	file, header, err := c.Request.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "上传失败",
			"data": gin.H{
				"error_details": "请选择要上传的头像文件",
			},
		})
		return
	}
	defer file.Close()

	// 验证文件类型
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}

	contentType := header.Header.Get("Content-Type")
	if !allowedTypes[contentType] {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    4004,
			"message": "文件类型不支持",
			"data": gin.H{
				"error_details": "只支持JPG、PNG、GIF格式的图片",
			},
		})
		return
	}

	// 验证文件大小（最大2MB）
	maxSize := int64(2 * 1024 * 1024)
	if header.Size > maxSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    4005,
			"message": "文件大小超限",
			"data": gin.H{
				"error_details": "头像文件大小不能超过2MB",
			},
		})
		return
	}

	// 保存文件 - 这里简化处理，实际项目中应该保存到文件系统或云存储
	// 生成文件名
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("avatar_%d_%d%s", userID, time.Now().Unix(), ext)

	// 实际项目中应该使用io.Copy将文件内容保存到指定路径
	// 这里简化处理，直接设置URL
	avatarURL := fmt.Sprintf("https://example.com/avatars/%s", filename)

	// 更新用户头像
	user.Avatar = avatarURL
	user.UpdatedAt = time.Now()

	if err := database.DB.Save(&user).Error; err != nil {
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
		"message": "头像更新成功",
		"data": gin.H{
			"user_id":    user.UserID,
			"avatar":     user.Avatar,
			"updated_at": user.UpdatedAt,
		},
	})
}

// parseUserIDFromToken 从JWT令牌中解析用户ID
func parseUserIDFromToken(tokenString string) (uint, error) {
	// 解析JWT令牌
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GlobalConfig.JWT.Secret), nil // 从配置中获取密钥
	})

	if err != nil {
		return 0, err
	}

	// 从令牌中获取用户ID
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	userID := uint(claims["user_id"].(float64))
	return userID, nil
}
