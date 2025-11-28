package database

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/family-bill/bill-server/internal/config"
	"github.com/family-bill/bill-server/models"
)

// DB 全局数据库连接
var DB *gorm.DB

// InitDatabase 初始化数据库连接
func InitDatabase(cfg *config.DatabaseConfig) error {
	var err error

	// 根据配置选择数据库类型
	switch cfg.Type {
	case "sqlite":
		// 确保数据目录存在
		dataDir := filepath.Dir(cfg.SQLite.Path)
		if err := os.MkdirAll(dataDir, 0755); err != nil {
			return fmt.Errorf("failed to create data directory: %w", err)
		}

		// SQLite 连接配置
		log.Info().Str("path", cfg.SQLite.Path).Msg("Using SQLite database")
		// 使用glebarez/sqlite驱动，它支持modernc.org/sqlite，不需要CGO编译
		DB, err = gorm.Open(sqlite.Open(cfg.SQLite.Path), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			return fmt.Errorf("failed to connect to sqlite database: %w", err)
		}

	case "mysql":
		// MySQL 连接配置
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.Host, cfg.MySQL.Port,
			cfg.MySQL.DBName, cfg.MySQL.Charset)
		log.Info().Str("dsn", dsn).Msg("Using MySQL database")
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

	default:
		return fmt.Errorf("unsupported database type: %s", cfg.Type)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// 自动迁移数据库模型
	if err := migrateDatabase(); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Info().Msg("Database connection established and migrated successfully")
	return nil
}

// migrateDatabase 自动迁移数据库模型
func migrateDatabase() error {
	// 自动迁移所有模型
	return DB.AutoMigrate(
		&models.User{},
		&models.Family{},
		&models.FamilyMember{},
		&models.Book{},
		&models.BookAccess{},
		&models.AccountType{},
		&models.AccountGroup{},
		&models.Account{},
		&models.Transaction{},
		&models.TransactionTag{},
		&models.Tag{},
		&models.RecurringTransaction{},
		&models.TransactionTemplate{},
		&models.Currency{},
		&models.ExchangeRate{},
		&models.File{},
		&models.HelpArticle{},
		&models.HelpCategory{},
		&models.Feedback{},
		&models.OperationLog{},
		&models.LoginDevice{},
		&models.TwoFactor{},
		&models.VerificationCode{},
	)
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return DB
}
