// // transaction-service/model/transaction.go
package model

import "time"

type Transaction struct {
    TransactionID string    `json:"transaction_id"`
    Bank          string    `json:"bank"`
    Amount        float64   `json:"amount"`
    Timestamp     time.Time `json:"timestamp"`
    MerchantID    string    `json:"merchant_id"`
}
