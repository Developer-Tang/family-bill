package v1

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/family-bill/bill-server/internal/database"
	"github.com/family-bill/bill-server/models"
	"github.com/family-bill/bill-server/utils"
)

// FamilyRequest 家庭组请求结构体
type FamilyRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=50"`
	Description string `json:"description" binding:"omitempty,max=255"`
	Avatar      string `json:"avatar" binding:"omitempty,url"`
}

// InviteRequest 邀请请求结构体
type InviteRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// GetFamilies 获取家庭组列表
func GetFamilies(c *gin.Context) {
	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	// 这里简化处理，使用固定值1
	userID := uint(1)

	// 查询用户加入的所有家庭组
	var familyMembers []models.FamilyMember
	result := database.DB.Where("user_id = ?", userID).Find(&familyMembers)
	if result.Error != nil {
		utils.ErrorResponseWithData(c, 500, "获取家庭组列表失败", gin.H{
			"error_details": result.Error.Error(),
		})
		return
	}

	// 提取家庭组ID列表
	var familyIDs []uint
	for _, member := range familyMembers {
		familyIDs = append(familyIDs, member.FamilyID)
	}

	// 查询家庭组详细信息
	var families []models.Family
	if len(familyIDs) > 0 {
		result = database.DB.Where("family_id IN ?", familyIDs).Find(&families)
		if result.Error != nil {
			utils.ErrorResponseWithData(c, 500, "获取家庭组详情失败", gin.H{
				"error_details": result.Error.Error(),
			})
			return
		}
	}

	// 准备响应数据
	var response []gin.H
	for _, family := range families {
		// 查询成员数量
		var memberCount int64
		database.DB.Model(&models.FamilyMember{}).Where("family_id = ?", family.FamilyID).Count(&memberCount)

		response = append(response, gin.H{
			"family_id":     family.FamilyID,
			"name":          family.Name,
			"description":   family.Description,
			"avatar":        family.Avatar,
			"creator_id":    family.CreatorID,
			"created_at":    family.CreatedAt,
			"members_count": memberCount,
		})
	}

	utils.SuccessResponseWithMessage(c, "获取成功", response)
}

// CreateFamily 创建家庭组
func CreateFamily(c *gin.Context) {
	var req FamilyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	// 这里简化处理，使用固定值1
	userID := uint(1)

	// 创建家庭组
	family := models.Family{
		Name:        req.Name,
		Description: req.Description,
		Avatar:      req.Avatar,
		CreatorID:   userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// 开始事务
	tx := database.DB.Begin()
	if err := tx.Create(&family).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponseWithData(c, 500, "创建家庭组失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 添加创建者为成员
	familyMember := models.FamilyMember{
		FamilyID:  family.FamilyID,
		UserID:    userID,
		Role:      "admin",
		JoinedAt:  time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := tx.Create(&familyMember).Error; err != nil {
		tx.Rollback()
		utils.ErrorResponseWithData(c, 500, "添加成员失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 提交事务
	tx.Commit()

	utils.SuccessResponseWithMessage(c, "创建成功", gin.H{
		"family_id":     family.FamilyID,
		"name":          family.Name,
		"description":   family.Description,
		"avatar":        family.Avatar,
		"created_at":    family.CreatedAt,
		"members_count": 1,
	})
}

// UpdateFamily 更新家庭组信息
func UpdateFamily(c *gin.Context) {
	// 获取家庭组ID
	familyIDStr := c.Param("id")
	familyID, err := strconv.ParseUint(familyIDStr, 10, 32)
	if err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": "无效的家庭组ID",
		})
		return
	}

	var req FamilyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	// 这里简化处理，使用固定值1
	userID := uint(1)

	// 检查用户是否是家庭组成员且具有管理员权限
	var familyMember models.FamilyMember
	result := database.DB.Where("family_id = ? AND user_id = ? AND role IN ?", familyID, userID, []string{"admin", "owner"}).First(&familyMember)
	if result.Error != nil {
		utils.ErrorResponseWithData(c, 403, "权限不足", gin.H{
			"error_details": "您没有权限修改该家庭组",
		})
		return
	}

	// 更新家庭组信息
	var family models.Family
	result = database.DB.First(&family, familyID)
	if result.Error != nil {
		utils.ErrorResponseWithData(c, 404, "家庭组不存在", gin.H{
			"error_details": result.Error.Error(),
		})
		return
	}

	family.Name = req.Name
	family.Description = req.Description
	if req.Avatar != "" {
		family.Avatar = req.Avatar
	}
	family.UpdatedAt = time.Now()

	if err := database.DB.Save(&family).Error; err != nil {
		utils.ErrorResponseWithData(c, 500, "更新失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	utils.SuccessResponseWithMessage(c, "更新成功", gin.H{
		"family_id":   family.FamilyID,
		"name":        family.Name,
		"description": family.Description,
		"avatar":      family.Avatar,
		"updated_at":  family.UpdatedAt,
	})
}

// DeleteFamily 删除家庭组
func DeleteFamily(c *gin.Context) {
	// 获取家庭组ID
	familyIDStr := c.Param("id")
	familyID, err := strconv.ParseUint(familyIDStr, 10, 32)
	if err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": "无效的家庭组ID",
		})
		return
	}

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	// 这里简化处理，使用固定值1
	userID := uint(1)

	// 检查用户是否是家庭组创建者
	var family models.Family
	result := database.DB.Where("family_id = ? AND creator_id = ?", familyID, userID).First(&family)
	if result.Error != nil {
		utils.ErrorResponseWithData(c, 403, "权限不足", gin.H{
			"error_details": "只有创建者可以删除家庭组",
		})
		return
	}

	// 删除家庭组（软删除）
	if err := database.DB.Delete(&family).Error; err != nil {
		utils.ErrorResponseWithData(c, 500, "删除失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	utils.SuccessResponseWithMessage(c, "删除成功", gin.H{
		"family_id":  family.FamilyID,
		"deleted_at": time.Now().Format(time.RFC3339),
	})
}

// InviteFamilyMember 邀请成员加入家庭组
func InviteFamilyMember(c *gin.Context) {
	// 获取家庭组ID
	familyIDStr := c.Param("id")
	familyID, err := strconv.ParseUint(familyIDStr, 10, 32)
	if err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": "无效的家庭组ID",
		})
		return
	}

	var req InviteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	// 这里简化处理，使用固定值1
	userID := uint(1)

	// 检查用户是否是家庭组成员且具有管理员权限
	var familyMember models.FamilyMember
	result := database.DB.Where("family_id = ? AND user_id = ? AND role IN ?", familyID, userID, []string{"admin", "owner"}).First(&familyMember)
	if result.Error != nil {
		utils.ErrorResponseWithData(c, 403, "权限不足", gin.H{
			"error_details": "您没有权限邀请成员",
		})
		return
	}

	// 检查邮箱是否已注册
	var existingUser models.User
	result = database.DB.Where("email = ?", req.Email).First(&existingUser)
	if result.Error != nil {
		utils.ErrorResponseWithData(c, 404, "用户不存在", gin.H{
			"error_details": "该邮箱尚未注册",
		})
		return
	}

	// 检查用户是否已加入该家庭组
	result = database.DB.Where("family_id = ? AND user_id = ?", familyID, existingUser.UserID).First(&models.FamilyMember{})
	if result.Error == nil {
		utils.ErrorResponseWithData(c, 400, "邀请失败", gin.H{
			"error_details": "该用户已加入该家庭组",
		})
		return
	}

	// 这里应该发送邀请邮件，实际项目中需要实现邮件发送逻辑
	// 生成邀请码或邀请链接
	inviteID := "invite_" + time.Now().Format("20060102150405")

	utils.SuccessResponseWithMessage(c, "邀请发送成功", gin.H{
		"family_id":  familyID,
		"invite_id":  inviteID,
		"email":      req.Email,
		"expires_in": 86400,
	})
}

// LeaveFamily 退出家庭组
func LeaveFamily(c *gin.Context) {
	// 获取家庭组ID
	familyIDStr := c.Param("id")
	familyID, err := strconv.ParseUint(familyIDStr, 10, 32)
	if err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": "无效的家庭组ID",
		})
		return
	}

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	// 这里简化处理，使用固定值1
	userID := uint(1)

	// 检查用户是否是家庭组成员
	var familyMember models.FamilyMember
	result := database.DB.Where("family_id = ? AND user_id = ?", familyID, userID).First(&familyMember)
	if result.Error != nil {
		utils.ErrorResponseWithData(c, 404, "您不是该家庭组成员", gin.H{
			"error_details": result.Error.Error(),
		})
		return
	}

	// 检查用户是否是创建者
	var family models.Family
	result = database.DB.Where("family_id = ? AND creator_id = ?", familyID, userID).First(&family)
	if result.Error == nil {
		utils.ErrorResponseWithData(c, 400, "退出失败", gin.H{
			"error_details": "创建者不能退出家庭组，只能删除家庭组",
		})
		return
	}

	// 删除家庭成员关系
	if err := database.DB.Delete(&familyMember).Error; err != nil {
		utils.ErrorResponseWithData(c, 500, "退出失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	utils.SuccessResponseWithMessage(c, "退出成功", gin.H{
		"family_id": familyID,
		"user_id":   userID,
		"left_at":   time.Now().Format(time.RFC3339),
	})
}

// GetFamilyMembers 获取家庭组成员列表
func GetFamilyMembers(c *gin.Context) {
	// 获取家庭组ID
	familyIDStr := c.Param("id")
	familyID, err := strconv.ParseUint(familyIDStr, 10, 32)
	if err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": "无效的家庭组ID",
		})
		return
	}

	// 从上下文获取用户ID - 实际项目中应该从JWT令牌中解析
	// 这里简化处理，使用固定值1
	userID := uint(1)

	// 检查用户是否是家庭组成员
	var familyMember models.FamilyMember
	result := database.DB.Where("family_id = ? AND user_id = ?", familyID, userID).First(&familyMember)
	if result.Error != nil {
		utils.ErrorResponseWithData(c, 403, "权限不足", gin.H{
			"error_details": "您不是该家庭组成员",
		})
		return
	}

	// 查询家庭组成员列表
	var members []models.FamilyMember
	result = database.DB.Where("family_id = ?", familyID).Preload("User").Find(&members)
	if result.Error != nil {
		utils.ErrorResponseWithData(c, 500, "获取成员列表失败", gin.H{
			"error_details": result.Error.Error(),
		})
		return
	}

	// 准备响应数据
	var response []gin.H
	for _, member := range members {
		response = append(response, gin.H{
			"user_id":   member.User.UserID,
			"username":  member.User.Username,
			"avatar":    member.User.Avatar,
			"role":      member.Role,
			"joined_at": member.JoinedAt,
		})
	}

	utils.SuccessResponseWithMessage(c, "获取成功", response)
}

// RemoveFamilyMember 移除家庭组成员
func RemoveFamilyMember(c *gin.Context) {
	// 获取家庭组ID和要移除的用户ID
	familyIDStr := c.Param("id")
	familyID, err := strconv.ParseUint(familyIDStr, 10, 32)
	if err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": "无效的家庭组ID",
		})
		return
	}

	removeUserIDStr := c.Param("userId")
	removeUserID, err := strconv.ParseUint(removeUserIDStr, 10, 32)
	if err != nil {
		utils.ErrorResponseWithData(c, 400, "请求参数错误", gin.H{
			"error_details": "无效的用户ID",
		})
		return
	}

	// 从上下文获取当前用户ID - 实际项目中应该从JWT令牌中解析
	// 这里简化处理，使用固定值1
	currentUserID := uint(1)

	// 检查当前用户是否是家庭组成员且具有管理员权限
	var currentMember models.FamilyMember
	result := database.DB.Where("family_id = ? AND user_id = ? AND role IN ?", familyID, currentUserID, []string{"admin", "owner"}).First(&currentMember)
	if result.Error != nil {
		utils.ErrorResponseWithData(c, 403, "权限不足", gin.H{
			"error_details": "您没有权限移除成员",
		})
		return
	}

	// 检查要移除的用户是否是家庭组成员
	var removeMember models.FamilyMember
	result = database.DB.Where("family_id = ? AND user_id = ?", familyID, removeUserID).First(&removeMember)
	if result.Error != nil {
		utils.ErrorResponseWithData(c, 404, "成员不存在", gin.H{
			"error_details": "要移除的用户不是该家庭组成员",
		})
		return
	}

	// 不能移除创建者
	var family models.Family
	result = database.DB.Where("family_id = ? AND creator_id = ?", familyID, removeUserID).First(&family)
	if result.Error == nil {
		utils.ErrorResponseWithData(c, 400, "移除失败", gin.H{
			"error_details": "不能移除家庭组创建者",
		})
		return
	}

	// 移除成员
	if err := database.DB.Delete(&removeMember).Error; err != nil {
		utils.ErrorResponseWithData(c, 500, "移除失败", gin.H{
			"error_details": err.Error(),
		})
		return
	}

	utils.SuccessResponseWithMessage(c, "移除成功", gin.H{
		"family_id":  familyID,
		"user_id":    removeUserID,
		"removed_at": time.Now().Format(time.RFC3339),
	})
}
