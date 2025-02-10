package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterPing(router *mux.Router) {
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	}).Methods(http.MethodPost)
}
