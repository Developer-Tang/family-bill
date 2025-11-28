package main

import (
	"fmt"
	"log"

	_ "github.com/family-bill/bill-server/docs"
	"github.com/family-bill/bill-server/internal/rest/controller"
	"github.com/family-bill/bill-server/internal/rest/middleware"
	"github.com/family-bill/bill-server/pkg/db"
	"github.com/family-bill/bill-server/pkg/logger"
	"github.com/family-bill/bill-server/pkg/utils"
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
func main() {
	// 加载配置
	config := utils.LoadConfig()

	// 初始化日志
	logger.InitLogger(config)

	// 初始化数据库
	db.InitDB(config)

	// 自动迁移数据库表
	db.AutoMigrate()

	// 设置Gin模式
	if config.Logger.Level == "debug" {
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

	// 添加响应中间件，生成交易ID
	r.Use(middleware.ResponseMiddleware())

	// 添加请求日志中间件
	r.Use(middleware.RequestLogger())

	// 注册路由
	swaggerRoutes(r, config) // 注册Swagger路由

	// API分组
	v1 := r.Group("/api/v1")
	commonRoutes(v1) // 公共路由
	userRoutes(v1)   // 用户相关路由

	// 其他路由...

	// 启动服务器
	addr := fmt.Sprintf(":%d", config.Server.Port)
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// swaggerRoutes 注册Swagger路由
func swaggerRoutes(r *gin.Engine, config *utils.Config) {
	// 配置Swagger
	if config.Swagger.Enabled {
		// 注册Swagger路由，使用gin-swagger提供的默认配置
		r.GET(config.Swagger.Path+"/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		log.Printf("Swagger enabled at %s", config.Swagger.Path)
	}
}

// commonRoutes 公共路由
func commonRoutes(api *gin.RouterGroup) {
	api.GET("/health", func(c *gin.Context) {
		middleware.SuccessResponse(c, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})
}

// userRoutes 用户相关路由
func userRoutes(api *gin.RouterGroup) {
	// 创建用户控制器实例
	userController := controller.NewUserController()

	// 用户路由分组
	user := api.Group("/users")

	// 注册
	user.POST("/register", userController.Register)

	// 登录
	user.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Login route",
		})
	})

	// 获取用户信息
	user.GET("/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get user info route",
		})
	})
}
