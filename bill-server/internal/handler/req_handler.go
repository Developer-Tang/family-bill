package handler

import (
	"bytes"
	"io"
	"log"
	"strings"
	"time"

	"github.com/family-bill/bill-server/internal/constant"
	"github.com/family-bill/bill-server/internal/util"
	"github.com/gin-gonic/gin"
)

// InitFilter 请求中间件
func InitReqAroundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path

		// 解析请求客户端类型
		clientType := util.GetClientType(c)
		c.Set(constant.ClientTypeKey, clientType)

		// 解析请求用户ID
		userID := util.GetUserID(c)
		c.Set(constant.UserIDKey, userID)

		// 生成交易ID
		tradeID := util.GenerateUUID()
		// 将交易ID添加到上下文
		c.Set(constant.TradeIDKey, tradeID)

		// 记录请求参数
		params := c.Request.URL.Query()

		// 记录请求体
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 重置请求体，以便后续处理
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 创建响应体记录器
		responseBody := &bytes.Buffer{}
		c.Writer = &responseWriter{body: responseBody, ResponseWriter: c.Writer}

		// 开始时间
		startTime := time.Now()

		// 继续处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		// 计算用时
		latency := endTime.Sub(startTime)

		// 全部接口POST编写
		if !strings.EqualFold(method, "POST") {
			// 静态文件，只记录基本信息，不记录响应体
			log.Printf(
				"[REQUEST] Path: %s, Method: %s, Params: %v, Body: %s, Status: %d, Latency: %v, TradeID: %s",
				c.Request.URL.Path,
				c.Request.Method,
				params,
				string(requestBody),
				c.Writer.Status(),
				latency,
				tradeID,
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
			tradeID,
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
