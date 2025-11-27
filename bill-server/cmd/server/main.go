package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/family-bill/bill-server/internal/config"
	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/internal/logger"
	"github.com/family-bill/bill-server/internal/router"
)

func main() {
	// 加载配置文件
	cfg, err := config.LoadConfig("")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	// 初始化日志系统
	logger.InitLogger(&cfg.Log)
	// 初始化数据库连接
	if err := database.InitDatabase(&cfg.Database); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize database")
	}

	// 设置 Gin 模式
	gin.SetMode(cfg.Server.Mode)

	// 设置路由
	r := router.SetupRouter()

	// 启动服务器
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Info().Str("addr", addr).Msg("Server starting")

	if err := r.Run(addr); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
