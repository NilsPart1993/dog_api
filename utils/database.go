package utils

import "dog_api/models"

type Database interface {
	CreateDog(dog models.Dog) (string, error)
	GetDogs() ([]models.Dog, error)
	GetDogByID(id string) (models.Dog, error)
}