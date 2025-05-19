// merchant-service/model/merchant.go
package model

import (
    "time"

    "github.com/google/uuid"
)

type Merchant struct {
    ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
    Name       string    `gorm:"not null" json:"name"`
    WebhookURL string    `gorm:"not null;index" json:"webhook_url"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}
