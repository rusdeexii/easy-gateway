package main

import (
	"log"
	"os"

	"merchant-service/config"
	"merchant-service/router"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env.dev if it exists
	if err := godotenv.Load(".env.dev"); err != nil {
		log.Println("No .env.dev file found")
	}

	// Connect to the database
	config.ConnectDB()

	// Setup router
	r := router.SetupRouter()

	// Get port from environment variable, default to 8084 if not provided
	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}

	// Start the server
	log.Printf("Merchant Service running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
