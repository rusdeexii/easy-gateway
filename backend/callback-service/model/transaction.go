package model

import (
    "time"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

// Transaction represents a transaction record
type Transaction struct {
    ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    TransactionID string    `gorm:"not null;unique"`
    Bank          string    `gorm:"not null"`
    Amount        float64   `gorm:"type:decimal(10,2);not null"`
    Timestamp     time.Time `gorm:"not null"`
    MerchantID    string
    CreatedAt     time.Time
}

// SaveTransaction saves a transaction record to the database
func SaveTransaction(db *gorm.DB, tx Transaction) error {
    return db.Create(&tx).Error
}
