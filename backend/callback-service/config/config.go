package config

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var db *gorm.DB

// LoadEnv loads environment variables and initializes the database connection
func LoadEnv() {
    if err := godotenv.Load(".env.dev"); err != nil {
        log.Println("No .env.dev file found, relying on environment variables.")
    }

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    log.Printf("Loaded Environment Variables: DB_HOST=%s, DB_PORT=%s, DB_USER=%s, DB_PASSWORD=%s, DB_NAME=%s", dbHost, dbPort, dbUser, dbPassword, dbName)

    dsn := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=disable",
        dbUser, dbPassword, dbHost, dbPort, dbName,
    )

    log.Printf("DSN: %s", dsn)

    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
}

// GetDB returns the GORM database instance
func GetDB() *gorm.DB {
    return db
}
