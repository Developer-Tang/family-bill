package db

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/family-bill/bill-server/pkg/utils"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB(cfg *utils.Config) {
	var err error

	// 配置GORM日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// 根据数据库驱动选择连接方式
	// 根据数据库驱动选择连接方式
	switch cfg.DB.Driver {
	case "sqlite":
		// 确保数据目录存在
		dataDir := filepath.Dir(cfg.DB.DSN)
		if err := os.MkdirAll(dataDir, 0755); err != nil {
			log.Fatalf("Failed to create data directory: %v", err)
		}

		// SQLite 连接配置
		log.Printf("Using SQLite database with path: %s", cfg.DB.DSN)
		DB, err = gorm.Open(sqlite.Open(cfg.DB.DSN), &gorm.Config{
			Logger: newLogger,
		})

	case "mysql":
		// MySQL 连接配置
		log.Printf("Using MySQL database with DSN: %s", cfg.DB.DSN)
		DB, err = gorm.Open(mysql.Open(cfg.DB.DSN), &gorm.Config{
			Logger: newLogger,
		})

	default:
		log.Fatalf("unsupported database type: %s", cfg.DB.Driver)
	}

	if err != nil {
		log.Fatalf("Failed to connect to %s database: %v", cfg.DB.Driver, err)
	}

	// 获取底层sql.DB对象，配置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection pool: %v", err)
	}

	// 配置连接池
	sqlDB.SetMaxIdleConns(cfg.DB.MaxIdle)
	sqlDB.SetMaxOpenConns(cfg.DB.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.DB.MaxLife) * time.Second)

	log.Println("Database connected successfully")
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return DB
}
