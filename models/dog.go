package models

import (
	"time"
)

type Dog struct {
	ID        string    `json:"id"`
	BirthDate time.Time `json:"birthDate"`
	Breed     string    `json:"breed"`
	Name      string    `json:"name"`
}

type DogRequest struct {
	BirthDate string `json:"birthDate"`
	Breed     string `json:"breed"`
	Name      string `json:"name"`
}
