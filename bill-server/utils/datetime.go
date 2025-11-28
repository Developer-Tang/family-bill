package utils

import (
	"time"

	"github.com/family-bill/bill-server/internal/config"
)

// DateTimeFormatter 日期时间格式化器
type DateTimeFormatter struct {
	cfg *config.DateFormatConfig
}

// NewDateTimeFormatter 创建日期时间格式化器
func NewDateTimeFormatter(cfg *config.DateFormatConfig) *DateTimeFormatter {
	return &DateTimeFormatter{
		cfg: cfg,
	}
}

// FormatDateOnly 格式化日期（yyyy-MM-dd）
func (f *DateTimeFormatter) FormatDateOnly(t time.Time) string {
	return t.Format(f.cfg.DateOnly)
}

// FormatDateTime 格式化日期时间（yyyy-MM-dd HH:mm）
func (f *DateTimeFormatter) FormatDateTime(t time.Time) string {
	return t.Format(f.cfg.DateTime)
}

// FormatDateTimeFull 格式化完整日期时间（yyyy-MM-dd HH:mm:ss）
func (f *DateTimeFormatter) FormatDateTimeFull(t time.Time) string {
	return t.Format(f.cfg.DateTimeFull)
}

// FormatTimeOnly 格式化时间（HH:mm:ss）
func (f *DateTimeFormatter) FormatTimeOnly(t time.Time) string {
	return t.Format(f.cfg.TimeOnly)
}

// FormatTimeShort 格式化短时间（HH:mm）
func (f *DateTimeFormatter) FormatTimeShort(t time.Time) string {
	return t.Format(f.cfg.TimeShort)
}

// ParseDateOnly 解析日期（yyyy-MM-dd）
func (f *DateTimeFormatter) ParseDateOnly(s string) (time.Time, error) {
	return time.Parse(f.cfg.DateOnly, s)
}

// ParseDateTime 解析日期时间（yyyy-MM-dd HH:mm）
func (f *DateTimeFormatter) ParseDateTime(s string) (time.Time, error) {
	return time.Parse(f.cfg.DateTime, s)
}

// ParseDateTimeFull 解析完整日期时间（yyyy-MM-dd HH:mm:ss）
func (f *DateTimeFormatter) ParseDateTimeFull(s string) (time.Time, error) {
	return time.Parse(f.cfg.DateTimeFull, s)
}

// ParseTimeOnly 解析时间（HH:mm:ss）
func (f *DateTimeFormatter) ParseTimeOnly(s string) (time.Time, error) {
	return time.Parse(f.cfg.TimeOnly, s)
}

// ParseTimeShort 解析短时间（HH:mm）
func (f *DateTimeFormatter) ParseTimeShort(s string) (time.Time, error) {
	return time.Parse(f.cfg.TimeShort, s)
}

// ValidateDateFormat 验证日期格式是否有效
func (f *DateTimeFormatter) ValidateDateFormat(format string) bool {
	// 使用当前时间测试格式化
	now := time.Now()
	formatted := now.Format(format)
	_, err := time.Parse(format, formatted)
	return err == nil
}
