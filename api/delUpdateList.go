package api

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"sc-profile/service/updatelist"
)

func DelUpdateList(logger *zap.Logger, updateListService updatelist.IService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		itemId := mux.Vars(r)["itemId"]

		if err := updateListService.DelUpdateList(r.Context(), itemId); err != nil {
			logger.Error("del update list error", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
