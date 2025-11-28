package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
)

// GetIncomeExpenseAnalysis 获取收支对比分析
// @Summary 获取收支对比分析
// @Description 获取指定账本在特定时间范围内的收支对比数据，支持多维度对比分析
// @Tags 收支分析
// @Accept jsonZZ
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param period query string false "统计周期：day, week, month, year" default(month)
// @Param start_date query string true "开始日期，格式：YYYY-MM-DD或YYYY-MM"
// @Param end_date query string true "结束日期，格式：YYYY-MM-DD或YYYY-MM"
// @Param compare_period query string false "对比周期：previous（上一周期）, same（同比）" default(previous)
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/analysis/income-expense [get]
func GetIncomeExpenseAnalysis(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	// period := c.DefaultQuery("period", "month") // 未使用的参数
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	comparePeriod := c.DefaultQuery("compare_period", "previous")

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

	// 获取当前周期的收入和支出
	var currentIncome, currentExpense float64
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "income", startDate, endDate).Select("COALESCE(SUM(amount), 0)").Scan(&currentIncome)
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "expense", startDate, endDate).Select("COALESCE(SUM(amount), 0)").Scan(&currentExpense)

	// 获取对比周期的开始和结束时间
	var compareStartDate, compareEndDate time.Time
	switch comparePeriod {
	case "previous":
		// 计算上一周期的时间范围
		duration := endDate.Sub(startDate)
		compareStartDate = startDate.Add(-duration)
		compareEndDate = startDate.Add(-time.Second)
	case "same":
		// 计算同比周期的时间范围（去年同期）
		compareStartDate = startDate.AddDate(-1, 0, 0)
		compareEndDate = endDate.AddDate(-1, 0, 0)
	default:
		// 默认使用上一周期
		duration := endDate.Sub(startDate)
		compareStartDate = startDate.Add(-duration)
		compareEndDate = startDate.Add(-time.Second)
	}

	// 获取对比周期的收入和支出
	var compareIncome, compareExpense float64
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "income", compareStartDate, compareEndDate).Select("COALESCE(SUM(amount), 0)").Scan(&compareIncome)
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "expense", compareStartDate, compareEndDate).Select("COALESCE(SUM(amount), 0)").Scan(&compareExpense)

	// 计算变化率
	incomeRate := calculateGrowthRate(currentIncome, compareIncome)
	expenseRate := calculateGrowthRate(currentExpense, compareExpense)
	balanceRate := calculateGrowthRate(currentIncome-currentExpense, compareIncome-compareExpense)

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取收支对比分析成功",
		"data": gin.H{
			"current": gin.H{
				"income":  currentIncome,
				"expense": currentExpense,
				"balance": currentIncome - currentExpense,
			},
			"compare": gin.H{
				"income":  compareIncome,
				"expense": compareExpense,
				"balance": compareIncome - compareExpense,
			},
			"change": gin.H{
				"income_rate":  incomeRate,
				"expense_rate": expenseRate,
				"balance_rate": balanceRate,
			},
		},
	})
}

// GetTrendAnalysis 获取收支趋势分析
// @Summary 获取收支趋势分析
// @Description 获取指定账本在特定时间范围内的收支趋势数据，支持按日、周、月、年统计
// @Tags 收支分析
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param type query string false "数据类型：income（收入）, expense（支出）, both（收支）" default(both)
// @Param period query string false "统计周期：day, week, month, year" default(month)
// @Param start_date query string true "开始日期，格式：YYYY-MM-DD或YYYY-MM"
// @Param end_date query string true "结束日期，格式：YYYY-MM-DD或YYYY-MM"
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/analysis/trend [get]
func GetTrendAnalysis(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	dataType := c.DefaultQuery("type", "both")
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

	// 构建查询条件
	query := database.DB.Model(&models.Transaction{}).Where("book_id = ?", bookID)

	// 添加数据类型筛选
	query = query.Where("type = ?", dataType)

	// 查询趋势数据
	type TrendData struct {
		Date    string  `json:"date"`
		Income  float64 `json:"income"`
		Expense float64 `json:"expense"`
		Balance float64 `json:"balance"`
	}
	var trendData []TrendData

	// 遍历时间序列，查询每个时间段的数据
	for _, dateRange := range timeSeries {
		var income, expense float64

		// 查询收入
		if dataType == "both" || dataType == "income" {
			database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "income", dateRange.Start, dateRange.End).Select("COALESCE(SUM(amount), 0)").Scan(&income)
		}

		// 查询支出
		if dataType == "both" || dataType == "expense" {
			database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "expense", dateRange.Start, dateRange.End).Select("COALESCE(SUM(amount), 0)").Scan(&expense)
		}

		// 计算余额
		balance := income - expense

		// 添加到结果集
		trendData = append(trendData, TrendData{
			Date:    dateRange.Label,
			Income:  income,
			Expense: expense,
			Balance: balance,
		})
	}

	// 计算汇总数据
	var totalIncome, totalExpense, totalBalance float64
	for _, data := range trendData {
		totalIncome += data.Income
		totalExpense += data.Expense
		totalBalance += data.Balance
	}

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取收支趋势分析成功",
		"data": gin.H{
			"period":     period,
			"start_date": startDateStr,
			"end_date":   endDateStr,
			"data":       trendData,
			"summary": gin.H{
				"total_income":  totalIncome,
				"total_expense": totalExpense,
				"total_balance": totalBalance,
			},
		},
	})
}

// GetFlowAnalysis 获取资金流向分析
// @Summary 获取资金流向分析
// @Description 获取指定账本在特定时间范围内的资金流向分析数据，展示资金的来源和去向
// @Tags 收支分析
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param period query string false "统计周期：day, week, month, year" default(month)
// @Param start_date query string true "开始日期，格式：YYYY-MM-DD或YYYY-MM"
// @Param end_date query string true "结束日期，格式：YYYY-MM-DD或YYYY-MM"
// @Param min_amount query number false "最小金额过滤" default(0)
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/analysis/flow [get]
func GetFlowAnalysis(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	// period := c.DefaultQuery("period", "month") // 未使用的参数
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	minAmountStr := c.DefaultQuery("min_amount", "0")

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

	// 解析最小金额
	minAmount, err := strconv.ParseFloat(minAmountStr, 64)
	if err != nil {
		minAmount = 0
	}

	// 查询收入流向
	type IncomeFlow struct {
		CategoryID   uint    `json:"category_id"`
		CategoryName string  `json:"category_name"`
		Amount       float64 `json:"amount"`
		Percentage   float64 `json:"percentage"`
	}
	var incomeFlow []IncomeFlow

	// 查询总收入
	var totalIncome float64
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ? AND amount >= ?", bookID, "income", startDate, endDate, minAmount).Select("COALESCE(SUM(amount), 0)").Scan(&totalIncome)

	// 查询收入分类统计
	var incomeCategories []struct {
		CategoryID uint    `json:"category_id"`
		Name       string  `json:"name"`
		Amount     float64 `json:"amount"`
	}
	database.DB.Table("transactions").Select("categories.category_id, categories.name, COALESCE(SUM(transactions.amount), 0) as amount").
		Joins("LEFT JOIN categories ON transactions.category_id = categories.category_id").
		Where("transactions.book_id = ? AND transactions.type = ? AND transactions.date BETWEEN ? AND ? AND transactions.amount >= ?", bookID, "income", startDate, endDate, minAmount).
		Group("categories.category_id, categories.name").
		Order("amount DESC").
		Scan(&incomeCategories)

	// 计算收入分类占比
	for _, ic := range incomeCategories {
		percentage := 0.0
		if totalIncome > 0 {
			percentage = (ic.Amount / totalIncome) * 100
		}
		incomeFlow = append(incomeFlow, IncomeFlow{
			CategoryID:   ic.CategoryID,
			CategoryName: ic.Name,
			Amount:       ic.Amount,
			Percentage:   percentage,
		})
	}

	// 查询支出流向
	type ExpenseFlow struct {
		CategoryID   uint    `json:"category_id"`
		CategoryName string  `json:"category_name"`
		Amount       float64 `json:"amount"`
		Percentage   float64 `json:"percentage"`
	}
	var expenseFlow []ExpenseFlow

	// 查询总支出
	var totalExpense float64
	database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ? AND amount >= ?", bookID, "expense", startDate, endDate, minAmount).Select("COALESCE(SUM(amount), 0)").Scan(&totalExpense)

	// 查询支出分类统计
	var expenseCategories []struct {
		CategoryID uint    `json:"category_id"`
		Name       string  `json:"name"`
		Amount     float64 `json:"amount"`
	}
	database.DB.Table("transactions").Select("categories.category_id, categories.name, COALESCE(SUM(transactions.amount), 0) as amount").
		Joins("LEFT JOIN categories ON transactions.category_id = categories.category_id").
		Where("transactions.book_id = ? AND transactions.type = ? AND transactions.date BETWEEN ? AND ? AND transactions.amount >= ?", bookID, "expense", startDate, endDate, minAmount).
		Group("categories.category_id, categories.name").
		Order("amount DESC").
		Scan(&expenseCategories)

	// 计算支出分类占比
	for _, ec := range expenseCategories {
		percentage := 0.0
		if totalExpense > 0 {
			percentage = (ec.Amount / totalExpense) * 100
		}
		expenseFlow = append(expenseFlow, ExpenseFlow{
			CategoryID:   ec.CategoryID,
			CategoryName: ec.Name,
			Amount:       ec.Amount,
			Percentage:   percentage,
		})
	}

	// 查询转账流向
	type TransferFlow struct {
		FromAccountID   uint    `json:"from_account_id"`
		FromAccountName string  `json:"from_account_name"`
		ToAccountID     uint    `json:"to_account_id"`
		ToAccountName   string  `json:"to_account_name"`
		Amount          float64 `json:"amount"`
	}
	var transferFlow []TransferFlow

	// 查询转账记录
	var transfers []struct {
		FromAccountID uint    `json:"from_account_id"`
		FromName      string  `json:"from_name"`
		ToAccountID   uint    `json:"to_account_id"`
		ToName        string  `json:"to_name"`
		Amount        float64 `json:"amount"`
	}
	database.DB.Table("transactions").Select("from_accounts.account_id as from_account_id, from_accounts.name as from_name, to_accounts.account_id as to_account_id, to_accounts.name as to_name, transactions.amount").
		Joins("LEFT JOIN accounts as from_accounts ON transactions.from_account_id = from_accounts.account_id").
		Joins("LEFT JOIN accounts as to_accounts ON transactions.to_account_id = to_accounts.account_id").
		Where("transactions.book_id = ? AND transactions.type = ? AND transactions.date BETWEEN ? AND ? AND transactions.amount >= ?", bookID, "transfer", startDate, endDate, minAmount).
		Scan(&transfers)

	// 构建转账流向数据
	for _, t := range transfers {
		transferFlow = append(transferFlow, TransferFlow{
			FromAccountID:   t.FromAccountID,
			FromAccountName: t.FromName,
			ToAccountID:     t.ToAccountID,
			ToAccountName:   t.ToName,
			Amount:          t.Amount,
		})
	}

	// 构建响应数据
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取资金流向分析成功",
		"data": gin.H{
			"income_flow":   incomeFlow,
			"expense_flow":  expenseFlow,
			"transfer_flow": transferFlow,
		},
	})
}

// 辅助函数：生成时间序列
func generateTimeSeries(startDate, endDate time.Time, period string) []struct {
	Start time.Time
	End   time.Time
	Label string
} {
	var timeSeries []struct {
		Start time.Time
		End   time.Time
		Label string
	}

	switch period {
	case "day":
		// 按天生成时间序列
		for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
			start := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
			end := start.Add(24 * time.Hour).Add(-time.Second)
			timeSeries = append(timeSeries, struct {
				Start time.Time
				End   time.Time
				Label string
			}{Start: start, End: end, Label: start.Format("2006-01-02")})
		}

	case "week":
		// 按周生成时间序列
		current := startDate
		for !current.After(endDate) {
			// 获取本周一
			weekday := int(current.Weekday())
			if weekday == 0 {
				weekday = 7
			}
			weekStart := time.Date(current.Year(), current.Month(), current.Day()-weekday+1, 0, 0, 0, 0, current.Location())
			weekEnd := weekStart.Add(7 * 24 * time.Hour).Add(-time.Second)

			// 确保不超过结束日期
			if weekEnd.After(endDate) {
				weekEnd = endDate
			}

			timeSeries = append(timeSeries, struct {
				Start time.Time
				End   time.Time
				Label string
			}{Start: weekStart, End: weekEnd, Label: weekStart.Format("2006-01-02") + "~" + weekEnd.Format("2006-01-02")})

			// 移动到下一周
			current = weekEnd.AddDate(0, 0, 1)
		}

	case "month":
		// 按月生成时间序列
		current := startDate
		for !current.After(endDate) {
			monthStart := time.Date(current.Year(), current.Month(), 1, 0, 0, 0, 0, current.Location())
			monthEnd := monthStart.AddDate(0, 1, 0).Add(-time.Second)

			// 确保不超过结束日期
			if monthEnd.After(endDate) {
				monthEnd = endDate
			}

			timeSeries = append(timeSeries, struct {
				Start time.Time
				End   time.Time
				Label string
			}{Start: monthStart, End: monthEnd, Label: monthStart.Format("2006-01")})

			// 移动到下一月
			current = monthEnd.AddDate(0, 1, 0)
		}

	case "year":
		// 按年生成时间序列
		current := startDate
		for !current.After(endDate) {
			yearStart := time.Date(current.Year(), 1, 1, 0, 0, 0, 0, current.Location())
			yearEnd := yearStart.AddDate(1, 0, 0).Add(-time.Second)

			// 确保不超过结束日期
			if yearEnd.After(endDate) {
				yearEnd = endDate
			}

			timeSeries = append(timeSeries, struct {
				Start time.Time
				End   time.Time
				Label string
			}{Start: yearStart, End: yearEnd, Label: yearStart.Format("2006")})

			// 移动到下一年
			current = yearEnd.AddDate(1, 0, 0)
		}

	default:
		// 默认按月生成时间序列
		current := startDate
		for !current.After(endDate) {
			monthStart := time.Date(current.Year(), current.Month(), 1, 0, 0, 0, 0, current.Location())
			monthEnd := monthStart.AddDate(0, 1, 0).Add(-time.Second)

			// 确保不超过结束日期
			if monthEnd.After(endDate) {
				monthEnd = endDate
			}

			timeSeries = append(timeSeries, struct {
				Start time.Time
				End   time.Time
				Label string
			}{Start: monthStart, End: monthEnd, Label: monthStart.Format("2006-01")})

			// 移动到下一月
			current = monthEnd.AddDate(0, 1, 0)
		}
	}

	return timeSeries
}
