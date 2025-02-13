package updatelist

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"sc-profile/models"
	"sc-profile/repository/updatelist"
)

type IService interface {
	AddUpdateList(context.Context, models.UpdateList) error
}

type Service struct {
	logger               *zap.Logger
	updateListRepository updatelist.IRepository
}

func NewService(logger *zap.Logger, updateListRepository updatelist.IRepository) *Service {
	return &Service{logger: logger, updateListRepository: updateListRepository}
}

func (s *Service) AddUpdateList(ctx context.Context, updateList models.UpdateList) error {
	if err := s.updateListRepository.InsertUpdateList(ctx, updateList); err != nil {
		return fmt.Errorf("insertUpdateList error: %w", err)
	}

	return nil
}
