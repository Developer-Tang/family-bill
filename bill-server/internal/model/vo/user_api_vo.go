package vo

// AuthVO 登录响应
type AuthVO struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Avatar       string `json:"avatar"`
	Role         string `json:"role"`
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// UserVO 用户信息
type UserVO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Role     string `json:"role"`
}
