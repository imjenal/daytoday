package app

import (
	"daytoday/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func addRoutes(router *mux.Router) error {
	router.Handle("/", handler.Heartbeat()).Methods(http.MethodGet)
	return nil
}