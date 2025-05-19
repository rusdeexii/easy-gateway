package model

import "gorm.io/gorm"

type Merchant struct {
    gorm.Model
    Name       string `json:"name"`
    WebhookURL string `json:"webhook_url"`
}
