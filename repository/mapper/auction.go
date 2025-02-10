package mapper

import (
	"sc-profile/repository/auction"
	"sc-profile/service/scapi"
)

func AuctionHistoryPricesToDbDeals(itemId string, auctionHistoryPrices []scapi.AuctionHistoryPrices) []auction.DbAuctionHistoryDeal {
	auctionHistoryDeals := make([]auction.DbAuctionHistoryDeal, 0, len(auctionHistoryPrices))

	for _, item := range auctionHistoryPrices {
		auctionHistoryDeals = append(auctionHistoryDeals, auction.DbAuctionHistoryDeal{
			ItemId:     itemId,
			Amount:     item.Amount,
			Price:      item.Price,
			Time:       item.Time,
			Additional: item.Additional,
		})
	}

	return auctionHistoryDeals
}

func AuctionHistoryPriceToDbDeal(itemId string, auctionHistoryPrice scapi.AuctionHistoryPrices) auction.DbAuctionHistoryDeal {
	return auction.DbAuctionHistoryDeal{
		ItemId:     itemId,
		Amount:     auctionHistoryPrice.Amount,
		Price:      auctionHistoryPrice.Price,
		Time:       auctionHistoryPrice.Time,
		Additional: auctionHistoryPrice.Additional,
	}
}
