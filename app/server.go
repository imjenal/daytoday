package app

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"fmt"
)

const (
	PORT  = "8080"
	HOST = "0.0.0.0"
)

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)

	allowedHeaders := handlers.AllowedHeaders([]string{"Origin", "Accept", "Content-Type", "Authorization", "DateUsed", "X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions})

	if err := addRoutes(router); err != nil {
		log.Println("Failed adding routes", err.Error())
	}
	address := net.JoinHostPort(HOST, PORT)
	fmt.Println("Server started successfully at", address)
	if err := http.ListenAndServe(address, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)); err != nil {
		log.Fatal(err)
	}


}
