package middleware

import (
	"github.com/gin-gonic/gin"
)

// RequestMiddleware
func RequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 继续处理请求
		c.Next()
	}
}

// 获取指定请求头
func GetRequestHeader(c *gin.Context, headerName string) string {
	return c.Request.Header.Get(headerName)
}

// 获取user-agent
func GetUserAgent(c *gin.Context) string {
	return c.Request.UserAgent()
}

// 根据user-agent判断是web/h5/app
