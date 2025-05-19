package config

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// DB ใช้เก็บการเชื่อมต่อฐานข้อมูล
var DB *gorm.DB

// LoadEnv ใช้โหลดค่าการตั้งค่าจากไฟล์ .env
func LoadEnv() {
    if err := godotenv.Load(".env.dev"); err != nil {
        log.Println("No .env.dev file found")
    }
}

// ConnectDB เชื่อมต่อกับฐานข้อมูล
func ConnectDB() {
    LoadEnv() // โหลดไฟล์ .env ก่อน

    // สร้าง DSN (Data Source Name) สำหรับเชื่อมต่อฐานข้อมูล
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"))

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
}

// GetDB ใช้ดึงการเชื่อมต่อฐานข้อมูล
func GetDB() *gorm.DB {
    return DB
}
