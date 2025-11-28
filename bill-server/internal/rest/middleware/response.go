package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Response 统一响应结构体
type Response struct {
	Code    int         `json:"code"`     // 状态码，0表示成功，其他表示错误
	Msg     string      `json:"msg"`      // 消息
	Data    interface{} `json:"data"`     // 数据
	Time    string      `json:"time"`     // 时间
	TradeID string      `json:"trade_id"` // 交易ID
}

// TradeIDKey 交易ID上下文键
const TradeIDKey = "trade_id"

// GenerateTradeID 生成交易ID
func GenerateTradeID() string {
	return uuid.New().String()
}

// ResponseMiddleware 响应中间件，生成交易ID
func ResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生成交易ID
		tradeID := GenerateTradeID()
		// 将交易ID添加到上下文
		c.Set(TradeIDKey, tradeID)
		// 继续处理请求
		c.Next()
	}
}

// SuccessResponse 成功响应
func SuccessResponse(c *gin.Context, data interface{}) {
	// 从上下文获取交易ID
	tradeID, _ := c.Get(TradeIDKey)
	// 构造响应
	response := Response{
		Code:    0,
		Msg:     "success",
		Data:    data,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		TradeID: tradeID.(string),
	}
	// 返回响应
	c.JSON(http.StatusOK, response)
}

// ErrorResponse 错误响应
func ErrorResponse(c *gin.Context, code int, msg string) {
	// 从上下文获取交易ID
	tradeID, _ := c.Get(TradeIDKey)
	// 构造响应
	response := Response{
		Code:    code,
		Msg:     msg,
		Data:    nil,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		TradeID: tradeID.(string),
	}
	// 返回响应
	c.JSON(http.StatusBadRequest, response)
}

// ServerErrorResponse 服务器错误响应
func ServerErrorResponse(c *gin.Context, msg string) {
	// 从上下文获取交易ID
	tradeID, _ := c.Get(TradeIDKey)
	// 构造响应
	response := Response{
		Code:    500,
		Msg:     msg,
		Data:    nil,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		TradeID: tradeID.(string),
	}
	// 返回响应
	c.JSON(http.StatusInternalServerError, response)
}
