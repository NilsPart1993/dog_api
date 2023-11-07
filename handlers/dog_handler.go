package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"dog_api/models"
	"dog_api/utils"
	"time"
)

var db utils.Database

func SetDatabase(database utils.Database) {
	db = database
}

func CreateDog(w http.ResponseWriter, r *http.Request) {
	var dogRequest models.DogRequest
	if err := utils.ParseJSON(r, &dogRequest); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	birthDate, err := time.Parse("2006-01-02", dogRequest.BirthDate)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid birthDate format")
		return
	}

	id, err := db.CreateDog(models.Dog{
		BirthDate: birthDate,
		Breed:     dogRequest.Breed,
		Name:      dogRequest.Name,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create dog")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, id)
}

func GetDogs(w http.ResponseWriter, r *http.Request) {
	dogs, err := db.GetDogs()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve dogs")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, dogs)
}

func GetDogByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dogID := params["id"]

	dog, err := db.GetDogByID(dogID)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, dog)
}
