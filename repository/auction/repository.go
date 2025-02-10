package auction

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
)

type IRepository interface {
	InsertDeal(deal DbAuctionHistoryDeal) error
	BulkInsertDeal(deals []DbAuctionHistoryDeal) error
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

//go:embed sql/insert_deal.sql
var insertDealSql string

func (r *Repository) InsertDeal(deal DbAuctionHistoryDeal) error {
	stmt, err := r.db.PrepareNamed(insertDealSql)
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(deal); err != nil {
		return err
	}

	return nil
}

func (r *Repository) BulkInsertDeal(deals []DbAuctionHistoryDeal) error {
	upsertQuery := "INSERT INTO auction_history (item_id, amount, price, time, additional) VALUES (:item_id, :amount, :price, :time, :additional)"
	onConflictStatement := " ON CONFLICT (price, time) DO NOTHING"

	query, queryArgs, err := r.db.BindNamed(upsertQuery, deals)
	if err != nil {
		return err
	}

	query = r.db.Rebind(query)
	query = query + onConflictStatement

	rows, err := r.db.Queryx(query, queryArgs...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
