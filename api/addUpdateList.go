package api

import (
	"encoding/json"
	"net/http"
	"sc-profile/models"
	"sc-profile/service/updatelist"
)

func AddUpdateList(updateListService updatelist.IService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateList models.UpdateList

		if err := json.NewDecoder(r.Body).Decode(&updateList); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := updateListService.AddUpdateList(updateList); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
