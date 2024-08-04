package main

import (
	"log"

	"github.com/Vghxv/GinHub/database"
	"github.com/Vghxv/GinHub/router"
)

func main() {
	// Initialize database connection
	database.Connect()

	// Initialize router
	r := router.SetupRouter()

	// Run the server
	log.Fatal(r.Run("0.0.0.0:8080"))
}
