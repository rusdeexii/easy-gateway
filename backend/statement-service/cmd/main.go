package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "statement-service/config"
    "statement-service/router"
)

func main() {
    if err := godotenv.Load(".env.dev"); err != nil {
        log.Println("No .env.dev file found")
    }

    // เรียกใช้ฟังก์ชัน ConnectDB จาก config
    config.ConnectDB()

    // กำหนด router
    r := router.SetupRouter()

    // กำหนด port
    port := os.Getenv("PORT")
    if port == "" {
        port = "8086"
    }

    log.Printf("Statement Service running on port %s", port)
    r.Run(":" + port)
}
