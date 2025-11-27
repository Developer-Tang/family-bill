package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLineChartData 获取折线图数据
func GetLineChartData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取折线图数据成功",
		"data": gin.H{
			"title":  "2023年收支趋势",
			"x_axis": []string{"2023-01", "2023-02", "2023-03", "2023-04", "2023-05", "2023-06", "2023-07", "2023-08", "2023-09", "2023-10", "2023-11", "2023-12"},
			"series": []gin.H{
				{
					"name": "收入",
					"data": []float64{8000.00, 8200.00, 8500.00, 8300.00, 8600.00, 8400.00, 8700.00, 8500.00, 8800.00, 8600.00, 8900.00, 9000.00},
				},
				{
					"name": "支出",
					"data": []float64{3500.00, 3800.00, 3600.00, 3700.00, 3900.00, 3800.00, 4000.00, 3900.00, 4100.00, 4000.00, 4200.00, 4300.00},
				},
			},
		},
	})
}

// GetPieChartData 获取饼图数据
func GetPieChartData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取饼图数据成功",
		"data": gin.H{
			"title": "2023年1月支出分类占比",
			"series": []gin.H{
				{
					"name": "支出分类",
					"data": []gin.H{
						{"name": "餐饮", "value": 1200.00},
						{"name": "住房", "value": 1000.00},
						{"name": "交通", "value": 500.00},
						{"name": "购物", "value": 400.00},
						{"name": "娱乐", "value": 300.00},
						{"name": "其他", "value": 100.00},
					},
				},
			},
		},
	})
}

// GetBarChartData 获取柱状图数据
func GetBarChartData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取柱状图数据成功",
		"data": gin.H{
			"title":  "2023年1-3月支出分类对比",
			"x_axis": []string{"餐饮", "住房", "交通", "购物", "娱乐"},
			"series": []gin.H{
				{
					"name": "1月",
					"data": []float64{1200.00, 1000.00, 500.00, 400.00, 300.00},
				},
				{
					"name": "2月",
					"data": []float64{1300.00, 1000.00, 550.00, 450.00, 350.00},
				},
				{
					"name": "3月",
					"data": []float64{1250.00, 1000.00, 520.00, 420.00, 330.00},
				},
			},
		},
	})
}

// GetRadarChartData 获取雷达图数据
func GetRadarChartData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取雷达图数据成功",
		"data": gin.H{
			"title": "2023年Q1-Q2支出分类对比",
			"indicator": []gin.H{
				{"name": "餐饮", "max": 2000},
				{"name": "住房", "max": 1500},
				{"name": "交通", "max": 1000},
				{"name": "购物", "max": 1000},
				{"name": "娱乐", "max": 800},
				{"name": "其他", "max": 500},
			},
			"series": []gin.H{
				{
					"name": "Q1",
					"data": []float64{1800.00, 1500.00, 750.00, 600.00, 450.00, 150.00},
				},
				{
					"name": "Q2",
					"data": []float64{1950.00, 1500.00, 825.00, 675.00, 525.00, 165.00},
				},
			},
		},
	})
}
