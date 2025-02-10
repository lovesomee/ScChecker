package auction

import "time"

type DbAuctionHistoryDeal struct {
	ItemId     string    `db:"item_id"`
	Amount     int       `db:"amount"`
	Price      int       `db:"price"`
	Time       time.Time `db:"time"`
	Additional any       `db:"additional"`
}
