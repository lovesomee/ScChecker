package api

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"sc-profile/service/updatelist"
)

func NewServer(logger *zap.Logger, updateListService updatelist.IService) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/update-list", AddUpdateList(logger, updateListService)).Methods(http.MethodPost)
	router.HandleFunc("/ping", Ping()).Methods(http.MethodGet)

	return &http.Server{
		Handler: router,
		Addr:    ":80",
	}
}
