package logger

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/family-bill/bill-server/internal/config"
)

// InitLogger 初始化日志系统
func InitLogger(cfg *config.LogConfig) {
	// 创建日志目录
	logDir := filepath.Dir(cfg.File)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal().Err(err).Msg("Failed to create log directory")
	}

	// 打开日志文件
	logFile, err := os.OpenFile(cfg.File, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open log file")
	}

	// 同时输出到控制台和文件
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// 设置日志级别
	level := strings.ToLower(cfg.Level)
	switch level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// 配置日志格式
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(multiWriter)

	log.Info().Str("level", level).Str("file", cfg.File).Msg("Logger initialized")
}
