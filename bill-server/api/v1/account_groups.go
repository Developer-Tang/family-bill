package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAccountGroups 获取账户分组列表
func GetAccountGroups(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "获取成功", "data": []gin.H{}})
}

// CreateAccountGroup 创建账户分组
func CreateAccountGroup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "创建成功", "data": gin.H{}})
}

// UpdateAccountGroup 更新账户分组
func UpdateAccountGroup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "更新成功", "data": gin.H{}})
}

// DeleteAccountGroup 删除账户分组
func DeleteAccountGroup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "删除成功", "data": nil})
}

// AddAccountsToGroup 账户加入分组
func AddAccountsToGroup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "操作成功", "data": gin.H{}})
}