package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
)

// GetOperationLogs 获取操作审计日志
// @Summary 获取操作审计日志
// @Description 获取用户的操作审计日志列表，支持分页和筛选
// @Tags 系统安全
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param page query int false "页码" default(1)
// @Param size query int false "每页数量" default(10)
// @Param start_date query string false "开始日期，格式：YYYY-MM-DD"
// @Param end_date query string false "结束日期，格式：YYYY-MM-DD"
// @Param action query string false "操作类型"
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/logs/operations [get]
func GetOperationLogs(c *gin.Context) {
	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	offset := (page - 1) * pageSize

	// 获取筛选参数
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	action := c.Query("action")

	// 构建查询
	query := database.DB.Model(&models.OperationLog{}).Where("user_id = ?", userID)

	// 添加日期筛选
	if startDate != "" {
		query = query.Where("created_at >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("created_at <= ?", endDate)
	}

	// 添加操作类型筛选
	if action != "" {
		query = query.Where("action = ?", action)
	}

	// 查询总数
	var total int64
	query.Count(&total)

	// 查询日志列表
	var logs []models.OperationLog
	query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&logs)

	// 构建响应数据
	var responseLogs []gin.H
	for _, log := range logs {
		responseLogs = append(responseLogs, gin.H{
			"log_id":      log.LogID,
			"user_id":     log.UserID,
			"username":    log.Username,
			"action":      log.Action,
			"description": log.Description,
			"ip_address":  log.IPAddress,
			"user_agent":  log.UserAgent,
			"created_at":  log.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取操作审计日志成功",
		"data": gin.H{
			"logs":  responseLogs,
			"total": total,
			"page":  page,
			"size":  pageSize,
		},
	})
}

// SetupTwoFactor 设置两步验证
// @Summary 设置两步验证
// @Description 为用户设置或更新两步验证
// @Tags 系统安全
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param enabled body bool true "是否启用两步验证"
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/security/2fa [post]
func SetupTwoFactor(c *gin.Context) {
	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 解析请求体
	var req struct {
		Enabled bool `json:"enabled" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	// 查询现有两步验证设置
	var twoFactor models.TwoFactor
	result := database.DB.Where("user_id = ?", userID).First(&twoFactor)

	if result.Error != nil {
		// 创建新的两步验证设置
		twoFactor = models.TwoFactor{
			UserID:  userID,
			Enabled: req.Enabled,
			// 实际项目中应该生成真实的密钥和恢复代码
			SecretKey:     "JBSWY3DPEHPK3PXP",
			RecoveryCodes: "123456,234567,345678,456789,567890",
		}
		if err := database.DB.Create(&twoFactor).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "设置两步验证失败",
				"data": gin.H{
					"error_details": err.Error(),
				},
			})
			return
		}
	} else {
		// 更新现有设置
		twoFactor.Enabled = req.Enabled
		if err := database.DB.Save(&twoFactor).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "更新两步验证失败",
				"data": gin.H{
					"error_details": err.Error(),
				},
			})
			return
		}
	}

	// 构建响应数据
	response := gin.H{
		"enabled":     twoFactor.Enabled,
		"secret_key":  twoFactor.SecretKey,
		"qr_code_url": "https://example.com/qrcode.png", // 实际项目中应该生成真实的二维码
		"recovery_codes": []string{
			"123456",
			"234567",
			"345678",
			"456789",
			"567890",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "设置两步验证成功",
		"data":    response,
	})
}

// GetLoginDevices 获取登录设备列表
// @Summary 获取登录设备列表
// @Description 获取用户的登录设备列表
// @Tags 系统安全
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Success 200 {object} gin.H{code=int, message=string, data=gin.H}
// @Router /api/v1/security/devices [get]
func GetLoginDevices(c *gin.Context) {
	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 查询登录设备列表
	var devices []models.LoginDevice
	result := database.DB.Where("user_id = ?", userID).Order("last_login DESC").Find(&devices)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取登录设备列表失败",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 构建响应数据
	var responseDevices []gin.H
	for _, device := range devices {
		responseDevices = append(responseDevices, gin.H{
			"device_id":   device.DeviceID,
			"device_name": device.DeviceName,
			"ip_address":  device.IPAddress,
			"user_agent":  device.UserAgent,
			"last_login":  device.LastLogin,
			"is_current":  device.IsCurrent,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取登录设备列表成功",
		"data": gin.H{
			"devices": responseDevices,
			"total":   len(responseDevices),
			"page":    1,
			"size":    len(responseDevices),
		},
	})
}

// LogoutDevice 下线异常设备
// @Summary 下线异常设备
// @Description 下线指定的登录设备
// @Tags 系统安全
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT令牌"
// @Param id path int true "设备ID"
// @Success 200 {object} gin.H{code=int, message=string, data=nil}
// @Router /api/v1/security/devices/{id} [delete]
func LogoutDevice(c *gin.Context) {
	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	userID := uint(1)

	// 获取设备ID
	deviceIDStr := c.Param("id")
	deviceID, err := strconv.ParseUint(deviceIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "请求参数错误",
			"data": gin.H{
				"error_details": "设备ID格式错误",
			},
		})
		return
	}

	// 查询设备是否存在且属于当前用户
	var device models.LoginDevice
	result := database.DB.Where("device_id = ? AND user_id = ?", deviceID, userID).First(&device)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "设备不存在",
			"data": gin.H{
				"error_details": result.Error.Error(),
			},
		})
		return
	}

	// 删除设备记录
	if err := database.DB.Delete(&device).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "下线设备失败",
			"data": gin.H{
				"error_details": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "下线设备成功",
		"data":    nil,
	})
}
