package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config 全局配置结构体
type Config struct {
	Server  ServerConfig  `mapstructure:"server"`
	DB      DBConfig      `mapstructure:"db"`
	JWT     JWTConfig     `mapstructure:"jwt"`
	Swagger SwaggerConfig `mapstructure:"swagger"`
	Logger  LoggerConfig  `mapstructure:"logger"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port int `mapstructure:"port"`
}

// DBConfig 数据库配置
type DBConfig struct {
	Driver  string `mapstructure:"driver"`
	DSN     string `mapstructure:"dsn"`
	MaxIdle int    `mapstructure:"max_idle"`
	MaxOpen int    `mapstructure:"max_open"`
	MaxLife int    `mapstructure:"max_life"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret string `mapstructure:"secret"`
}

// SwaggerConfig Swagger配置
type SwaggerConfig struct {
	Enabled     bool   `mapstructure:"enabled"`
	Path        string `mapstructure:"path"`
	Title       string `mapstructure:"title"`
	Version     string `mapstructure:"version"`
	Description string `mapstructure:"description"`
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	Level      string `mapstructure:"level"`
	Path       string `mapstructure:"path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

var (
	YamlConfig *Config
)

// LoadConfig 加载配置
func LoadConfig() *Config {
	if YamlConfig != nil {
		return YamlConfig
	}

	viper.AddConfigPath("./data")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}
	YamlConfig = &Config{}
	if err := viper.Unmarshal(YamlConfig); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	return YamlConfig
}
