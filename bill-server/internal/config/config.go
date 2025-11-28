package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config 应用配置结构体
type Config struct {
	Version    string           `yaml:"version"`
	Server     ServerConfig     `yaml:"server"`
	Database   DatabaseConfig   `yaml:"database"`
	JWT        JWTConfig        `yaml:"jwt"`
	Log        LogConfig        `yaml:"log"`
	DateFormat DateFormatConfig `yaml:"date_format"`
}

// GlobalConfig 全局配置变量
var GlobalConfig *Config

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
	Mode string `yaml:"mode"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Type   string       `yaml:"type"`
	SQLite SQLiteConfig `yaml:"sqlite"`
	MySQL  MySQLConfig  `yaml:"mysql"`
}

// SQLiteConfig SQLite配置
type SQLiteConfig struct {
	Path string `yaml:"path"`
}

// MySQLConfig MySQL配置
type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret string `yaml:"secret"`
	Expire int    `yaml:"expire"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `yaml:"level"`
	File       string `yaml:"file"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
}

// DateFormatConfig 日期格式配置
type DateFormatConfig struct {
	DateOnly     string `yaml:"date_only"`      // yyyy-MM-dd
	DateTime     string `yaml:"date_time"`      // yyyy-MM-dd HH:mm
	DateTimeFull string `yaml:"date_time_full"` // yyyy-MM-dd HH:mm:ss
	TimeOnly     string `yaml:"time_only"`      // HH:mm:ss
	TimeShort    string `yaml:"time_short"`     // HH:mm
}

// LoadConfig 加载配置文件
func LoadConfig(configPath string) (*Config, error) {
	// 如果没有提供配置路径，使用默认路径
	if configPath == "" {
		configPath = filepath.Join("config.yaml")
	}

	// 读取配置文件
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	// 解析配置文件
	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	// 从环境变量覆盖配置
	overrideFromEnv(&config)

	// 设置全局配置
	GlobalConfig = &config

	return &config, nil
}

// overrideFromEnv 从环境变量覆盖配置
func overrideFromEnv(config *Config) {
	// 服务器配置
	if port := os.Getenv("SERVER_PORT"); port != "" {
		config.Server.Port = port
	}
	if host := os.Getenv("SERVER_HOST"); host != "" {
		config.Server.Host = host
	}
	if mode := os.Getenv("SERVER_MODE"); mode != "" {
		config.Server.Mode = mode
	}

	// 数据库配置
	if dbType := os.Getenv("DATABASE_TYPE"); dbType != "" {
		config.Database.Type = dbType
	}

	// JWT配置
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		config.JWT.Secret = secret
	}
	if expire := os.Getenv("JWT_EXPIRE"); expire != "" {
		// 这里可以添加字符串转整数的逻辑
	}

	// 日志配置
	if level := os.Getenv("LOG_LEVEL"); level != "" {
		config.Log.Level = level
	}

	// 日期格式配置
	if dateOnly := os.Getenv("DATE_FORMAT_DATE_ONLY"); dateOnly != "" {
		config.DateFormat.DateOnly = dateOnly
	}
	if dateTime := os.Getenv("DATE_FORMAT_DATE_TIME"); dateTime != "" {
		config.DateFormat.DateTime = dateTime
	}
	if dateTimeFull := os.Getenv("DATE_FORMAT_DATE_TIME_FULL"); dateTimeFull != "" {
		config.DateFormat.DateTimeFull = dateTimeFull
	}
	if timeOnly := os.Getenv("DATE_FORMAT_TIME_ONLY"); timeOnly != "" {
		config.DateFormat.TimeOnly = timeOnly
	}
	if timeShort := os.Getenv("DATE_FORMAT_TIME_SHORT"); timeShort != "" {
		config.DateFormat.TimeShort = timeShort
	}
}
