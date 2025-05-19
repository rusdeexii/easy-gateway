package consumer

import (
    "encoding/json"
    "log"
    "notification-service/model"
    "notification-service/handler"
    "notification-service/config" // เพิ่มการ import config
    "github.com/streadway/amqp"
)

func StartConsumer() {
    // ใช้การตั้งค่าจาก config
    uri := config.GetRabbitMQURI() // ใช้ฟังก์ชั่น GetRabbitMQURI ที่เราสร้างใน config.go

    conn, err := amqp.Dial(uri)
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open channel: %v", err)
    }
    defer ch.Close()

    q, err := ch.QueueDeclare("transaction.created", true, false, false, false, nil)
    if err != nil {
        log.Fatalf("Failed to declare queue: %v", err)
    }

    msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
    if err != nil {
        log.Fatalf("Failed to register consumer: %v", err)
    }

    log.Println("Waiting for transaction messages...")

    for msg := range msgs {
        var tx model.Transaction
        if err := json.Unmarshal(msg.Body, &tx); err != nil {
            log.Printf("Invalid message format: %v", err)
            continue
        }
        go handler.SendNotification(tx)
    }
}
