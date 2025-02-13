package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
	"net/http"
	"sc-profile/api"
	"sc-profile/config"
	"sc-profile/cron"
	dbauction "sc-profile/repository/auction"
	dbupdatelist "sc-profile/repository/updatelist"
	"sc-profile/service/auction"
	"sc-profile/service/scapi"
	"sc-profile/service/updatelist"
)

func main() {
	cfg := config.Read()
	logger, _ := zap.NewProduction()
	db := newDatabase(cfg)
	err := goose.Up(db.DB, "migrations")
	if err != nil {
		panic(err)
	}

	auctionRepository := dbauction.NewRepository(logger, db)
	updateListRepository := dbupdatelist.NewRepository(logger, db)

	stalcraftApi := scapi.NewScApi(logger, cfg, http.DefaultClient)
	updateListService := updatelist.NewService(logger, updateListRepository)
	auctionService := auction.NewService(logger, stalcraftApi, auctionRepository, updateListRepository)

	scCron := cron.NewScCron(logger, auctionService)
	scCron.Start()

	server := api.NewServer(logger, cfg, updateListService)
	server.ListenAndServe()
}

func newDatabase(cfg config.Settings) *sqlx.DB {
	pool, err := pgxpool.New(context.Background(), cfg.Database.PostgresConnection)
	if err != nil {
		panic(err)
	}

	return sqlx.NewDb(stdlib.OpenDBFromPool(pool), "pgx")
}

//func calculatePercent(minValue, maxValue, statsRandom float64) float64 {
//	calcRange := maxValue - minValue
//	return minValue + calcRange*((statsRandom+2)/4)
//}
