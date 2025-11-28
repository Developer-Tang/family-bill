package middleware

import (
	"bytes"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 静态文件后缀列表
var staticFileExtensions = map[string]bool{
	".css":   true,
	".js":    true,
	".png":   true,
	".jpg":   true,
	".jpeg":  true,
	".gif":   true,
	".svg":   true,
	".ico":   true,
	".woff":  true,
	".woff2": true,
	".ttf":   true,
	".otf":   true,
	".map":   true,
}

// isStaticFile 判断是否为静态文件
func isStaticFile(path string) bool {
	// 获取文件后缀
	ext := strings.ToLower(filepath.Ext(path))
	// 检查是否为静态文件后缀
	if _, ok := staticFileExtensions[ext]; ok {
		return true
	}
	// 特殊处理Swagger相关文件
	if strings.Contains(path, "/swagger/") {
		return true
	}
	return false
}

// RequestLogger 请求日志中间件
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文获取交易ID
		tradeID, _ := c.Get(TradeIDKey)

		// 开始时间
		startTime := time.Now()

		// 记录请求体
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 重置请求体，以便后续处理
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 记录请求参数
		params := c.Request.URL.Query()

		// 创建响应体记录器
		responseBody := &bytes.Buffer{}
		c.Writer = &responseWriter{body: responseBody, ResponseWriter: c.Writer}

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		// 计算用时
		latency := endTime.Sub(startTime)

		// 判断是否为静态文件
		path := c.Request.URL.Path
		if isStaticFile(path) {
			// 静态文件，只记录基本信息，不记录响应体
			log.Printf(
				"[REQUEST] Path: %s, Method: %s, Params: %v, Body: %s, Status: %d, Latency: %v, TradeID: %s",
				path,
				c.Request.Method,
				params,
				string(requestBody),
				c.Writer.Status(),
				latency,
				tradeID.(string),
			)
			return
		}

		// 非静态文件，记录完整信息
		// 记录日志
		log.Printf(
			"[REQUEST] Path: %s, Method: %s, Params: %v, Body: %s, Status: %d, Response: %s, Latency: %v, TradeID: %s",
			path,
			c.Request.Method,
			params,
			string(requestBody),
			c.Writer.Status(),
			responseBody.String(),
			latency,
			tradeID.(string),
		)
	}
}

// responseWriter 响应体记录器
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write 写入响应体
func (w *responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func parseValidationError(err error) string {
	// 从原始错误信息中提取Error:后面的内容
	errStr := err.Error()
	if idx := strings.Index(errStr, "Error:"); idx != -1 {
		return strings.TrimSpace(errStr[idx+6:])
	}

	return errStr
}
