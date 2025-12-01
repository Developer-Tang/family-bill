package main

import (
	"fmt"
	"log"

	"github.com/family-bill/bill-server/internal/config"
	"github.com/family-bill/bill-server/internal/db"
	"github.com/family-bill/bill-server/internal/router"
	"github.com/family-bill/bill-server/internal/util"
)

func main() {
	// 加载配置
	config := config.LoadConfig()

	// 初始化日志
	util.InitLogger()

	// 初始化数据库
	db.InitDB()

	// 自动迁移数据库表
	db.AutoMigrate()

	// 初始化路由
	r := router.InitRouter()

	// 启动服务器
	addr := fmt.Sprintf(":%d", config.Server.Port)
	log.Printf("Server starting on http://127.0.0.1:%s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
