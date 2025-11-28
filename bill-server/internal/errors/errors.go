package errors

import (
	"fmt"
)

// ErrorCode 错误码类型
type ErrorCode int

// 错误码定义
const (
	// 系统错误
	ErrInternalServer ErrorCode = 500

	// 请求参数错误
	ErrBadRequest ErrorCode = 400

	// 未授权
	ErrUnauthorized ErrorCode = 401

	// 权限不足
	ErrForbidden ErrorCode = 403

	// 资源不存在
	ErrNotFound ErrorCode = 404

	// 用户相关错误
	ErrUserExists              ErrorCode = 600
	ErrUserNotFound            ErrorCode = 601
	ErrPasswordIncorrect       ErrorCode = 602
	ErrVerificationCodeInvalid ErrorCode = 603

	// 家庭组相关错误
	ErrFamilyExists   ErrorCode = 604
	ErrFamilyNotFound ErrorCode = 605

	// 账本相关错误
	ErrBookExists           ErrorCode = 606
	ErrBookNotFound         ErrorCode = 607
	ErrBookPermissionDenied ErrorCode = 608
)

// AppError 自定义应用错误类型
type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Detail  string    `json:"detail,omitempty"`
}

// Error 实现error接口
func (e *AppError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Detail)
	}
	return e.Message
}

// New 创建新的应用错误
func New(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// NewWithDetail 创建带详细信息的应用错误
func NewWithDetail(code ErrorCode, message string, detail string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
}

// BadRequest 创建请求参数错误
func BadRequest(message string) *AppError {
	return New(ErrBadRequest, message)
}

// BadRequestWithDetail 创建带详细信息的请求参数错误
func BadRequestWithDetail(message string, detail string) *AppError {
	return NewWithDetail(ErrBadRequest, message, detail)
}

// Unauthorized 创建未授权错误
func Unauthorized(message string) *AppError {
	return New(ErrUnauthorized, message)
}

// Forbidden 创建权限不足错误
func Forbidden(message string) *AppError {
	return New(ErrForbidden, message)
}

// NotFound 创建资源不存在错误
func NotFound(message string) *AppError {
	return New(ErrNotFound, message)
}

// InternalServer 创建系统内部错误
func InternalServer(message string) *AppError {
	return New(ErrInternalServer, message)
}

// InternalServerWithDetail 创建带详细信息的系统内部错误
func InternalServerWithDetail(message string, detail string) *AppError {
	return NewWithDetail(ErrInternalServer, message, detail)
}

// UserExists 创建用户已存在错误
func UserExists(message string) *AppError {
	return New(ErrUserExists, message)
}

// UserNotFound 创建用户不存在错误
func UserNotFound(message string) *AppError {
	return New(ErrUserNotFound, message)
}

// PasswordIncorrect 创建密码错误
func PasswordIncorrect(message string) *AppError {
	return New(ErrPasswordIncorrect, message)
}

// VerificationCodeInvalid 创建验证码错误
func VerificationCodeInvalid(message string) *AppError {
	return New(ErrVerificationCodeInvalid, message)
}

// FamilyExists 创建家庭组已存在错误
func FamilyExists(message string) *AppError {
	return New(ErrFamilyExists, message)
}

// FamilyNotFound 创建家庭组不存在错误
func FamilyNotFound(message string) *AppError {
	return New(ErrFamilyNotFound, message)
}

// BookExists 创建账本已存在错误
func BookExists(message string) *AppError {
	return New(ErrBookExists, message)
}

// BookNotFound 创建账本不存在错误
func BookNotFound(message string) *AppError {
	return New(ErrBookNotFound, message)
}

// BookPermissionDenied 创建账本权限不足错误
func BookPermissionDenied(message string) *AppError {
	return New(ErrBookPermissionDenied, message)
}
