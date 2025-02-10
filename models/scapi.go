package models

import "time"

type AuctionHistoryResponse struct {
	Total  int                    `json:"total"`
	Prices []AuctionHistoryPrices `json:"prices"`
}

type AuctionHistoryPrices struct {
	Amount     int       `json:"amount"`
	Price      int       `json:"price"`
	Time       time.Time `json:"time"`
	Additional any       `json:"additional"`
}
