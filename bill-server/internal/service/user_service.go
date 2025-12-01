package service

import (
	"fmt"
	"log"

	"github.com/family-bill/bill-server/internal/db"
	"github.com/family-bill/bill-server/internal/model"
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
	result := db.GetDB().Exec(`
		INSERT INTO tb_users (username, email, password, have_role, status, created_at, updated_at)
		SELECT ?, ?, ?, 'admin', 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
		WHERE NOT EXISTS (SELECT 1 FROM tb_users WHERE status = 1)
	`, username, email, hashedPassword)

	if err := result.Error; err != nil || result.RowsAffected == 0 {
		log.Printf("Register error: %v, RowsAffected: %d", err, result.RowsAffected)
		return nil, fmt.Errorf("注册失败，仅第一位用户支持注册")
	}

	// 返回注册成功的用户信息
	var user model.User
	db.GetDB().Where("username = ? AND status = 1", username).First(&user)
	return &user, nil
}

// Login 用户登录
func (s *UserService) Login(username, password string) (*model.User, error) {
	// 根据用户名查找用户
	var user model.User
	db := db.GetDB().Where("username = ? AND status = 1", username)
	if err := db.First(&user).Error; err != nil {
		return nil, err
	}

	// 验证密码（暂时使用简单比较，实际项目中应使用bcrypt等安全加密算法）
	if user.Password != password {
		return nil, fmt.Errorf("密码错误")
	}

	return &user, nil
}

// FindUserById 根据用户ID查找用户
func (s *UserService) FindUserById(userId uint) (*model.User, error) {
	// 根据用户名查找用户
	var user model.User
	db := db.GetDB().Where("id = ? AND status = 1", userId)
	if err := db.First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
