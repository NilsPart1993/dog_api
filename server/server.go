package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer() {
	router := mux.NewRouter()
	SetDogRoutes(router)
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}