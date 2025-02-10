package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"net/http"
	"sc-profile/api"
	"sc-profile/cron"
	dbauction "sc-profile/repository/auction"
	"sc-profile/repository/updatelist"
	"sc-profile/service/auction"
	"sc-profile/service/scapi"
)

func main() {
	router := mux.NewRouter()
	db := newDatabase("postgres://root:rpass@localhost:5432/sc_db?sslmode=disable")
	err := goose.Up(db.DB, "migrations")
	if err != nil {
		panic(err)
	}

	auctionRepository := dbauction.NewRepository(db)
	updateListRepository := updatelist.NewRepository(db)
	stalcraftApi := scapi.NewScApi(http.DefaultClient, "https://eapi.stalcraft.net")
	auctionService := auction.NewService(stalcraftApi, auctionRepository, updateListRepository)
	scCron := cron.NewScCron(auctionService)
	scCron.Start()

	api.RegisterPing(router)
	server := newServer(router)
	server.ListenAndServe()
}

func newServer(router http.Handler) *http.Server {
	return &http.Server{
		Handler: router,
		Addr:    ":80",
	}
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
