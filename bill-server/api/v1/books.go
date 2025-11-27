package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBooks 获取账本列表
func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "获取成功", "data": []gin.H{}})
}

// CreateBook 创建账本
func CreateBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "创建成功", "data": gin.H{}})
}

// GetBook 获取账本详情
func GetBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "获取成功", "data": gin.H{}})
}

// UpdateBook 更新账本信息
func UpdateBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "更新成功", "data": gin.H{}})
}

// DeleteBook 删除账本
func DeleteBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "删除成功", "data": nil})
}

// GetBookMembers 获取账本成员列表
func GetBookMembers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "获取成功", "data": []gin.H{}})
}

// AddBookMember 添加账本成员
func AddBookMember(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "添加成功", "data": gin.H{}})
}

// RemoveBookMember 移除账本成员
func RemoveBookMember(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "移除成功", "data": nil})
}

// UpdateBookMemberRole 更新账本成员角色
func UpdateBookMemberRole(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "更新成功", "data": gin.H{}})
}

// SetDefaultBook 设置默认账本
func SetDefaultBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "设置成功", "data": gin.H{}})
}
