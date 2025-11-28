package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// GenerateVerificationCode 生成指定长度的随机数字验证码
func GenerateVerificationCode(length int) (string, error) {
	if length <= 0 {
		length = 6 // 默认6位
	}

	code := ""
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		code += fmt.Sprintf("%d", num)
	}

	return code, nil
}

// GenerateToken 生成随机令牌
func GenerateToken(length int) (string, error) {
	if length <= 0 {
		length = 32 // 默认32位
	}

	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := ""
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		token += string(chars[index.Int64()])
	}

	return token, nil
}

// IsExpired 检查时间是否已过期
func IsExpired(expiresAt time.Time) bool {
	return time.Now().After(expiresAt)
}
