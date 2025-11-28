package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
)

// GetDashboardSummary 获取财务概览数据
// @Summary 获取财务概览数据
// @Description 获取指定账本的财务概览数据，包括总收入、总支出、余额、同比增长率等
// @Tags 数据看板
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param period query string false "统计周期：day, week, month, year" default(month)
// @Param date query string false "统计日期，格式：YYYY-MM-DD或YYYY-MM" default(当前日期)
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/dashboard/summary [get]
func GetDashboardSummary(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	period := c.DefaultQuery("period", "month")
	dateStr := c.DefaultQuery("date", time.Now().Format("2006-01"))

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

	// 获取当前周期的开始和结束时间
	startDate, endDate := getPeriodRange(date, period)

	// 获取上一周期的开始和结束时间
	prevStartDate, prevEndDate := getPreviousPeriodRange(date, period)

	// 查询当前周期的收入和支出
	var currentIncome, currentExpense float64
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "income", startDate, endDate).Select("COALESCE(SUM(amount), 0)").Scan(&currentIncome)
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "expense", startDate, endDate).Select("COALESCE(SUM(amount), 0)").Scan(&currentExpense)

	// 查询上一周期的收入和支出
	var prevIncome, prevExpense float64
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "income", prevStartDate, prevEndDate).Select("COALESCE(SUM(amount), 0)").Scan(&prevIncome)
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "expense", prevStartDate, prevEndDate).Select("COALESCE(SUM(amount), 0)").Scan(&prevExpense)

	// 计算同比增长率
	incomeRate := calculateGrowthRate(currentIncome, prevIncome)
	expenseRate := calculateGrowthRate(currentExpense, prevExpense)

	// 查询最高支出分类
	type CategoryExpense struct {
		CategoryID   uint    `json:"id"`
		CategoryName string  `json:"name"`
		Amount       float64 `json:"amount"`
	}
	var topExpenseCategory CategoryExpense
	database.DB.Table("transactions").Select("categories.category_id, categories.name as category_name, COALESCE(SUM(transactions.amount), 0) as amount").
		Joins("LEFT JOIN categories ON transactions.category_id = categories.category_id").
		Where("transactions.book_id = ? AND transactions.type = ? AND transactions.date BETWEEN ? AND ?", bookID, "expense", startDate, endDate).
		Group("categories.category_id, categories.name").
		Order("amount DESC").
		Limit(1).
		Scan(&topExpenseCategory)

	// 计算最高支出分类占比
	percentage := 0.0
	if currentExpense > 0 {
		percentage = (topExpenseCategory.Amount / currentExpense) * 100
	}

	// 查询账户汇总信息
	var totalBalance float64
	var accountCount int64
	database.DB.Model(&models.Account{}).Where("book_id = ?", bookID).Select("COALESCE(SUM(balance), 0)").Scan(&totalBalance)
	database.DB.Model(&models.Account{}).Where("book_id = ?", bookID).Count(&accountCount)

	// 查询预算执行情况
	var totalBudget, usedBudget float64
	// 这里假设预算表存在，实际实现时需要根据数据库结构调整
	database.DB.Model(&models.Budget{}).Where("book_id = ? AND period = ? AND date = ?", bookID, period, dateStr).Select("COALESCE(SUM(amount), 0)").Scan(&totalBudget)
	usedBudget = currentExpense

	budgetPercentage := 0.0
	if totalBudget > 0 {
		budgetPercentage = (usedBudget / totalBudget) * 100
	}

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取财务概览数据成功",
		"data": gin.H{
			"total_income":  currentIncome,
			"total_expense": currentExpense,
			"balance":       currentIncome - currentExpense,
			"income_rate":   incomeRate,
			"expense_rate":  expenseRate,
			"top_expense_category": gin.H{
				"id":         topExpenseCategory.CategoryID,
				"name":       topExpenseCategory.CategoryName,
				"amount":     topExpenseCategory.Amount,
				"percentage": percentage,
			},
			"budget_status": gin.H{
				"total":      totalBudget,
				"used":       usedBudget,
				"percentage": budgetPercentage,
				"alert":      budgetPercentage >= 80,
			},
			"account_summary": gin.H{
				"total_balance": totalBalance,
				"account_count": accountCount,
			},
		},
	})
}

// GetDashboardQuickStats 获取快速统计数据
// @Summary 获取快速统计数据
// @Description 获取指定账本的快速统计数据，包括收支笔数、平均收支等
// @Tags 数据看板
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param period query string false "统计周期：day, week, month, year" default(month)
// @Param date query string false "统计日期，格式：YYYY-MM-DD或YYYY-MM" default(当前日期)
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/dashboard/quick-stats [get]
func GetDashboardQuickStats(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	period := c.DefaultQuery("period", "month")
	dateStr := c.DefaultQuery("date", time.Now().Format("2006-01"))

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

	// 获取当前周期的开始和结束时间
	startDate, endDate := getPeriodRange(date, period)

	// 查询收入笔数和平均收入
	var incomeCount int64
	var avgIncome float64
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "income", startDate, endDate).Count(&incomeCount)
	if incomeCount > 0 {
		database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "income", startDate, endDate).Select("COALESCE(AVG(amount), 0)").Scan(&avgIncome)
	}

	// 查询支出笔数和平均支出
	var expenseCount int64
	var avgExpense float64
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "expense", startDate, endDate).Count(&expenseCount)
	if expenseCount > 0 {
		database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "expense", startDate, endDate).Select("COALESCE(AVG(amount), 0)").Scan(&avgExpense)
	}

	// 计算总交易笔数
	totalTransactions := incomeCount + expenseCount

	// 查询最高单笔收入和支出
	var highestIncome, highestExpense float64
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "income", startDate, endDate).Select("COALESCE(MAX(amount), 0)").Scan(&highestIncome)
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "expense", startDate, endDate).Select("COALESCE(MAX(amount), 0)").Scan(&highestExpense)

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取快速统计数据成功",
		"data": gin.H{
			"income_count":       incomeCount,
			"expense_count":      expenseCount,
			"avg_income":         avgIncome,
			"avg_expense":        avgExpense,
			"total_transactions": totalTransactions,
			"highest_income":     highestIncome,
			"highest_expense":    highestExpense,
		},
	})
}

// GetDashboardRecentTransactions 获取最近收支记录
// @Summary 获取最近收支记录
// @Description 获取指定账本的最近收支记录，支持分页和筛选
// @Tags 数据看板
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param limit query int false "每页记录数，最大50" default(10)
// @Param offset query int false "偏移量" default(0)
// @Param type query string false "交易类型：income（收入）, expense（支出）"
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/dashboard/recent-transactions [get]
func GetDashboardRecentTransactions(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	transactionType := c.Query("type")

	// 验证参数
	bookID, err := strconv.ParseUint(bookIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的账本ID",
		})
		return
	}

	limit, err := strconv.ParseUint(limitStr, 10, 32)
	if err != nil || limit > 50 {
		limit = 10
	}

	offset, err := strconv.ParseUint(offsetStr, 10, 32)
	if err != nil {
		offset = 0
	}

	// 构建查询
	query := database.DB.Table("transactions").Select("transactions.transaction_id as id, transactions.amount, transactions.type, categories.name as category_name, accounts.name as account_name, transactions.description as remark, transactions.created_at").
		Joins("LEFT JOIN categories ON transactions.category_id = categories.category_id").
		Joins("LEFT JOIN accounts ON transactions.account_id = accounts.account_id").
		Where("transactions.book_id = ?", bookID)

	// 添加交易类型筛选
	if transactionType != "" {
		query = query.Where("transactions.type = ?", transactionType)
	}

	// 查询总记录数
	var total int64
	query.Count(&total)

	// 查询记录列表
	type RecentTransaction struct {
		ID           uint      `json:"id"`
		Amount       float64   `json:"amount"`
		Type         string    `json:"type"`
		CategoryName string    `json:"category_name"`
		AccountName  string    `json:"account_name"`
		Remark       string    `json:"remark"`
		CreatedAt    time.Time `json:"created_at"`
	}
	var transactions []RecentTransaction
	query.Order("transactions.created_at DESC").Limit(int(limit)).Offset(int(offset)).Scan(&transactions)

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取最近收支记录成功",
		"data": gin.H{
			"list":   transactions,
			"total":  total,
			"limit":  limit,
			"offset": offset,
		},
	})
}

// GetDashboardBudgetProgress 获取预算执行进度
// @Summary 获取预算执行进度
// @Description 获取指定账本的预算执行进度数据
// @Tags 数据看板
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param period query string false "统计周期：month, quarter, year" default(month)
// @Param date query string false "统计日期，格式：YYYY-MM" default(当前日期)
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/dashboard/budget-progress [get]
func GetDashboardBudgetProgress(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	period := c.DefaultQuery("period", "month")
	dateStr := c.DefaultQuery("date", time.Now().Format("2006-01"))

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

	// 获取当前周期的开始和结束时间
	startDate, endDate := getPeriodRange(date, period)

	// 查询总预算
	var totalBudget float64
	database.DB.Model(&models.Budget{}).Where("book_id = ? AND period = ? AND date = ?", bookID, period, dateStr).Select("COALESCE(SUM(amount), 0)").Scan(&totalBudget)

	// 查询已使用预算
	var usedBudget float64
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "expense", startDate, endDate).Select("COALESCE(SUM(amount), 0)").Scan(&usedBudget)

	// 计算剩余预算和使用率
	remainingBudget := totalBudget - usedBudget
	usageRate := 0.0
	if totalBudget > 0 {
		usageRate = (usedBudget / totalBudget) * 100
	}

	// 查询分类预算执行情况
	type CategoryBudget struct {
		CategoryID   uint    `json:"category_id"`
		CategoryName string  `json:"category_name"`
		Budget       float64 `json:"budget"`
		Used         float64 `json:"used"`
		UsageRate    float64 `json:"usage_rate"`
	}
	var categories []CategoryBudget

	// 先查询所有分类预算
	var categoryBudgets []struct {
		CategoryID uint    `json:"category_id"`
		Amount     float64 `json:"amount"`
	}
	database.DB.Model(&models.Budget{}).Where("book_id = ? AND period = ? AND date = ? AND category_id IS NOT NULL", bookID, period, dateStr).Select("category_id, amount").Scan(&categoryBudgets)

	// 构建分类ID映射
	categoryBudgetMap := make(map[uint]float64)
	for _, cb := range categoryBudgets {
		categoryBudgetMap[cb.CategoryID] = cb.Amount
	}

	// 查询分类支出
	var categoryExpenses []struct {
		CategoryID uint    `json:"category_id"`
		Name       string  `json:"name"`
		Amount     float64 `json:"amount"`
	}
	database.DB.Table("transactions").Select("categories.category_id, categories.name, COALESCE(SUM(transactions.amount), 0) as amount").
		Joins("LEFT JOIN categories ON transactions.category_id = categories.category_id").
		Where("transactions.book_id = ? AND transactions.type = ? AND transactions.date BETWEEN ? AND ?", bookID, "expense", startDate, endDate).
		Group("categories.category_id, categories.name").
		Scan(&categoryExpenses)

	// 构建分类预算执行情况
	for _, ce := range categoryExpenses {
		budget := categoryBudgetMap[ce.CategoryID]
		usageRate := 0.0
		if budget > 0 {
			usageRate = (ce.Amount / budget) * 100
		}

		categories = append(categories, CategoryBudget{
			CategoryID:   ce.CategoryID,
			CategoryName: ce.Name,
			Budget:       budget,
			Used:         ce.Amount,
			UsageRate:    usageRate,
		})
	}

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取预算执行进度成功",
		"data": gin.H{
			"total_budget":     totalBudget,
			"used_budget":      usedBudget,
			"remaining_budget": remainingBudget,
			"usage_rate":       usageRate,
			"categories":       categories,
		},
	})
}

// 辅助函数：根据周期获取时间范围
func getPeriodRange(date time.Time, period string) (startDate, endDate time.Time) {
	year, month, day := date.Date()

	switch period {
	case "day":
		startDate = time.Date(year, month, day, 0, 0, 0, 0, date.Location())
		endDate = startDate.Add(24 * time.Hour).Add(-time.Second)
	case "week":
		// 获取本周一
		weekday := int(date.Weekday())
		if weekday == 0 {
			weekday = 7
		}
		startDate = time.Date(year, month, day-weekday+1, 0, 0, 0, 0, date.Location())
		endDate = startDate.Add(7 * 24 * time.Hour).Add(-time.Second)
	case "month":
		startDate = time.Date(year, month, 1, 0, 0, 0, 0, date.Location())
		endDate = startDate.AddDate(0, 1, 0).Add(-time.Second)
	case "year":
		startDate = time.Date(year, 1, 1, 0, 0, 0, 0, date.Location())
		endDate = startDate.AddDate(1, 0, 0).Add(-time.Second)
	default:
		startDate = time.Date(year, month, 1, 0, 0, 0, 0, date.Location())
		endDate = startDate.AddDate(0, 1, 0).Add(-time.Second)
	}

	return
}

// 辅助函数：获取上一周期的时间范围
func getPreviousPeriodRange(date time.Time, period string) (startDate, endDate time.Time) {
	year, month, day := date.Date()

	switch period {
	case "day":
		startDate = time.Date(year, month, day-1, 0, 0, 0, 0, date.Location())
		endDate = startDate.Add(24 * time.Hour).Add(-time.Second)
	case "week":
		// 获取上周一
		weekday := int(date.Weekday())
		if weekday == 0 {
			weekday = 7
		}
		startDate = time.Date(year, month, day-weekday+1-7, 0, 0, 0, 0, date.Location())
		endDate = startDate.Add(7 * 24 * time.Hour).Add(-time.Second)
	case "month":
		startDate = time.Date(year, month-1, 1, 0, 0, 0, 0, date.Location())
		endDate = startDate.AddDate(0, 1, 0).Add(-time.Second)
	case "year":
		startDate = time.Date(year-1, 1, 1, 0, 0, 0, 0, date.Location())
		endDate = startDate.AddDate(1, 0, 0).Add(-time.Second)
	default:
		startDate = time.Date(year, month-1, 1, 0, 0, 0, 0, date.Location())
		endDate = startDate.AddDate(0, 1, 0).Add(-time.Second)
	}

	return
}

// 辅助函数：计算增长率
func calculateGrowthRate(current, previous float64) float64 {
	if previous == 0 {
		return 0
	}
	return ((current - previous) / previous) * 100
}
