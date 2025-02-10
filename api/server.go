package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"sc-profile/service/updatelist"
)

func NewServer(updateListService updatelist.IService) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/update-list", AddUpdateList(updateListService)).Methods(http.MethodPost)
	router.HandleFunc("/ping", Ping()).Methods(http.MethodGet)

	return &http.Server{
		Handler: router,
		Addr:    ":80",
	}
}
