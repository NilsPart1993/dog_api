package server

import (
	"github.com/gorilla/mux"
	"dog_api/handlers"
)

func SetDogRoutes(router *mux.Router) {
	router.HandleFunc("/dogs", handlers.CreateDog).Methods("POST")
	router.HandleFunc("/dogs", handlers.GetDogs).Methods("GET")
	router.HandleFunc("/dogs/{id}", handlers.GetDogByID).Methods("GET")
}