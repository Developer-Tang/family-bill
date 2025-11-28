package v1

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
)

// GetLineChartData 获取折线图数据
// @Summary 获取折线图数据
// @Description 获取指定账本的折线图数据，支持多种数据维度和时间范围
// @Tags 图表数据
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param type query string true "图表类型：income-expense（收支对比）, category-trend（分类趋势）, account-trend（账户趋势）"
// @Param period query string false "统计周期：day, week, month, year" default(month)
// @Param start_date query string true "开始日期，格式：YYYY-MM-DD或YYYY-MM"
// @Param end_date query string true "结束日期，格式：YYYY-MM-DD或YYYY-MM"
// @Param category_ids query string false "分类ID列表，多个用逗号分隔"
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/charts/line [get]
func GetLineChartData(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	chartType := c.Query("type")
	period := c.DefaultQuery("period", "month")
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	categoryIDsStr := c.Query("category_ids")

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

	// 解析分类ID列表
	var categoryIDs []uint
	if categoryIDsStr != "" {
		idStrs := strings.Split(categoryIDsStr, ",")
		for _, idStr := range idStrs {
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err == nil {
				categoryIDs = append(categoryIDs, uint(id))
			}
		}
	}

	// 构建响应数据
	response := gin.H{
		"code":    200,
		"message": "获取折线图数据成功",
		"data": gin.H{
			"title":  "",
			"x_axis": []string{},
			"series": []gin.H{},
		},
	}

	// 根据图表类型生成数据
	switch chartType {
	case "income-expense":
		// 收支对比折线图
		response["data"].(gin.H)["title"] = startDate.Format("2006") + "年收支趋势"

		// 生成X轴数据
		var xAxis []string
		for _, ts := range timeSeries {
			xAxis = append(xAxis, ts.Label)
		}
		response["data"].(gin.H)["x_axis"] = xAxis

		// 生成收入和支出数据
		var incomeData, expenseData []float64
		for _, ts := range timeSeries {
			var income, expense float64
			database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "income", ts.Start, ts.End).Select("COALESCE(SUM(amount), 0)").Scan(&income)
			database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "expense", ts.Start, ts.End).Select("COALESCE(SUM(amount), 0)").Scan(&expense)
			incomeData = append(incomeData, income)
			expenseData = append(expenseData, expense)
		}

		// 构建系列数据
		series := []gin.H{
			{
				"name": "收入",
				"data": incomeData,
			},
			{
				"name": "支出",
				"data": expenseData,
			},
		}
		response["data"].(gin.H)["series"] = series

	case "category-trend":
		// 分类趋势折线图
		response["data"].(gin.H)["title"] = startDate.Format("2006") + "年分类趋势"

		// 生成X轴数据
		var xAxis []string
		for _, ts := range timeSeries {
			xAxis = append(xAxis, ts.Label)
		}
		response["data"].(gin.H)["x_axis"] = xAxis

		// 查询分类列表
		var categories []models.Category
		query := database.DB.Model(&models.Category{})
		if len(categoryIDs) > 0 {
			query = query.Where("category_id IN ?", categoryIDs)
		}
		query.Find(&categories)

		// 生成系列数据
		var series []gin.H
		for _, category := range categories {
			var data []float64
			for _, ts := range timeSeries {
				var amount float64
				database.DB.Model(&models.Transaction{}).Where("book_id = ? AND category_id = ? AND date BETWEEN ? AND ?", bookID, category.CategoryID, ts.Start, ts.End).Select("COALESCE(SUM(amount), 0)").Scan(&amount)
				data = append(data, amount)
			}
			series = append(series, gin.H{
				"name": category.Name,
				"data": data,
			})
		}
		response["data"].(gin.H)["series"] = series

	case "account-trend":
		// 账户趋势折线图
		response["data"].(gin.H)["title"] = startDate.Format("2006") + "年账户趋势"

		// 生成X轴数据
		var xAxis []string
		for _, ts := range timeSeries {
			xAxis = append(xAxis, ts.Label)
		}
		response["data"].(gin.H)["x_axis"] = xAxis

		// 查询账户列表
		var accounts []models.Account
		database.DB.Model(&models.Account{}).Where("book_id = ?", bookID).Find(&accounts)

		// 生成系列数据
		var series []gin.H
		for _, account := range accounts {
			var data []float64
			for _, ts := range timeSeries {
				// 查询账户在该时间段的余额变化
				var income, expense float64
				database.DB.Model(&models.Transaction{}).Where("book_id = ? AND account_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, account.AccountID, "income", ts.Start, ts.End).Select("COALESCE(SUM(amount), 0)").Scan(&income)
				database.DB.Model(&models.Transaction{}).Where("book_id = ? AND account_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, account.AccountID, "expense", ts.Start, ts.End).Select("COALESCE(SUM(amount), 0)").Scan(&expense)
				balanceChange := income - expense
				data = append(data, balanceChange)
			}
			series = append(series, gin.H{
				"name": account.Name,
				"data": data,
			})
		}
		response["data"].(gin.H)["series"] = series

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的图表类型",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetPieChartData 获取饼图数据
// @Summary 获取饼图数据
// @Description 获取指定账本的饼图数据，支持多种数据维度
// @Tags 图表数据
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param type query string true "图表类型：category（分类占比）, account（账户占比）, member（成员占比）"
// @Param period query string false "统计周期：day, week, month, year" default(month)
// @Param date query string true "统计日期，格式：YYYY-MM-DD或YYYY-MM"
// @Param data_type query string false "数据类型：income（收入）, expense（支出）" default(expense)
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/charts/pie [get]
func GetPieChartData(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	chartType := c.Query("type")
	period := c.DefaultQuery("period", "month")
	dateStr := c.Query("date")
	dataType := c.DefaultQuery("data_type", "expense")

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

	// 构建响应数据
	response := gin.H{
		"code":    200,
		"message": "获取饼图数据成功",
		"data": gin.H{
			"title":  "",
			"series": []gin.H{},
		},
	}

	// 根据图表类型生成数据
	switch chartType {
	case "category":
		// 分类占比饼图
		response["data"].(gin.H)["title"] = date.Format("2006年01月") + dataType + "分类占比"

		// 查询分类统计
		var categories []struct {
			Name   string  `json:"name"`
			Amount float64 `json:"value"`
		}
		database.DB.Table("transactions").Select("categories.name, COALESCE(SUM(transactions.amount), 0) as amount").
			Joins("LEFT JOIN categories ON transactions.category_id = categories.category_id").
			Where("transactions.book_id = ? AND transactions.type = ? AND transactions.date BETWEEN ? AND ?", bookID, dataType, startDate, endDate).
			Group("categories.name").
			Order("amount DESC").
			Scan(&categories)

		// 构建系列数据
		series := []gin.H{
			{
				"name": dataType + "分类",
				"data": categories,
			},
		}
		response["data"].(gin.H)["series"] = series

	case "account":
		// 账户占比饼图
		response["data"].(gin.H)["title"] = date.Format("2006年01月") + dataType + "账户占比"

		// 查询账户统计
		var accounts []struct {
			Name   string  `json:"name"`
			Amount float64 `json:"value"`
		}
		database.DB.Table("transactions").Select("accounts.name, COALESCE(SUM(transactions.amount), 0) as amount").
			Joins("LEFT JOIN accounts ON transactions.account_id = accounts.account_id").
			Where("transactions.book_id = ? AND transactions.type = ? AND transactions.date BETWEEN ? AND ?", bookID, dataType, startDate, endDate).
			Group("accounts.name").
			Order("amount DESC").
			Scan(&accounts)

		// 构建系列数据
		series := []gin.H{
			{
				"name": dataType + "账户",
				"data": accounts,
			},
		}
		response["data"].(gin.H)["series"] = series

	case "member":
		// 成员占比饼图
		response["data"].(gin.H)["title"] = date.Format("2006年01月") + dataType + "成员占比"

		// 查询成员统计
		var members []struct {
			Name   string  `json:"name"`
			Amount float64 `json:"value"`
		}
		database.DB.Table("transactions").Select("users.username as name, COALESCE(SUM(transactions.amount), 0) as amount").
			Joins("LEFT JOIN users ON transactions.user_id = users.user_id").
			Where("transactions.book_id = ? AND transactions.type = ? AND transactions.date BETWEEN ? AND ?", bookID, dataType, startDate, endDate).
			Group("users.username").
			Order("amount DESC").
			Scan(&members)

		// 构建系列数据
		series := []gin.H{
			{
				"name": dataType + "成员",
				"data": members,
			},
		}
		response["data"].(gin.H)["series"] = series

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的图表类型",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetBarChartData 获取柱状图数据
// @Summary 获取柱状图数据
// @Description 获取指定账本的柱状图数据，支持多种数据维度和时间范围
// @Tags 图表数据
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param type query string true "图表类型：income-expense（收支对比）, category（分类对比）, account（账户对比）"
// @Param period query string false "统计周期：day, week, month, year" default(month)
// @Param start_date query string true "开始日期，格式：YYYY-MM-DD或YYYY-MM"
// @Param end_date query string true "结束日期，格式：YYYY-MM-DD或YYYY-MM"
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/charts/bar [get]
func GetBarChartData(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	chartType := c.Query("type")
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

	// 构建响应数据
	response := gin.H{
		"code":    200,
		"message": "获取柱状图数据成功",
		"data": gin.H{
			"title":  "",
			"x_axis": []string{},
			"series": []gin.H{},
		},
	}

	// 根据图表类型生成数据
	switch chartType {
	case "income-expense":
		// 收支对比柱状图
		response["data"].(gin.H)["title"] = startDate.Format("2006年") + "收支对比"

		// 生成X轴数据
		var xAxis []string
		for _, ts := range timeSeries {
			xAxis = append(xAxis, ts.Label)
		}
		response["data"].(gin.H)["x_axis"] = xAxis

		// 生成收入和支出数据
		var incomeData, expenseData []float64
		for _, ts := range timeSeries {
			var income, expense float64
			database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "income", ts.Start, ts.End).Select("COALESCE(SUM(amount), 0)").Scan(&income)
			database.DB.Model(&models.Transaction{}).Where("book_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, "expense", ts.Start, ts.End).Select("COALESCE(SUM(amount), 0)").Scan(&expense)
			incomeData = append(incomeData, income)
			expenseData = append(expenseData, expense)
		}

		// 构建系列数据
		series := []gin.H{
			{
				"name": "收入",
				"data": incomeData,
			},
			{
				"name": "支出",
				"data": expenseData,
			},
		}
		response["data"].(gin.H)["series"] = series

	case "category":
		// 分类对比柱状图
		response["data"].(gin.H)["title"] = startDate.Format("2006年") + "分类对比"

		// 查询分类列表
		var categories []models.Category
		database.DB.Model(&models.Category{}).Where("type = ?", "expense").Find(&categories)

		// 生成X轴数据
		var xAxis []string
		for _, category := range categories {
			xAxis = append(xAxis, category.Name)
		}
		response["data"].(gin.H)["x_axis"] = xAxis

		// 生成系列数据
		var series []gin.H
		for _, ts := range timeSeries {
			var data []float64
			for _, category := range categories {
				var amount float64
				database.DB.Model(&models.Transaction{}).Where("book_id = ? AND category_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, category.CategoryID, "expense", ts.Start, ts.End).Select("COALESCE(SUM(amount), 0)").Scan(&amount)
				data = append(data, amount)
			}
			series = append(series, gin.H{
				"name": ts.Label,
				"data": data,
			})
		}
		response["data"].(gin.H)["series"] = series

	case "account":
		// 账户对比柱状图
		response["data"].(gin.H)["title"] = startDate.Format("2006年") + "账户对比"

		// 查询账户列表
		var accounts []models.Account
		database.DB.Model(&models.Account{}).Where("book_id = ?", bookID).Find(&accounts)

		// 生成X轴数据
		var xAxis []string
		for _, account := range accounts {
			xAxis = append(xAxis, account.Name)
		}
		response["data"].(gin.H)["x_axis"] = xAxis

		// 生成系列数据
		var series []gin.H
		for _, ts := range timeSeries {
			var data []float64
			for _, account := range accounts {
				var amount float64
				database.DB.Model(&models.Transaction{}).Where("book_id = ? AND account_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, account.AccountID, "expense", ts.Start, ts.End).Select("COALESCE(SUM(amount), 0)").Scan(&amount)
				data = append(data, amount)
			}
			series = append(series, gin.H{
				"name": ts.Label,
				"data": data,
			})
		}
		response["data"].(gin.H)["series"] = series

	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的图表类型",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetRadarChartData 获取雷达图数据
// @Summary 获取雷达图数据
// @Description 获取指定账本的雷达图数据，支持多维度对比分析
// @Tags 图表数据
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param book_id query int true "账本ID"
// @Param period query string false "统计周期：month, quarter, year" default(month)
// @Param start_date query string true "开始日期，格式：YYYY-MM"
// @Param compare_period query string false "对比周期：previous（上一周期）, same（同比）" default(previous)
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/charts/radar [get]
func GetRadarChartData(c *gin.Context) {
	// 获取查询参数
	bookIDStr := c.Query("book_id")
	period := c.DefaultQuery("period", "month")
	startDateStr := c.Query("start_date")
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
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的开始日期格式",
		})
		return
	}

	// 获取当前周期的开始和结束时间
	currentStartDate, currentEndDate := getPeriodRange(startDate, period)

	// 获取对比周期的开始和结束时间
	var compareStartDate, compareEndDate time.Time
	switch comparePeriod {
	case "previous":
		// 计算上一周期的时间范围
		duration := currentEndDate.Sub(currentStartDate)
		compareStartDate = currentStartDate.Add(-duration)
		compareEndDate = currentStartDate.Add(-time.Second)
	case "same":
		// 计算同比周期的时间范围（去年同期）
		compareStartDate = currentStartDate.AddDate(-1, 0, 0)
		compareEndDate = currentEndDate.AddDate(-1, 0, 0)
	default:
		// 默认使用上一周期
		duration := currentEndDate.Sub(currentStartDate)
		compareStartDate = currentStartDate.Add(-duration)
		compareEndDate = currentStartDate.Add(-time.Second)
	}

	// 查询分类列表
	var categories []models.Category
	database.DB.Model(&models.Category{}).Where("type = ?", "expense").Find(&categories)

	// 构建响应数据
	response := gin.H{
		"code":    200,
		"message": "获取雷达图数据成功",
		"data": gin.H{
			"title":     startDate.Format("2006年") + "支出分类对比",
			"indicator": []gin.H{},
			"series":    []gin.H{},
		},
	}

	// 生成指标数据
	var indicators []gin.H
	for _, category := range categories {
		// 查询该分类的最大支出金额（用于设置雷达图最大值）
		var maxAmount float64
		database.DB.Model(&models.Transaction{}).Where("book_id = ? AND category_id = ? AND type = ?", bookID, category.CategoryID, "expense").Select("COALESCE(MAX(amount), 1000)").Scan(&maxAmount)

		// 设置最大值为当前最大值的1.2倍，留出空间
		maxAmount = maxAmount * 1.2

		indicators = append(indicators, gin.H{
			"name": category.Name,
			"max":  maxAmount,
		})
	}
	response["data"].(gin.H)["indicator"] = indicators

	// 生成当前周期数据
	var currentData []float64
	for _, category := range categories {
		var amount float64
		database.DB.Model(&models.Transaction{}).Where("book_id = ? AND category_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, category.CategoryID, "expense", currentStartDate, currentEndDate).Select("COALESCE(SUM(amount), 0)").Scan(&amount)
		currentData = append(currentData, amount)
	}

	// 生成对比周期数据
	var compareData []float64
	for _, category := range categories {
		var amount float64
		database.DB.Model(&models.Transaction{}).Where("book_id = ? AND category_id = ? AND type = ? AND date BETWEEN ? AND ?", bookID, category.CategoryID, "expense", compareStartDate, compareEndDate).Select("COALESCE(SUM(amount), 0)").Scan(&amount)
		compareData = append(compareData, amount)
	}

	// 构建系列数据
	series := []gin.H{
		{
			"name": startDate.Format("2006年01月"),
			"data": currentData,
		},
		{
			"name": compareStartDate.Format("2006年01月"),
			"data": compareData,
		},
	}
	response["data"].(gin.H)["series"] = series

	c.JSON(http.StatusOK, response)
}
