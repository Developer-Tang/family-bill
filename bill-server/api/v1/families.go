package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetFamilies 获取家庭组列表
func GetFamilies(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": []gin.H{
			{
				"family_id": 1,
				"name": "我的家庭",
				"description": "家庭记账群组",
				"avatar": "https://example.com/family_avatar.jpg",
				"creator_id": 1,
				"created_at": "2023-01-01T12:00:00Z",
				"members_count": 4
			}
		}
	})
}

// CreateFamily 创建家庭组
func CreateFamily(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "创建成功",
		"data": gin.H{
			"family_id": 1,
			"name": "我的家庭",
			"description": "家庭记账群组",
			"created_at": "2023-01-01T12:00:00Z",
			"members_count": 1
		}
	})
}

// UpdateFamily 更新家庭组信息
func UpdateFamily(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "更新成功",
		"data": gin.H{
			"family_id": 1,
			"name": "我的新家庭",
			"description": "更新后的家庭记账群组",
			"avatar": "https://example.com/new_family_avatar.jpg",
			"updated_at": "2023-01-05T12:30:00Z"
		}
	})
}

// DeleteFamily 删除家庭组
func DeleteFamily(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
		"data": gin.H{
			"family_id": 1,
			"deleted_at": "2023-01-05T12:30:00Z"
		}
	})
}

// InviteFamilyMember 邀请成员加入家庭组
func InviteFamilyMember(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "邀请发送成功",
		"data": gin.H{
			"family_id": 1,
			"invite_id": "invite_123",
			"email": "member@example.com",
			"expires_in": 86400
		}
	})
}

// LeaveFamily 退出家庭组
func LeaveFamily(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "退出成功",
		"data": gin.H{
			"family_id": 1,
			"user_id": 2,
			"left_at": "2023-01-05T12:30:00Z"
		}
	})
}

// GetFamilyMembers 获取家庭组成员列表
func GetFamilyMembers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取成功",
		"data": []gin.H{
			{
				"user_id": 1,
				"username": "张三",
				"avatar": "https://example.com/avatar1.jpg",
				"role": "admin",
				"joined_at": "2023-01-01T12:00:00Z"
			},
			{
				"user_id": 2,
				"username": "李四",
				"avatar": "https://example.com/avatar2.jpg",
				"role": "member",
				"joined_at": "2023-01-02T12:00:00Z"
			}
		}
	})
}

// RemoveFamilyMember 移除家庭组成员
func RemoveFamilyMember(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "移除成功",
		"data": gin.H{
			"family_id": 1,
			"user_id": 2,
			"removed_at": "2023-01-05T12:30:00Z"
		}
