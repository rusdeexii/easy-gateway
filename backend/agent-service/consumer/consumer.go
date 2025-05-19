package consumer

import (
    "encoding/json"
    "fmt"
    "log"
    "agent-service/handler"
    "agent-service/model"
    "agent-service/config" // import config package
    "github.com/streadway/amqp"
)

func StartConsumer() {
    // โหลดค่าการตั้งค่าจาก config
    rabbitConfig := config.GetRabbitMQConfig()

    // สร้าง URI สำหรับ RabbitMQ
    uri := fmt.Sprintf("amqp://%s:%s@%s:%s/",
        rabbitConfig.User,
        rabbitConfig.Password,
        rabbitConfig.Host,
        rabbitConfig.Port,
    )

    // เชื่อมต่อกับ RabbitMQ
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

    // สร้างคิวที่ใช้สำหรับการดึงข้อความ
    q, err := ch.QueueDeclare("transaction.created", true, false, false, false, nil)
    if err != nil {
        log.Fatalf("Failed to declare queue: %v", err)
    }

    // ลงทะเบียน consumer สำหรับคิวที่ประกาศ
    msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
    if err != nil {
        log.Fatalf("Failed to register consumer: %v", err)
    }

    log.Println("Waiting for messages from queue...")

    // รับและประมวลผลข้อความจากคิว
    for msg := range msgs {
        var tx model.Transaction
        if err := json.Unmarshal(msg.Body, &tx); err != nil {
            log.Printf("Invalid message format: %v", err)
            continue
        }
        go handler.SendWebhook(tx)
    }
}
