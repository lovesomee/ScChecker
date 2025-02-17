package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"sc-profile/config"
	"sc-profile/service/updatelist"
)

func NewServer(logger *zap.Logger, cfg config.Settings, updateListService updatelist.IService) *http.Server {
	router := mux.NewRouter()

	middleware := NewMiddleware(logger)

	router.Use(middleware.panicRecoverMiddleware, middleware.loggingMiddleware)
	router.HandleFunc("/update-list", AddUpdateList(logger, updateListService)).Methods(http.MethodPost)
	router.HandleFunc("/update-list", GetUpdateList(logger, updateListService)).Methods(http.MethodGet)
	router.HandleFunc("/ping", Ping()).Methods(http.MethodGet)

	return &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", cfg.Port),
	}
}
