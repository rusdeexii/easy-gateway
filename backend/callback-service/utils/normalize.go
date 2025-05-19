package utils

import (
    "callback-service/model"
    "errors"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

func NormalizePayload(bank string, c *gin.Context) (model.Transaction, error) {
    var tx model.Transaction
    merchantId := c.Param("merchantId")

    switch bank {
    case "SCB":
        var body struct {
            TxnID     string  `json:"txn_id"`
            Amount    float64 `json:"amount"`
            Timestamp string  `json:"timestamp"`
        }
        if err := c.ShouldBindJSON(&body); err != nil {
            return tx, err
        }

        t, err := time.Parse(time.RFC3339, body.Timestamp)
        if err != nil {
            return tx, err
        }

        tx = model.Transaction{
            ID:            uuid.New(),
            Bank:          "SCB",
            TransactionID: body.TxnID,
            Amount:        body.Amount,
            Timestamp:     t,
            MerchantID:    merchantId,
        }

    case "KBANK":
        var body struct {
            TransactionID string `json:"transaction_id"`
            Total         string `json:"total"`
            Time          string `json:"time"`
        }
        if err := c.ShouldBindJSON(&body); err != nil {
            return tx, err
        }

        amt, err := strconv.ParseFloat(body.Total, 64)
        if err != nil {
            return tx, err
        }

        t, err := time.Parse("2006-01-02 15:04:05", body.Time)
        if err != nil {
            return tx, err
        }

        tx = model.Transaction{
            ID:            uuid.New(),
            Bank:          "KBANK",
            TransactionID: body.TransactionID,
            Amount:        amt,
            Timestamp:     t,
            MerchantID:    merchantId,
        }

    case "TTB":
        var body struct {
            Ref string  `json:"ref"`
            Amt float64 `json:"amt"`
            Ts  int64   `json:"ts"`
        }
        if err := c.ShouldBindJSON(&body); err != nil {
            return tx, err
        }

        t := time.Unix(body.Ts, 0)

        tx = model.Transaction{
            ID:            uuid.New(),
            Bank:          "TTB",
            TransactionID: body.Ref,
            Amount:        body.Amt,
            Timestamp:     t,
            MerchantID:    merchantId,
        }

    default:
        return tx, errors.New("unsupported bank")
    }

    return tx, nil
}
