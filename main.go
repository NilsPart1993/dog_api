package main

import (
	"dog_api/handlers"
	"dog_api/server"
	"dog_api/utils"
)

func main() {
	// Initialize the in-memory database
	db := utils.NewInMemoryDB()

	// Set the database implementation in your handlers
	handlers.SetDatabase(db)

	server.StartServer()
}