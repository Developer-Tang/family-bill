package util

import (
	"os"
	"path/filepath"

	"github.com/family-bill/bill-server/internal/config"
	"github.com/sirupsen/logrus"
)

// InitLogger 初始化日志
func InitLogger() {
	Log := logrus.New()

	cfg := config.YamlConfig.Logger

	// 设置日志级别
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	Log.SetLevel(level)

	// 设置日志格式
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 创建日志目录
	if err = os.MkdirAll(cfg.Path, 0755); err != nil {
		Log.Fatalf("Failed to create log directory: %v", err)
	}

	// 设置日志文件
	logFile := filepath.Join(cfg.Path, "app.log")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		Log.Fatalf("Failed to open log file: %v", err)
	}

	// 同时输出到控制台和文件
	Log.SetOutput(os.Stdout)
	Log.AddHook(&fileHook{
		file:     file,
		filename: logFile,
		config:   cfg,
	})
}

// fileHook 日志文件钩子
type fileHook struct {
	file     *os.File
	filename string
	config   config.LoggerConfig
}

// Write 写入日志
func (hook *fileHook) Write(p []byte) (n int, err error) {
	return hook.file.Write(p)
}

// Fire 触发日志写入
func (hook *fileHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Write([]byte(line))
	return err
}

// Levels 支持的日志级别
func (hook *fileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
