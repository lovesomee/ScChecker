package api

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"sc-profile/service/updatelist"
)

func GetUpdateList(logger *zap.Logger, updateListService updatelist.IService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		updateList, err := updateListService.GetUpdateList(r.Context())
		if err != nil {
			logger.Error("getUpdateList error", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonResp, err := json.Marshal(updateList)
		if err != nil {
			logger.Error("getUpdateList marshal error", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}
}
