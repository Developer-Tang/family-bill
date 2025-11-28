package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/family-bill/bill-server/internal/errors"
)

// Response 统一API响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SuccessResponse 成功响应
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "操作成功",
		Data:    data,
	})
}

// SuccessResponseWithMessage 带自定义消息的成功响应
func SuccessResponseWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse 错误响应
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// ErrorResponseWithData 带数据的错误响应
func ErrorResponseWithData(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// BadRequestResponse 请求参数错误响应
func BadRequestResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    400,
		Message: message,
		Data:    nil,
	})
}

// UnauthorizedResponse 未授权响应
func UnauthorizedResponse(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, Response{
		Code:    401,
		Message: message,
		Data:    nil,
	})
}

// ForbiddenResponse 权限不足响应
func ForbiddenResponse(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, Response{
		Code:    403,
		Message: message,
		Data:    nil,
	})
}

// NotFoundResponse 资源不存在响应
func NotFoundResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Code:    404,
		Message: message,
		Data:    nil,
	})
}

// InternalServerErrorResponse 服务器内部错误响应
func InternalServerErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Response{
		Code:    500,
		Message: message,
		Data:    nil,
	})
}

// ErrorResponseFromError 从错误对象生成错误响应
func ErrorResponseFromError(c *gin.Context, err error) {
	// 检查是否为自定义应用错误
	if appErr, ok := err.(*errors.AppError); ok {
		c.JSON(http.StatusOK, Response{
			Code:    int(appErr.Code),
			Message: appErr.Message,
			Data: gin.H{
				"error_details": appErr.Detail,
			},
		})
		return
	}

	// 其他类型错误，返回系统内部错误
	c.JSON(http.StatusInternalServerError, Response{
		Code:    500,
		Message: "服务器内部错误",
		Data: gin.H{
			"error_details": err.Error(),
		},
	})
}
