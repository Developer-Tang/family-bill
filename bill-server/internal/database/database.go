package database

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"github.com/family-bill/bill-server/internal/config"
)

// DB 全局数据库连接
var DB *gorm.DB

// InitDatabase 初始化数据库连接
func InitDatabase(cfg *config.DatabaseConfig) error {
	// 根据配置选择数据库类型
	switch cfg.Type {
	case "sqlite":
		// SQLite 连接配置
		log.Info().Str("path", cfg.SQLite.Path).Msg("Using SQLite database")

	case "mysql":
		// MySQL 连接配置
		log.Info().Str("host", cfg.MySQL.Host).Str("dbname", cfg.MySQL.DBName).Msg("Using MySQL database")

	default:
		return fmt.Errorf("unsupported database type: %s", cfg.Type)
	}

	// 连接数据库 - 这里我们暂时跳过实际连接，只做日志记录
	// 实际项目中应该使用真实的数据库连接
	log.Info().Msg("Database connection established (mock)")
	return nil
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return DB
}
