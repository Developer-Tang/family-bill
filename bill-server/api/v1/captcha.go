package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetImageCaptcha 获取图形验证码
func GetImageCaptcha(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": gin.H{
			"captcha_token": "captcha_123",
			"image_data":    "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA...",
			"expires_in":    180,
		},
	})
}

// VerifyImageCaptcha 验证图形验证码
func VerifyImageCaptcha(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "验证码验证成功",
		"data": gin.H{
			"is_valid": true,
		},
	})
}

// SendSmsCaptcha 发送短信验证码
func SendSmsCaptcha(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "短信验证码发送成功",
		"data": gin.H{
			"phone":     "13800138000",
			"expires_in": 300,
		},
	})
}

// SendEmailCaptcha 发送邮箱验证码
func SendEmailCaptcha(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "邮箱验证码发送成功",
		"data": gin.H{
			"email":     "user@example.com",
			"expires_in": 300,
		},
	})
}

// SendVoiceCaptcha 发送语音验证码
func SendVoiceCaptcha(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "语音验证码发送成功",
		"data": gin.H{
			"phone":     "13800138000",
			"expires_in": 300,
		},
	})
}
