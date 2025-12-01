package util

import (
	"errors"
	"time"

	"github.com/family-bill/bill-server/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims JWT声明结构体
type JWTClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT令牌，根据客户端类型设置不同有效期
func GenerateToken(userID uint, username string, expirationTime time.Duration) (string, error) {
	claims := JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.YamlConfig.JWT.Secret))
}

// ParseToken 解析JWT令牌
func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.YamlConfig.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)

	if ok && token.Valid && claims.ExpiresAt.After(time.Now()) {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
