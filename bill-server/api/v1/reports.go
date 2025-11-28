package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
)

// GetCategoryReport 获取分类统计报表
// @Summary 获取分类统计报表
// @Description 获取指定账本在特定时间范围内的分类统计数据，支持按收入或支出类型统计
// @Tags 统计报表
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param type query string true "统计类型：income（收入）, expense（支出）"
// @Param period query string false "统计周期：day, week, month, year" default(month)
// @Param date query string true "统计日期，格式：YYYY-MM-DD或YYYY-MM"
// @Param level query int false "分类层级：1（一级分类）, 2（二级分类）" default(1)
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/reports/category [get]
func GetCategoryReport(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	reportType := c.Query("type")
	period := c.DefaultQuery("period", "month")
	dateStr := c.Query("date")
	levelStr := c.DefaultQuery("level", "1")

	// 验证参数
	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的账本ID",
		})
		return
	}

	level, err := strconv.Atoi(levelStr)
	if err != nil || level < 1 || level > 2 {
		level = 1
	}

	// 解析日期
	date, err := time.Parse("2006-01", dateStr)
	if err != nil {
		// 尝试解析完整日期格式
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "无效的日期格式",
			})
			return
		}
	}

	// 获取时间范围
	startDate, endDate := getPeriodRange(date, period)

	// 查询分类统计数据
	type CategoryStat struct {
		CategoryID   uint    `json:"category_id"`
		CategoryName string  `json:"category_name"`
		Amount       float64 `json:"amount"`
		Percentage   float64 `json:"percentage"`
		Count        int64   `json:"count"`
		Icon         string  `json:"icon"`
		Color        string  `json:"color"`
	}
	var categories []CategoryStat

	// 构建查询
	query := database.DB.Table("transactions").Select("categories.category_id, categories.name as category_name, COALESCE(SUM(transactions.amount), 0) as amount, COUNT(transactions.transaction_id) as count, categories.icon, categories.color").
		Joins("LEFT JOIN categories ON transactions.category_id = categories.category_id").
		Where("transactions.book_id = ? AND transactions.type = ? AND transactions.date BETWEEN ? AND ?", bookID, reportType, startDate, endDate)

	// 添加分类层级过滤
	if level == 1 {
		query = query.Where("categories.parent_id IS NULL")
	} else {
		query = query.Where("categories.parent_id IS NOT NULL")
	}

	// 执行查询
	query.Group("categories.category_id, categories.name, categories.icon, categories.color").
		Order("amount DESC").
		Scan(&categories)

	// 查询总金额
	var totalAmount float64
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, reportType, startDate, endDate).Select("COALESCE(SUM(amount), 0)").Scan(&totalAmount)

	// 计算百分比
	for i := range categories {
		if totalAmount > 0 {
			categories[i].Percentage = (categories[i].Amount / totalAmount) * 100
		}
	}

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取分类统计报表成功",
		"data": gin.H{
			"type":         reportType,
			"period":       period,
			"date":         dateStr,
			"total_amount": totalAmount,
			"categories":   categories,
		},
	})
}

// GetCategoryTrendReport 获取分类趋势报表
// @Summary 获取分类趋势报表
// @Description 获取指定账本在特定时间范围内的分类趋势数据，支持按收入或支出类型统计
// @Tags 统计报表
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param type query string true "统计类型：income（收入）, expense（支出）"
// @Param category_id query int true "分类ID"
// @Param period query string false "统计周期：day, week, month, year" default(month)
// @Param start_date query string true "开始日期，格式：YYYY-MM-DD或YYYY-MM"
// @Param end_date query string true "结束日期，格式：YYYY-MM-DD或YYYY-MM"
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/reports/category/trend [get]
func GetCategoryTrendReport(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	reportType := c.Query("type")
	categoryIDStr := c.Query("category_id")
	period := c.DefaultQuery("period", "month")
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	// 验证参数
	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的账本ID",
		})
		return
	}

	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的分类ID",
		})
		return
	}

	// 解析日期
	startDate, err := time.Parse("2006-01", startDateStr)
	if err != nil {
		// 尝试解析完整日期格式
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "无效的开始日期格式",
			})
			return
		}
	}

	endDate, err := time.Parse("2006-01", endDateStr)
	if err != nil {
		// 尝试解析完整日期格式
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "无效的结束日期格式",
			})
			return
		}
	}

	// 生成时间序列
	timeSeries := generateTimeSeries(startDate, endDate, period)

	// 查询分类信息
	var category models.Category
	database.DB.First(&category, categoryID)

	// 查询趋势数据
	type TrendData struct {
		Date   string  `json:"date"`
		Amount float64 `json:"amount"`
		Count  int64   `json:"count"`
	}
	var trendData []TrendData

	// 遍历时间序列，查询每个时间段的数据
	for _, ts := range timeSeries {
		var amount float64
		var count int64
		database.DB.Model(&models.Transaction{}).Where("book_id = ? AND category_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, categoryID, reportType, ts.Start, ts.End).Select("COALESCE(SUM(amount), 0)").Scan(&amount)
		database.DB.Model(&models.Transaction{}).Where("book_id = ? AND category_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, categoryID, reportType, ts.Start, ts.End).Count(&count)
		trendData = append(trendData, TrendData{
			Date:   ts.Label,
			Amount: amount,
			Count:  count,
		})
	}

	// 计算汇总数据
	var totalAmount float64
	var totalCount int64
	var avgAmount float64
	for _, data := range trendData {
		totalAmount += data.Amount
		totalCount += data.Count
	}
	if totalCount > 0 {
		avgAmount = totalAmount / float64(totalCount)
	}

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取分类趋势报表成功",
		"data": gin.H{
			"category_id":   categoryID,
			"category_name": category.Name,
			"type":          reportType,
			"period":        period,
			"start_date":    startDateStr,
			"end_date":      endDateStr,
			"data":          trendData,
			"summary": gin.H{
				"total_amount": totalAmount,
				"total_count":  totalCount,
				"avg_amount":   avgAmount,
			},
		},
	})
}

// GetAccountReport 获取账户收支报表
// @Summary 获取账户收支报表
// @Description 获取指定账本在特定时间范围内的账户收支数据，支持按账户类型统计
// @Tags 统计报表
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param type query string false "数据类型：income（收入）, expense（支出）, both（收支）" default(both)
// @Param period query string false "统计周期：day, week, month, year" default(month)
// @Param date query string true "统计日期，格式：YYYY-MM-DD或YYYY-MM"
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/reports/account [get]
func GetAccountReport(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	dataType := c.DefaultQuery("type", "both")
	period := c.DefaultQuery("period", "month")
	dateStr := c.Query("date")

	// 验证参数
	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的账本ID",
		})
		return
	}

	// 解析日期
	date, err := time.Parse("2006-01", dateStr)
	if err != nil {
		// 尝试解析完整日期格式
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "无效的日期格式",
			})
			return
		}
	}

	// 获取时间范围
	startDate, endDate := getPeriodRange(date, period)

	// 查询账户列表
	var accounts []models.Account
	database.DB.Where("book_id = ?", bookID).Find(&accounts)

	// 构建账户收支数据
	type AccountStat struct {
		AccountID     uint    `json:"account_id"`
		AccountName   string  `json:"account_name"`
		AccountType   string  `json:"account_type"`
		Income        float64 `json:"income"`
		Expense       float64 `json:"expense"`
		BalanceChange float64 `json:"balance_change"`
	}
	var accountStats []AccountStat

	for _, account := range accounts {
		var income, expense float64

		// 查询收入
		if dataType == "both" || dataType == "income" {
			database.DB.Model(&models.Transaction{}).Where("book_id = ? AND account_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, account.AccountID, "income", startDate, endDate).Select("COALESCE(SUM(amount), 0)").Scan(&income)
		}

		// 查询支出
		if dataType == "both" || dataType == "expense" {
			database.DB.Model(&models.Transaction{}).Where("book_id = ? AND account_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, account.AccountID, "expense", startDate, endDate).Select("COALESCE(SUM(amount), 0)").Scan(&expense)
		}

		// 计算余额变化
		balanceChange := income - expense

		// 查询账户类型
		var accountType models.AccountType
		database.DB.First(&accountType, account.AccountTypeID)

		accountStats = append(accountStats, AccountStat{
			AccountID:     account.AccountID,
			AccountName:   account.Name,
			AccountType:   accountType.Name,
			Income:        income,
			Expense:       expense,
			BalanceChange: balanceChange,
		})
	}

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取账户收支报表成功",
		"data": gin.H{
			"type":     dataType,
			"period":   period,
			"date":     dateStr,
			"accounts": accountStats,
		},
	})
}

// GetAccountBalanceReport 获取账户余额报表
// @Summary 获取账户余额报表
// @Description 获取指定账本在特定日期的账户余额报表，包括各账户余额及按账户类型统计的余额分布
// @Tags 统计报表
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param date query string true "统计日期，格式：YYYY-MM-DD"
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/reports/account/balance [get]
func GetAccountBalanceReport(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	dateStr := c.Query("date")

	// 验证参数
	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的账本ID",
		})
		return
	}

	// 验证日期格式
	_, err = time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的日期格式",
		})
		return
	}

	// 查询账户列表
	var accounts []models.Account
	database.DB.Where("book_id = ?", bookID).Find(&accounts)

	// 构建账户余额数据
	type AccountBalance struct {
		AccountID   uint    `json:"account_id"`
		AccountName string  `json:"account_name"`
		AccountType string  `json:"account_type"`
		Balance     float64 `json:"balance"`
		Percentage  float64 `json:"percentage"`
		Currency    string  `json:"currency"`
		Icon        string  `json:"icon"`
		Color       string  `json:"color"`
	}
	var accountBalances []AccountBalance

	// 计算总余额
	var totalBalance float64

	for _, account := range accounts {
		// 查询账户类型
		var accountType models.AccountType
		database.DB.First(&accountType, account.AccountTypeID)

		// 查询账户余额
		var balance float64
		// 这里假设账户余额直接存储在accounts表中
		// 实际实现中可能需要根据交易记录计算
		balance = account.Balance
		totalBalance += balance

		accountBalances = append(accountBalances, AccountBalance{
			AccountID:   account.AccountID,
			AccountName: account.Name,
			AccountType: accountType.Name,
			Balance:     balance,
			Percentage:  0, // 先初始化为0，后面统一计算
			Currency:    account.Currency,
			Icon:        accountType.Icon,
			Color:       accountType.Color,
		})
	}

	// 计算各账户余额占比
	for i := range accountBalances {
		if totalBalance > 0 {
			accountBalances[i].Percentage = (accountBalances[i].Balance / totalBalance) * 100
		}
	}

	// 按账户类型统计余额分布
	type BalanceByType struct {
		Type       string  `json:"type"`
		Amount     float64 `json:"amount"`
		Percentage float64 `json:"percentage"`
	}
	balanceByTypeMap := make(map[string]float64)

	// 统计各类型账户余额
	for _, ab := range accountBalances {
		balanceByTypeMap[ab.AccountType] += ab.Balance
	}

	// 构建按类型统计的余额分布
	var balanceByType []BalanceByType
	for accountType, amount := range balanceByTypeMap {
		percentage := 0.0
		if totalBalance > 0 {
			percentage = (amount / totalBalance) * 100
		}
		balanceByType = append(balanceByType, BalanceByType{
			Type:       accountType,
			Amount:     amount,
			Percentage: percentage,
		})
	}

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取账户余额报表成功",
		"data": gin.H{
			"date":            dateStr,
			"total_balance":   totalBalance,
			"accounts":        accountBalances,
			"balance_by_type": balanceByType,
		},
	})
}

// GetMemberReport 获取成员收支报表
// @Summary 获取成员收支报表
// @Description 获取指定账本在特定时间范围内的成员收支数据，支持按成员统计
// @Tags 统计报表
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param period query string false "统计周期：day, week, month, year" default(month)
// @Param date query string true "统计日期，格式：YYYY-MM-DD或YYYY-MM"
// @Param member_id query int false "成员ID，不填则统计所有成员"
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/reports/member [get]
func GetMemberReport(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	period := c.DefaultQuery("period", "month")
	dateStr := c.Query("date")
	memberIDStr := c.Query("member_id")

	// 验证参数
	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的账本ID",
		})
		return
	}

	// 解析日期
	date, err := time.Parse("2006-01", dateStr)
	if err != nil {
		// 尝试解析完整日期格式
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "无效的日期格式",
			})
			return
		}
	}

	// 获取时间范围
	startDate, endDate := getPeriodRange(date, period)

	// 构建查询
	query := database.DB.Table("transactions").Select("users.user_id as member_id, users.username as member_name, COALESCE(SUM(CASE WHEN transactions.type = 'income' THEN transactions.amount ELSE 0 END), 0) as income, COALESCE(SUM(CASE WHEN transactions.type = 'expense' THEN transactions.amount ELSE 0 END), 0) as expense").
		Joins("LEFT JOIN users ON transactions.user_id = users.user_id").
		Where("transactions.book_id = ? AND transactions.date BETWEEN ? AND ?", bookID, startDate, endDate)

	// 添加成员ID筛选
	if memberIDStr != "" {
		memberID, err := strconv.ParseUint(memberIDStr, 10, 32)
		if err == nil {
			query = query.Where("users.user_id = ?", memberID)
		}
	}

	// 执行查询
	type MemberStat struct {
		MemberID   uint    `json:"member_id"`
		MemberName string  `json:"member_name"`
		Income     float64 `json:"income"`
		Expense    float64 `json:"expense"`
		Balance    float64 `json:"balance"`
	}
	var memberStats []MemberStat
	query.Group("users.user_id, users.username").Scan(&memberStats)

	// 计算余额
	for i := range memberStats {
		memberStats[i].Balance = memberStats[i].Income - memberStats[i].Expense
	}

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成员收支报表成功",
		"data": gin.H{
			"period":  period,
			"date":    dateStr,
			"members": memberStats,
		},
	})
}

// GetMemberContributionReport 获取成员贡献度分析
// @Summary 获取成员贡献度分析
// @Description 获取指定账本在特定时间范围内的成员贡献度数据，分析各成员的财务贡献
// @Tags 统计报表
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param period query string false "统计周期：day, week, month, year" default(month)
// @Param date query string true "统计日期，格式：YYYY-MM-DD或YYYY-MM"
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/reports/member/contribution [get]
func GetMemberContributionReport(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	period := c.DefaultQuery("period", "month")
	dateStr := c.Query("date")

	// 验证参数
	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的账本ID",
		})
		return
	}

	// 解析日期
	date, err := time.Parse("2006-01", dateStr)
	if err != nil {
		// 尝试解析完整日期格式
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    400,
				"message": "无效的日期格式",
			})
			return
		}
	}

	// 获取时间范围
	startDate, endDate := getPeriodRange(date, period)

	// 查询总收支
	var totalIncome, totalExpense float64
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "income", startDate, endDate).Select("COALESCE(SUM(amount), 0)").Scan(&totalIncome)
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "expense", startDate, endDate).Select("COALESCE(SUM(amount), 0)").Scan(&totalExpense)

	// 查询成员贡献度
	type MemberContribution struct {
		MemberID          uint    `json:"member_id"`
		MemberName        string  `json:"member_name"`
		Income            float64 `json:"income"`
		IncomePercentage  float64 `json:"income_percentage"`
		Expense           float64 `json:"expense"`
		ExpensePercentage float64 `json:"expense_percentage"`
		NetContribution   float64 `json:"net_contribution"`
	}
	var memberContributions []MemberContribution

	database.DB.Table("transactions").Select("users.user_id as member_id, users.username as member_name, COALESCE(SUM(CASE WHEN transactions.type = 'income' THEN transactions.amount ELSE 0 END), 0) as income, COALESCE(SUM(CASE WHEN transactions.type = 'expense' THEN transactions.amount ELSE 0 END), 0) as expense").
		Joins("LEFT JOIN users ON transactions.user_id = users.user_id").
		Where("transactions.book_id = ? AND transactions.date BETWEEN ? AND ?", bookID, startDate, endDate).
		Group("users.user_id, users.username").
		Scan(&memberContributions)

	// 计算百分比和净贡献
	for i := range memberContributions {
		// 计算收入占比
		if totalIncome > 0 {
			memberContributions[i].IncomePercentage = (memberContributions[i].Income / totalIncome) * 100
		}

		// 计算支出占比
		if totalExpense > 0 {
			memberContributions[i].ExpensePercentage = (memberContributions[i].Expense / totalExpense) * 100
		}

		// 计算净贡献
		memberContributions[i].NetContribution = memberContributions[i].Income - memberContributions[i].Expense
	}

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成员贡献度分析成功",
		"data": gin.H{
			"period":        period,
			"date":          dateStr,
			"total_income":  totalIncome,
			"total_expense": totalExpense,
			"members":       memberContributions,
		},
	})
}

// GetBudgetReport 获取预算执行报表
// @Summary 获取预算执行报表
// @Description 获取指定账本在特定时间范围内的预算执行情况，支持按分类统计
// @Tags 统计报表
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param period query string false "统计周期：month, quarter, year" default(month)
// @Param date query string true "统计日期，格式：YYYY-MM"
// @Param category_id query int false "分类ID，不填则统计所有分类"
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/reports/budget [get]
func GetBudgetReport(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	period := c.DefaultQuery("period", "month")
	dateStr := c.Query("date")
	categoryIDStr := c.Query("category_id")

	// 验证参数
	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的账本ID",
		})
		return
	}

	// 解析日期
	date, err := time.Parse("2006-01", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的日期格式",
		})
		return
	}

	// 获取时间范围
	startDate, endDate := getPeriodRange(date, period)

	// 查询总预算
	var totalBudget float64
	query := database.DB.Model(&models.Budget{}).Where("book_id = ? AND period = ? AND date = ?", bookID, period, dateStr)
	if categoryIDStr != "" {
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
		if err == nil {
			query = query.Where("category_id = ?", categoryID)
		}
	}
	query.Select("COALESCE(SUM(amount), 0)").Scan(&totalBudget)

	// 查询已使用预算
	var totalUsed float64
	usedQuery := database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "expense", startDate, endDate)
	if categoryIDStr != "" {
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
		if err == nil {
			usedQuery = usedQuery.Where("category_id = ?", categoryID)
		}
	}
	usedQuery.Select("COALESCE(SUM(amount), 0)").Scan(&totalUsed)

	// 计算剩余预算和使用率
	totalRemaining := totalBudget - totalUsed
	overallUsageRate := 0.0
	if totalBudget > 0 {
		overallUsageRate = (totalUsed / totalBudget) * 100
	}

	// 查询分类预算执行情况
	type CategoryBudgetStat struct {
		CategoryID   uint    `json:"category_id"`
		CategoryName string  `json:"category_name"`
		Budget       float64 `json:"budget"`
		Used         float64 `json:"used"`
		Remaining    float64 `json:"remaining"`
		UsageRate    float64 `json:"usage_rate"`
		IsOverBudget bool    `json:"is_over_budget"`
	}
	var categoryBudgetStats []CategoryBudgetStat

	// 查询分类预算
	var categoryBudgets []struct {
		CategoryID uint    `json:"category_id"`
		Name       string  `json:"name"`
		Amount     float64 `json:"amount"`
	}
	budgetQuery := database.DB.Table("budgets").Select("budgets.category_id, categories.name, budgets.amount").
		Joins("LEFT JOIN categories ON budgets.category_id = categories.category_id").
		Where("budgets.book_id = ? AND budgets.period = ? AND budgets.date = ?", bookID, period, dateStr)
	if categoryIDStr != "" {
		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
		if err == nil {
			budgetQuery = budgetQuery.Where("budgets.category_id = ?", categoryID)
		}
	}
	budgetQuery.Scan(&categoryBudgets)

	// 查询每个分类的使用情况
	for _, cb := range categoryBudgets {
		var used float64
		database.DB.Model(&models.Transaction{}).Where("book_id = ? AND category_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, cb.CategoryID, "expense", startDate, endDate).Select("COALESCE(SUM(amount), 0)").Scan(&used)

		remaining := cb.Amount - used
		usageRate := 0.0
		if cb.Amount > 0 {
			usageRate = (used / cb.Amount) * 100
		}
		isOverBudget := used > cb.Amount

		categoryBudgetStats = append(categoryBudgetStats, CategoryBudgetStat{
			CategoryID:   cb.CategoryID,
			CategoryName: cb.Name,
			Budget:       cb.Amount,
			Used:         used,
			Remaining:    remaining,
			UsageRate:    usageRate,
			IsOverBudget: isOverBudget,
		})
	}

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取预算执行报表成功",
		"data": gin.H{
			"period":             period,
			"date":               dateStr,
			"total_budget":       totalBudget,
			"total_used":         totalUsed,
			"total_remaining":    totalRemaining,
			"overall_usage_rate": overallUsageRate,
			"categories":         categoryBudgetStats,
		},
	})
}

// GetBudgetAlertReport 获取预算超支提醒
// @Summary 获取预算超支提醒
// @Description 获取指定账本的预算超支提醒信息，支持按分类和成员筛选
// @Tags 统计报表
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param period query string false "统计周期：month, quarter, year" default(month)
// @Param date query string false "统计日期，格式：YYYY-MM" default(当前日期)
// @Param threshold query number false "提醒阈值，百分比（0-100）" default(80)
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/reports/budget/alert [get]
func GetBudgetAlertReport(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	period := c.DefaultQuery("period", "month")
	dateStr := c.DefaultQuery("date", time.Now().Format("2006-01"))
	thresholdStr := c.DefaultQuery("threshold", "80")

	// 验证参数
	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的账本ID",
		})
		return
	}

	threshold, err := strconv.ParseFloat(thresholdStr, 64)
	if err != nil || threshold < 0 || threshold > 100 {
		threshold = 80
	}

	// 解析日期
	date, err := time.Parse("2006-01", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的日期格式",
		})
		return
	}

	// 获取时间范围
	startDate, endDate := getPeriodRange(date, period)

	// 查询预算超支提醒
	type BudgetAlert struct {
		CategoryID   uint    `json:"category_id"`
		CategoryName string  `json:"category_name"`
		Budget       float64 `json:"budget"`
		Used         float64 `json:"used"`
		UsageRate    float64 `json:"usage_rate"`
		OverAmount   float64 `json:"over_amount"`
	}
	var budgetAlerts []BudgetAlert

	// 查询分类预算
	var categoryBudgets []struct {
		CategoryID uint    `json:"category_id"`
		Name       string  `json:"name"`
		Amount     float64 `json:"amount"`
	}
	database.DB.Table("budgets").Select("budgets.category_id, categories.name, budgets.amount").
		Joins("LEFT JOIN categories ON budgets.category_id = categories.category_id").
		Where("budgets.book_id = ? AND budgets.period = ? AND budgets.date = ?", bookID, period, dateStr).
		Scan(&categoryBudgets)

	// 查询每个分类的使用情况并检查是否超支
	for _, cb := range categoryBudgets {
		var used float64
		database.DB.Model(&models.Transaction{}).Where("book_id = ? AND category_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, cb.CategoryID, "expense", startDate, endDate).Select("COALESCE(SUM(amount), 0)").Scan(&used)

		usageRate := 0.0
		if cb.Amount > 0 {
			usageRate = (used / cb.Amount) * 100
		}

		// 检查是否超过阈值
		if usageRate >= threshold {
			overAmount := used - cb.Amount
			if overAmount < 0 {
				overAmount = 0
			}

			budgetAlerts = append(budgetAlerts, BudgetAlert{
				CategoryID:   cb.CategoryID,
				CategoryName: cb.Name,
				Budget:       cb.Amount,
				Used:         used,
				UsageRate:    usageRate,
				OverAmount:   overAmount,
			})
		}
	}

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取预算超支提醒成功",
		"data": gin.H{
			"period":    period,
			"date":      dateStr,
			"threshold": threshold,
			"alerts":    budgetAlerts,
		},
	})
}
