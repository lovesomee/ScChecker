package updatelist

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"sc-profile/repository/updatelist"
)

type IService interface {
	AddUpdateList(context.Context, []string) error
	GetUpdateList(ctx context.Context) ([]string, error)
	DelUpdateList(ctx context.Context, updateList string) error
}

type Service struct {
	logger               *zap.Logger
	updateListRepository updatelist.IRepository
}

func NewService(logger *zap.Logger, updateListRepository updatelist.IRepository) *Service {
	return &Service{logger: logger, updateListRepository: updateListRepository}
}

func (s *Service) AddUpdateList(ctx context.Context, updateList []string) error {
	if err := s.updateListRepository.InsertUpdateList(ctx, updateList); err != nil {
		return fmt.Errorf("insertUpdateList error: %w", err)
	}

	return nil
}

func (s *Service) GetUpdateList(ctx context.Context) ([]string, error) {
	updateList, err := s.updateListRepository.SelectUpdateList(ctx)
	if err != nil {
		return nil, fmt.Errorf("selectUpdateList error: %w", err)
	}

	return updateList, err
}

func (s *Service) DelUpdateList(ctx context.Context, updateList string) error {
	if err := s.updateListRepository.DelUpdateList(ctx, updateList); err != nil {
		return fmt.Errorf("delUpdateList error: %w", err)
	}

	return nil
}
