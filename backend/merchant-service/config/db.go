package config

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "merchant-service/model"
)

var DB *gorm.DB

// ConnectDB connects to the database and runs AutoMigrate
func ConnectDB() {
    // ตรวจสอบค่าของตัวแปรสภาพแวดล้อมก่อน
    requiredEnvVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
    for _, envVar := range requiredEnvVars {
        if os.Getenv(envVar) == "" {
            log.Fatalf("Environment variable %s is not set", envVar)
        }
    }

    // สร้าง DSN สำหรับการเชื่อมต่อ
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Perform AutoMigrate for Merchant model
    if err := DB.AutoMigrate(&model.Merchant{}); err != nil {
        log.Fatalf("Auto migration failed: %v", err)
    }

    log.Println("Successfully connected to the database and migrated the Merchant model.")
}

// GetDB returns the database connection instance
func GetDB() *gorm.DB {
    return DB
}
