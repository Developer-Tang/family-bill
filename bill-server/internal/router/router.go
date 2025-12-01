package router

import (
	"log"

	"github.com/family-bill/bill-server/internal/config"
	"github.com/family-bill/bill-server/internal/handler"
	"github.com/family-bill/bill-server/internal/rest"
	"github.com/family-bill/bill-server/internal/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Family Bill API
// @version 1.0.0
// @description 家庭账单管理系统API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath /api/v1
func InitRouter() *gin.Engine {
	// 设置Gin模式
	if config.YamlConfig.Logger.Level == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建Gin引擎
	r := gin.Default()

	// 配置CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 添加请求过滤中间件
	r.Use(handler.InitReqAroundHandler())

	// 注册路由
	swaggerRoutes(r)         // 注册Swagger路由
	v1 := r.Group("/api/v1") // API分组
	commonRoutes(v1)         // 公共路由
	userRoutes(v1)           // 用户相关路由

	return r
}

// swaggerRoutes 注册Swagger路由
func swaggerRoutes(r *gin.Engine) {
	// 配置Swagger
	if config.YamlConfig.Swagger.Enabled {
		// 注册Swagger路由，使用gin-swagger提供的默认配置
		r.GET(config.YamlConfig.Swagger.Path+"/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		log.Printf("Swagger enabled at %s", config.YamlConfig.Swagger.Path)
	}
}

// commonRoutes 公共路由
func commonRoutes(api *gin.RouterGroup) {
	api.GET("/health", func(c *gin.Context) {
		util.RespOk(c, "Server is running")
	})
}

// userRoutes 用户相关路由
func userRoutes(api *gin.RouterGroup) {
	// 创建用户控制器实例
	rest := rest.InitUserRest()

	// 用户路由分组
	userApi := api.Group("/users")

	// 注册
	userApi.POST("/register", rest.Register)

	// 登录
	userApi.POST("/login", rest.Login)

	// 刷新token
	userApi.POST("/refresh-token", rest.RefreshToken)

	// 获取用户信息
	userApi.POST("/me", rest.GetUserInfo)
}
