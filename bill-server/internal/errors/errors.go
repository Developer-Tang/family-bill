package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIError API错误结构体
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

// NewAPIError 创建新的API错误
func NewAPIError(code int, message string, details any) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// Error 实现error接口
func (e *APIError) Error() string {
	return e.Message
}

// BadRequestError 400错误
func BadRequestError(message string, details any) *APIError {
	return NewAPIError(http.StatusBadRequest, message, details)
}

// UnauthorizedError 401错误
func UnauthorizedError(message string, details any) *APIError {
	return NewAPIError(http.StatusUnauthorized, message, details)
}

// ForbiddenError 403错误
func ForbiddenError(message string, details any) *APIError {
	return NewAPIError(http.StatusForbidden, message, details)
}

// NotFoundError 404错误
func NotFoundError(message string, details any) *APIError {
	return NewAPIError(http.StatusNotFound, message, details)
}

// InternalServerError 500错误
func InternalServerError(message string, details any) *APIError {
	return NewAPIError(http.StatusInternalServerError, message, details)
}

// HandleError 统一处理错误
func HandleError(c *gin.Context, err error) {
	// 检查是否为APIError类型
	if apiErr, ok := err.(*APIError); ok {
		c.JSON(apiErr.Code, apiErr)
		return
	}

	// 其他错误类型，返回500
	c.JSON(http.StatusInternalServerError, NewAPIError(
		http.StatusInternalServerError,
		"Internal Server Error",
		err.Error(),
	))
}
