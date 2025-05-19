package config

import (
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
)

// การตั้งค่าที่เกี่ยวข้องกับ RabbitMQ
var RabbitMQConfig struct {
    Host     string
    Port     string
    User     string
    Password string
}

// การตั้งค่าที่เกี่ยวข้องกับการส่งอีเมล (SMTP)
var SMTPConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    From     string
    To       string
}

// LoadEnv จะโหลดค่าการตั้งค่าจากไฟล์ .env
func LoadEnv() {
    if err := godotenv.Load(".env.dev"); err != nil {
        log.Println("No .env.dev file found")
    }

    // โหลดค่าการตั้งค่าสำหรับ RabbitMQ
    RabbitMQConfig.Host = os.Getenv("RABBITMQ_HOST")
    RabbitMQConfig.Port = os.Getenv("RABBITMQ_PORT")
    RabbitMQConfig.User = os.Getenv("RABBITMQ_USER")
    RabbitMQConfig.Password = os.Getenv("RABBITMQ_PASS")

    // โหลดค่าการตั้งค่าสำหรับ SMTP
    SMTPConfig.Host = os.Getenv("SMTP_HOST")
    SMTPConfig.Port = os.Getenv("SMTP_PORT")
    SMTPConfig.User = os.Getenv("SMTP_USER")
    SMTPConfig.Password = os.Getenv("SMTP_PASS")
    SMTPConfig.From = os.Getenv("NOTIFY_EMAIL_FROM")
    SMTPConfig.To = os.Getenv("NOTIFY_EMAIL_TO")

    log.Printf("Loaded Environment Variables: RabbitMQ Host=%s, SMTP Host=%s", RabbitMQConfig.Host, SMTPConfig.Host)
}

// RabbitMQ URI ใช้สำหรับการเชื่อมต่อ RabbitMQ
func GetRabbitMQURI() string {
    return fmt.Sprintf("amqp://%s:%s@%s:%s/",
        RabbitMQConfig.User,
        RabbitMQConfig.Password,
        RabbitMQConfig.Host,
        RabbitMQConfig.Port)
}
