package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

// RabbitMQ connection configuration
type RabbitMQConfig struct {
    Host     string
    Port     string
    User     string
    Password string
}

// LoadEnv ใช้โหลดค่าการตั้งค่าจากไฟล์ .env
func LoadEnv() {
    if err := godotenv.Load(".env.dev"); err != nil {
        log.Println("No .env.dev file found")
    }
}

// GetRabbitMQConfig ใช้ดึงค่าการตั้งค่า RabbitMQ
func GetRabbitMQConfig() *RabbitMQConfig {
    return &RabbitMQConfig{
        Host:     os.Getenv("RABBITMQ_HOST"),
        Port:     os.Getenv("RABBITMQ_PORT"),
        User:     os.Getenv("RABBITMQ_USER"),
        Password: os.Getenv("RABBITMQ_PASS"),
    }
}

