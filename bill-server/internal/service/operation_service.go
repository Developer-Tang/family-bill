package service

import (
	"github.com/family-bill/bill-server/internal/db"
	"github.com/family-bill/bill-server/internal/model"
)

// OperationService 操作服务
type OperationService struct{}

// NewOperationService 创建操作服务实例
func NewOperationService() *OperationService {
	return &OperationService{}
}

// CreateOperation 创建操作记录
func (s *OperationService) CreateOperation(operation *model.Operation) error {
	// 保存操作记录
	if err := db.GetDB().Create(operation).Error; err != nil {
		return err
	}
	return nil
}
