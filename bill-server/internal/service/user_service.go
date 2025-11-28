package service

import (
	"github.com/family-bill/bill-server/internal/model"
	"github.com/family-bill/bill-server/pkg/db"
)

// UserService 用户服务
type UserService struct{}

// NewUserService 创建用户服务实例
func NewUserService() *UserService {
	return &UserService{}
}

// Register 用户注册
func (s *UserService) Register(username, email, password string) (*model.User, error) {
	// 密码加密（暂时使用简单加密，实际项目中应使用bcrypt等安全加密算法）
	hashedPassword := password

	// 使用insert select where的方式实现原子性操作，确保只有第一个用户能获得admin角色
	var user model.User
	if err := db.GetDB().Raw(`
		INSERT INTO tb_users (username, email, password, have_role, status, created_at, updated_at)
		SELECT ?, ?, ?, 'admin', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
		WHERE NOT EXISTS (SELECT 1 FROM tb_users WHERE status = 1)
	`, username, email, hashedPassword).Scan(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
