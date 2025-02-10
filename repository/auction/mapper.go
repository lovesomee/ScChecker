package auction

import (
	"sc-profile/models"
)

func AuctionHistoryPricesToDbDeals(itemId string, auctionHistoryPrices []models.AuctionHistoryPrices) []DbAuctionHistoryDeal {
	auctionHistoryDeals := make([]DbAuctionHistoryDeal, 0, len(auctionHistoryPrices))

	for _, item := range auctionHistoryPrices {
		auctionHistoryDeals = append(auctionHistoryDeals, DbAuctionHistoryDeal{
			ItemId:     itemId,
			Amount:     item.Amount,
			Price:      item.Price,
			Time:       item.Time,
			Additional: item.Additional,
		})
	}

	return auctionHistoryDeals
}

func AuctionHistoryPriceToDbDeal(itemId string, auctionHistoryPrice models.AuctionHistoryPrices) DbAuctionHistoryDeal {
	return DbAuctionHistoryDeal{
		ItemId:     itemId,
		Amount:     auctionHistoryPrice.Amount,
		Price:      auctionHistoryPrice.Price,
		Time:       auctionHistoryPrice.Time,
		Additional: auctionHistoryPrice.Additional,
	}
}
