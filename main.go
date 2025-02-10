package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"net/http"
	"sc-profile/api"
	"sc-profile/cron"
	dbauction "sc-profile/repository/auction"
	dbupdatelist "sc-profile/repository/updatelist"
	"sc-profile/service/auction"
	"sc-profile/service/scapi"
	"sc-profile/service/updatelist"
)

func main() {
	db := newDatabase("postgres://root:rpass@localhost:5432/sc_db?sslmode=disable")
	err := goose.Up(db.DB, "migrations")
	if err != nil {
		panic(err)
	}

	auctionRepository := dbauction.NewRepository(db)
	updateListRepository := dbupdatelist.NewRepository(db)

	stalcraftApi := scapi.NewScApi(http.DefaultClient, "https://eapi.stalcraft.net")
	updateListService := updatelist.NewService(updateListRepository)
	auctionService := auction.NewService(stalcraftApi, auctionRepository, updateListRepository)

	scCron := cron.NewScCron(auctionService)
	scCron.Start()

	server := api.NewServer(updateListService)
	server.ListenAndServe()
}

func newDatabase(connect string) *sqlx.DB {
	pool, err := pgxpool.New(context.Background(), connect)
	if err != nil {
		panic(err)
	}

	return sqlx.NewDb(stdlib.OpenDBFromPool(pool), "pgx")
}

//func calculatePercent(minValue, maxValue, statsRandom float64) float64 {
//	calcRange := maxValue - minValue
//	return minValue + calcRange*((statsRandom+2)/4)
//}
