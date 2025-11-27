package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRoles 获取角色列表
func GetRoles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "获取成功", "data": []gin.H{}})
}

// CreateRole 创建自定义角色
func CreateRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "创建成功", "data": gin.H{}})
}

// UpdateRole 更新角色
func UpdateRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "更新成功", "data": gin.H{}})
}

// DeleteRole 删除角色
func DeleteRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "删除成功", "data": gin.H{}})
}