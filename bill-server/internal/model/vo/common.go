package vo

// R 统一响应结构体
type R struct {
	Code    int         `json:"code"`     // 状态码，0表示成功，其他表示错误
	Msg     string      `json:"msg"`      // 消息
	Data    interface{} `json:"data"`     // 数据
	Time    string      `json:"time"`     // 时间
	TradeID string      `json:"trade_id"` // 交易ID
}
