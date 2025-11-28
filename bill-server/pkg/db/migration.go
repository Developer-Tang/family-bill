package db

import (
	"log"

	"github.com/family-bill/bill-server/internal/model"
)

// AutoMigrate 自动迁移数据库表
func AutoMigrate() {
	log.Println("Starting database migration...")

	// 自动迁移所有模型
	if err := DB.AutoMigrate(
		&model.User{},
		&model.Login{},
		&model.Operation{},
		&model.VerifyCode{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully")
}
