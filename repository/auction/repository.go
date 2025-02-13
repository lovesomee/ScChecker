package auction

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"sc-profile/models"
)

type IRepository interface {
	InsertDeal(ctx context.Context, deal DbAuctionHistoryDeal) error
	BulkInsertDeal(ctx context.Context, itemId string, historyPrices []models.AuctionHistoryPrices) error
}

type Repository struct {
	logger *zap.Logger
	db     *sqlx.DB
}

func NewRepository(logger *zap.Logger, db *sqlx.DB) *Repository {
	return &Repository{logger: logger, db: db}
}

//go:embed sql/insert_deal.sql
var insertDealSql string

func (r *Repository) InsertDeal(ctx context.Context, deal DbAuctionHistoryDeal) error {
	stmt, err := r.db.PrepareNamed(insertDealSql)
	if err != nil {
		return fmt.Errorf("cannot prepare statement: %w", err)
	}

	if _, err = stmt.Exec(deal); err != nil {
		return fmt.Errorf("cannot execute statement: %w", err)
	}

	return nil
}

func (r *Repository) BulkInsertDeal(ctx context.Context, itemId string, historyPrices []models.AuctionHistoryPrices) error {
	upsertQuery := "INSERT INTO auction_history (item_id, amount, price, time, additional) VALUES (:item_id, :amount, :price, :time, :additional)"
	onConflictStatement := " ON CONFLICT (price, time) DO NOTHING"

	query, queryArgs, err := r.db.BindNamed(upsertQuery, AuctionHistoryPricesToDbDeals(itemId, historyPrices))
	if err != nil {
		return fmt.Errorf("error inserting deal into database: %w", err)
	}

	query = r.db.Rebind(query)
	query = query + onConflictStatement

	rows, err := r.db.QueryxContext(ctx, query, queryArgs...)
	if err != nil {
		return fmt.Errorf("error inserting deal into database: %w", err)
	}
	defer rows.Close()

	return nil
}
