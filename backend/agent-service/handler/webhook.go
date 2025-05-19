package handler

import (
    "agent-service/model"
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "time"
)

func getWebhookURL(merchantID string) (string, error) {
    baseURL := os.Getenv("MERCHANT_SERVICE_URL")
    fullURL := fmt.Sprintf("%s/merchants/%s/webhook-url", baseURL, merchantID)

    client := &http.Client{Timeout: 5 * time.Second}
    resp, err := client.Get(fullURL)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("merchant-service returned status %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var res map[string]string
    if err := json.Unmarshal(body, &res); err != nil {
        return "", err
    }

    return res["webhook_url"], nil
}

func SendWebhook(tx model.Transaction) {
    webhookURL, err := getWebhookURL(tx.MerchantID)
    if err != nil {
        log.Printf("[ERROR] getWebhookURL failed | merchantId=%s | err=%v", tx.MerchantID, err)
        return
    }

    payload, err := json.Marshal(map[string]interface{}{
        "bank":           tx.Bank,
        "transaction_id": tx.TransactionID,
        "amount":         tx.Amount,
        "timestamp":      tx.Timestamp,
    })
    if err != nil {
        log.Printf("[ERROR] marshal payload failed | txID=%s | err=%v", tx.TransactionID, err)
        return
    }

    client := &http.Client{Timeout: 5 * time.Second}

    maxRetries := 3
    for attempt := 1; attempt <= maxRetries; attempt++ {
        req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payload))
        req.Header.Set("Content-Type", "application/json")

        if err != nil {
            log.Printf("[ERROR] create request failed | attempt=%d | err=%v", attempt, err)
            continue
        }

        resp, err := client.Do(req)
        if err != nil || resp.StatusCode >= 400 {
            msg := "[ERROR] webhook failed"
            if err == nil {
                msg += fmt.Sprintf(" | status=%d", resp.StatusCode)
            }
            log.Printf("%s | merchantId=%s | attempt=%d | err=%v", msg, tx.MerchantID, attempt, err)

            time.Sleep(time.Duration(attempt*2) * time.Second) // simple backoff
            continue
        }

        log.Printf("[SUCCESS] webhook sent | merchantId=%s | txID=%s | status=%d", tx.MerchantID, tx.TransactionID, resp.StatusCode)
        resp.Body.Close()
        return
    }

    log.Printf("[FAIL] all attempts failed | txID=%s | merchantId=%s", tx.TransactionID, tx.MerchantID)
}
