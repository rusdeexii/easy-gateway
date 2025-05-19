package main

import (
    "log"
    "agent-service/consumer"
    "github.com/joho/godotenv"
)

func main() {
    // โหลดค่าจากไฟล์ .env
    if err := godotenv.Load(".env.dev"); err != nil {
        log.Println("No .env.dev file found")
    }

    log.Println("Agent Service starting...")

    // เรียกใช้ consumer
    consumer.StartConsumer()
}
