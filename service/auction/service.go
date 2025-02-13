package auction

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"sc-profile/repository/auction"
	"sc-profile/repository/updatelist"
	"sc-profile/service/scapi"
)

type IService interface {
	UpdateItemHistory(ctx context.Context) error
}

type Service struct {
	logger               *zap.Logger
	stalcraftApi         scapi.IScApi
	auctionRepository    auction.IRepository
	updateListRepository updatelist.IRepository
}

func NewService(logger *zap.Logger, stalcraftApi scapi.IScApi, auctionRepository auction.IRepository, updateListRepository updatelist.IRepository) *Service {
	return &Service{logger: logger, stalcraftApi: stalcraftApi, auctionRepository: auctionRepository, updateListRepository: updateListRepository}
}

func (s *Service) UpdateItemHistory(ctx context.Context) error {
	dbUpdateList, err := s.updateListRepository.SelectUpdateList(ctx) //получаем массив айдишников по которым будем искать
	if err != nil {
		return fmt.Errorf("selectItemHistory error: %w", err)
	}

	for _, item := range dbUpdateList {
		auctionHistory, err := s.stalcraftApi.GetAuctionHistory(ctx, item.ItemId, "ru", 200) //формируем http запрос на получение резульатов от api, возвращаем заполненную структуру с ответом
		if err != nil {
			return fmt.Errorf("GetAuctionHistory error: %w", err)
		}

		if err = s.auctionRepository.BulkInsertDeal(ctx, item.ItemId, auctionHistory.Prices); err != nil { //выполняем массовую вставку в БД
			return fmt.Errorf("bulkInsertDeal error: %w", err)
		}
	}

	return nil
}
