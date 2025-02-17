package api

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"sc-profile/service/updatelist"
)

func AddUpdateList(logger *zap.Logger, updateListService updatelist.IService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateList []string

		if err := json.NewDecoder(r.Body).Decode(&updateList); err != nil {
			logger.Error("incorrect json format error", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := updateListService.AddUpdateList(r.Context(), updateList); err != nil {
			logger.Error("addUpdateList error", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
