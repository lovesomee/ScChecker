package auction

import (
	"sc-profile/repository/auction"
	"sc-profile/repository/mapper"
	"sc-profile/repository/updatelist"
	"sc-profile/service/scapi"
)

type IService interface {
	UpdateItemHistory() error
}

type Service struct {
	StalcraftApi         scapi.IScApi
	AuctionRepository    auction.IRepository
	UpdateListRepository updatelist.IRepository
}

func NewService(stalcraftApi scapi.IScApi, auctionRepository auction.IRepository, updateListRepository updatelist.IRepository) *Service {
	return &Service{StalcraftApi: stalcraftApi, AuctionRepository: auctionRepository, UpdateListRepository: updateListRepository}
}

func (s *Service) UpdateItemHistory() error {
	dbUpdateList, err := s.UpdateListRepository.SelectUpdateList() //получаем массив айдишников по которым будем искать
	if err != nil {
		return err
	}

	for _, item := range dbUpdateList {
		auctionHistory, err := s.StalcraftApi.GetAuctionHistory(item.ItemId, "ru", 200) //формируем http запрос на получение резульатов от api, возвращаем заполненную структуру с ответом
		if err != nil {
			return err
		}

		if err = s.AuctionRepository.BulkInsertDeal(mapper.AuctionHistoryPricesToDbDeals(item.ItemId, auctionHistory.Prices)); err != nil { //выполняем массовую вставку в БД
			return err
		}
	}

	return nil
}
