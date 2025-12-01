package constant

import "time"

// TradeIDKey 交易ID上下文键
const TradeIDKey = "trade_id"

// ClientTypeKey 客户端类型上下文键
const ClientTypeKey = "client_type"

// UserIDKey 用户ID上下文键
const UserIDKey = "user_id"

// ClientType 客户端类型
const (
	ClientTypeApp     = "app"     // 移动端
	ClientTypeWeb     = "web"     // 网页端
	ClientTypeH5      = "h5"      // 管理端
	ClientTypeUnknown = "unknown" // 未知客户端类型
)

/* ==================== Token相关 ==================== */

// TokenExpire 令牌有效期
const (
	TokenExpireWeb     = 30 * time.Minute    // 网页端令牌有效期（默认30分钟）
	TokenExpireApp     = 7 * 24 * time.Hour  // 管理端令牌有效期（默认7天）
	TokenExpireRefresh = 30 * 24 * time.Hour // 刷新令牌有效期（默认30天）
)
