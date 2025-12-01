package bo

// RegisterBo 注册请求
type RegisterBo struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginBo 登录请求
// @Summary 登录请求
// @Description 用户登录请求参数
// @Tags 用户管理
type LoginBo struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RefreshTokenBo 刷新token请求
// @Summary 刷新token请求
// @Description 刷新token请求参数
// @Tags 用户管理
type RefreshTokenBo struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
