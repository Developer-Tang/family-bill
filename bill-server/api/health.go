package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthCheck 健康检查接口
// @Summary 健康检查
// @Description 检查服务是否正常运行
// @Tags 公共服务
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	// 从上下文获取版本号
	version, exists := c.Get("version")
	if !exists {
		version = "unknown"
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "服务正常运行",
		"data": gin.H{
			"status":    "ok",
			"timestamp": time.Now(),
			"version":   version.(string),
		},
	})
}
