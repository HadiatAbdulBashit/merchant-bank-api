package main

import (
	"log"
	"merchant-bank-api/db"
	"merchant-bank-api/router"
)

func main() {
	// Connect to the MySQL database
	db.Connect()

	// Setup and run the Gin router
	r := router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
