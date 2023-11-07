package utils

import (
	"dog_api/models"
	"fmt"
	"sync"
)

type InMemoryDB struct {
	Dogs      map[string]models.Dog
	idCounter int
	mu        sync.Mutex
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		Dogs: make(map[string]models.Dog),
	}
}

func (db *InMemoryDB) CreateDog(dog models.Dog) (string, error) {
	id := db.generateUUID()
	dog.ID = id
	db.Dogs[id] = dog
	return id, nil
}

func (db *InMemoryDB) GetDogs() ([]models.Dog, error) {
	dogs := make([]models.Dog, 0, len(db.Dogs))
	for _, dog := range db.Dogs {
		dogs = append(dogs, dog)
	}
	return dogs, nil
}

func (db *InMemoryDB) GetDogByID(id string) (models.Dog, error) {
	dog, found := db.Dogs[id]
	if !found {
		return models.Dog{}, fmt.Errorf("Dog not found")
	}
	return dog, nil
}

func (db *InMemoryDB) generateUUID() string {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.idCounter++
	return fmt.Sprintf("dog-%d", db.idCounter)
}
