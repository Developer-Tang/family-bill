package router

import (
	"github.com/gin-gonic/gin"

	"github.com/family-bill/bill-server/api"
	v1 "github.com/family-bill/bill-server/api/v1"
	"github.com/family-bill/bill-server/internal/config"
)

// SetupRouter 设置路由
func SetupRouter(cfg *config.Config) *gin.Engine {
	// 创建 Gin 引擎
	r := gin.Default()

	// 健康检查路由
	r.GET("/health", func(c *gin.Context) {
		c.Set("version", cfg.Version)
		api.HealthCheck(c)
	})

	// API 路由组
	apiGroup := r.Group("/api/v1")
	{
		// 认证授权
		auth := apiGroup.Group("/auth")
		{
			auth.POST("/login", v1.Login)
			auth.POST("/logout", v1.Logout)
			auth.POST("/register", v1.Register)
			auth.POST("/refresh", v1.RefreshToken)
			auth.POST("/verify-email", v1.VerifyEmail)
			auth.POST("/reset-password", v1.ResetPassword)
			auth.POST("/send-code", v1.SendCode)
		}

		// 帮助文档
		help := apiGroup.Group("/help")
		{
			help.GET("/articles", v1.GetHelpArticles)
		}

		// 用户管理
		users := apiGroup.Group("/users")
		{
			users.GET("/profile", v1.GetProfile)
			users.PUT("/profile", v1.UpdateProfile)
			users.POST("/avatar", v1.UpdateAvatar)
		}

		// 家庭组管理
		families := apiGroup.Group("/families")
		{
			families.GET("", v1.GetFamilies)
			families.POST("", v1.CreateFamily)
			families.PUT("/:id", v1.UpdateFamily)
			families.DELETE("/:id", v1.DeleteFamily)
			families.POST("/:id/invite", v1.InviteFamilyMember)
			families.POST("/:id/leave", v1.LeaveFamily)
			families.GET("/:id/members", v1.GetFamilyMembers)
			families.DELETE("/:id/members/:userId", v1.RemoveFamilyMember)
		}

		// 账本管理
		books := apiGroup.Group("/books")
		{
			books.GET("", v1.GetBooks)
			books.POST("", v1.CreateBook)
			books.PUT("/:id", v1.UpdateBook)
			books.DELETE("/:id", v1.DeleteBook)
		}

		// 权限控制
		roles := apiGroup.Group("/roles")
		{
			roles.GET("", v1.GetRoles)
			roles.POST("", v1.CreateRole)
			roles.PUT("/:id", v1.UpdateRole)
			roles.DELETE("/:id", v1.DeleteRole)
		}

		// 系统安全
		logs := apiGroup.Group("/logs")
		{
			logs.GET("/operations", v1.GetOperationLogs)
		}

		security := apiGroup.Group("/security")
		{
			security.POST("/2fa", v1.SetupTwoFactor)
			security.GET("/devices", v1.GetLoginDevices)
			security.DELETE("/devices/:id", v1.LogoutDevice)
		}

		// 账户类型管理
		accountTypes := apiGroup.Group("/account-types")
		{
			accountTypes.GET("", v1.GetAccountTypes)
			accountTypes.POST("", v1.CreateAccountType)
			accountTypes.PUT("/:id", v1.UpdateAccountType)
			accountTypes.DELETE("/:id", v1.DeleteAccountType)
		}

		// 账户管理
		accounts := apiGroup.Group("/accounts")
		{
			accounts.GET("", v1.GetAccounts)
			accounts.POST("", v1.CreateAccount)
			accounts.GET("/:id", v1.GetAccount)
			accounts.PUT("/:id", v1.UpdateAccount)
			accounts.DELETE("/:id", v1.DeleteAccount)
			accounts.PUT("/:id/balance", v1.AdjustAccountBalance)
			accounts.PUT("/:id/status", v1.UpdateAccountStatus)
		}

		// 账户分组管理
		accountGroups := apiGroup.Group("/account-groups")
		{
			accountGroups.GET("", v1.GetAccountGroups)
			accountGroups.POST("", v1.CreateAccountGroup)
			accountGroups.PUT("/:id", v1.UpdateAccountGroup)
			accountGroups.DELETE("/:id", v1.DeleteAccountGroup)
			accountGroups.POST("/:id/accounts", v1.AddAccountsToGroup)
		}

		// 转账管理
		transfers := apiGroup.Group("/transfers")
		{
			transfers.POST("", v1.CreateTransfer)
			transfers.GET("", v1.GetTransfers)
			transfers.GET("/:id", v1.GetTransfer)
		}

		// 货币和汇率管理
		currencies := apiGroup.Group("/currencies")
		{
			currencies.GET("", v1.GetCurrencies)
		}

		exchangeRates := apiGroup.Group("/exchange-rates")
		{
			exchangeRates.GET("", v1.GetExchangeRates)
			exchangeRates.POST("", v1.UpdateExchangeRates)
		}

		// 收支记录管理
		transactions := apiGroup.Group("/transactions")
		{
			transactions.GET("", v1.GetTransactions)
			transactions.POST("", v1.CreateTransaction)
			transactions.GET("/:id", v1.GetTransaction)
			transactions.PUT("/:id", v1.UpdateTransaction)
			transactions.DELETE("/:id", v1.DeleteTransaction)
			transactions.PUT("/:id/lock", v1.LockTransaction)
			transactions.POST("/batch", v1.BatchCreateTransactions)
			transactions.POST("/import", v1.ImportTransactions)
		}

		// 快速记账
		quickTransactions := apiGroup.Group("/quick-transactions")
		{
			quickTransactions.POST("", v1.CreateQuickTransaction)
		}

		// 记账模板
		templates := apiGroup.Group("/transaction-templates")
		{
			templates.GET("", v1.GetTransactionTemplates)
			templates.POST("", v1.CreateTransactionTemplate)
			templates.PUT("/:id", v1.UpdateTransactionTemplate)
			templates.DELETE("/:id", v1.DeleteTransactionTemplate)
		}

		// 周期记账
		recurringTransactions := apiGroup.Group("/recurring-transactions")
		{
			recurringTransactions.GET("", v1.GetRecurringTransactions)
			recurringTransactions.POST("", v1.CreateRecurringTransaction)
			recurringTransactions.PUT("/:id", v1.UpdateRecurringTransaction)
			recurringTransactions.DELETE("/:id", v1.DeleteRecurringTransaction)
			recurringTransactions.POST("/:id/trigger", v1.TriggerRecurringTransaction)
		}

		// 标签管理
		tags := apiGroup.Group("/tags")
		{
			tags.GET("", v1.GetTags)
			tags.POST("", v1.CreateTag)
			tags.PUT("/:id", v1.UpdateTag)
			tags.DELETE("/:id", v1.DeleteTag)
		}

		// 数据看板
		dashboard := apiGroup.Group("/dashboard")
		{
			dashboard.GET("/summary", v1.GetDashboardSummary)
			dashboard.GET("/quick-stats", v1.GetDashboardQuickStats)
			dashboard.GET("/recent-transactions", v1.GetDashboardRecentTransactions)
			dashboard.GET("/budget-progress", v1.GetDashboardBudgetProgress)
		}

		// 收支分析
		analysis := apiGroup.Group("/analysis")
		{
			analysis.GET("/income-expense", v1.GetIncomeExpenseAnalysis)
			analysis.GET("/trend", v1.GetTrendAnalysis)
			analysis.GET("/flow", v1.GetFlowAnalysis)
		}

		// 统计报表
		reports := apiGroup.Group("/reports")
		{
			reports.GET("/category", v1.GetCategoryReport)
			reports.GET("/category/trend", v1.GetCategoryTrendReport)
			reports.GET("/account", v1.GetAccountReport)
			reports.GET("/account/balance", v1.GetAccountBalanceReport)
			reports.GET("/member", v1.GetMemberReport)
			reports.GET("/member/contribution", v1.GetMemberContributionReport)
			reports.GET("/budget", v1.GetBudgetReport)
			reports.GET("/budget/alert", v1.GetBudgetAlertReport)
		}

		// 图表数据
		charts := apiGroup.Group("/charts")
		{
			charts.GET("/line", v1.GetLineChartData)
			charts.GET("/pie", v1.GetPieChartData)
			charts.GET("/bar", v1.GetBarChartData)
			charts.GET("/radar", v1.GetRadarChartData)
		}
	}

	return r
}
