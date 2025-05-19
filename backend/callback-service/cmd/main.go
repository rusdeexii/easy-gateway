package main

import (
    "log"
    "os"

    "callback-service/config"
    "callback-service/router"
    "callback-service/model"
)

func main() {
    // Load environment variables and connect to database
    config.LoadEnv()

    // Get the DB instance
    db := config.GetDB()

    // Run AutoMigrate to automatically create the table for Transaction
    if err := db.AutoMigrate(&model.Transaction{}); err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    // Set up the router
    r := router.SetupRouter()

    // Start the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8082"
    }

    log.Printf("Callback service is running on port %s", port)
    r.Run(":" + port)
}
