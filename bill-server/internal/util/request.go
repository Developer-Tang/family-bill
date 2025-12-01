package util

import (
	"strings"

	"github.com/family-bill/bill-server/internal/constant"
	"github.com/gin-gonic/gin"
)

// 获取指定请求头
func GetRequestHeader(c *gin.Context, headerName string) string {
	return c.Request.Header.Get(headerName)
}

// 获取user-agent
func GetUserAgent(c *gin.Context) string {
	return c.Request.UserAgent()
}

// 根据user-agent判断是web/h5/app
func GetClientType(c *gin.Context) string {
	if clientType, exists := c.Get("clientType"); exists {
		return clientType.(string)
	}

	userAgent := GetUserAgent(c)
	userAgent = strings.ToLower(userAgent)

	// 检查是否为App客户端
	// TODO: 具体逻辑待补充

	// 检查是否为H5客户端
	if StrContainsAny(userAgent, "mobile", "android", "iphone", "ipad") {
		return constant.ClientTypeH5
	}

	// 检查是否为Web客户端
	if StrContainsAny(userAgent, "mozilla", "chrome", "safari", "edge") {
		return constant.ClientTypeWeb
	}

	// 未知客户端
	return constant.ClientTypeUnknown
}

func GetUserID(c *gin.Context) uint {
	token := GetRequestHeader(c, "Authorization")
	if token == "" {
		return 0
	}

	token = strings.TrimPrefix(token, "Bearer ")
	claims, err := ParseToken(token)
	if err != nil {
		return 0
	}

	return claims.UserID
}

// GetIp 获取客户端IP
func GetIp(c *gin.Context) string {
	return c.ClientIP()
}
