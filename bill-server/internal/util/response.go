package util

import (
	"net/http"
	"strings"
	"time"

	"github.com/family-bill/bill-server/internal/constant"
	"github.com/family-bill/bill-server/internal/model/vo"
	"github.com/gin-gonic/gin"
)

// RespOk 成功响应
func RespOk(c *gin.Context, data interface{}) {
	// 从上下文获取交易ID
	tradeID, _ := c.Get(constant.TradeIDKey)
	// 构造响应
	response := vo.R{
		Code:    0,
		Msg:     "success",
		Data:    data,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		TradeID: tradeID.(string),
	}
	// 返回响应
	c.JSON(http.StatusOK, response)
}

// RespFail 错误响应
func RespFail(c *gin.Context, code int, msg string) {
	// 从上下文获取交易ID
	tradeID, _ := c.Get(constant.TradeIDKey)
	// 构造响应
	response := vo.R{
		Code:    code,
		Msg:     msg,
		Data:    nil,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		TradeID: tradeID.(string),
	}
	// 返回响应
	c.JSON(http.StatusBadRequest, response)
}

// 解析验证错误，只返回具体的错误信息
func ParseValidationError(err error) string {
	// 从原始错误信息中提取Error:后面的内容
	errStr := err.Error()
	if idx := strings.Index(errStr, "Error:"); idx != -1 {
		return strings.TrimSpace(errStr[idx+6:])
	}

	return errStr
}
