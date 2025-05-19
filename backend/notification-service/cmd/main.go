package main

import (
    "log"
    "notification-service/config"
    "notification-service/consumer"
)

func main() {
    // โหลดการตั้งค่า .env
    config.LoadEnv()

    log.Println("Notification Service starting...")

    // เริ่มตัว consumer ที่จะเชื่อมต่อกับ RabbitMQ
    consumer.StartConsumer()
}
