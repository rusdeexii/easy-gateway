package handler

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "net/smtp"
    "notification-service/model"
    "os"
    "strings"
    "text/template"
    "time"
)

const emailTemplate = `
📣 Transaction Notification
===========================

Merchant ID : {{.MerchantID}}
Transaction : {{.TransactionID}}
Bank        : {{.Bank}}
Amount      : {{printf "%.2f" .Amount}} THB
Timestamp   : {{.Timestamp.Format "2006-01-02 15:04:05"}}

---------------------------
📌 Sent by Easy-Gateway System
`

// ✅ เรียกจาก consumer.go
func SendNotification(tx model.Transaction) {
    log.Println("🔔 Dispatching notification for tx:", tx.TransactionID)

    go SendWebhook(tx)
    go SendEmail(tx)
}

func SendWebhook(tx model.Transaction) {
    webhookURL := os.Getenv("DEFAULT_WEBHOOK_URL")

    payload, _ := json.Marshal(map[string]interface{}{
        "type":           "transaction",
        "transaction_id": tx.TransactionID,
        "amount":         tx.Amount,
        "bank":           tx.Bank,
        "timestamp":      tx.Timestamp,
        "merchant_id":    tx.MerchantID,
    })

    req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payload))
    if err != nil {
        log.Printf("[ERROR] Create webhook request: %v", err)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{Timeout: 5 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("[ERROR] Send webhook: %v", err)
        return
    }
    defer resp.Body.Close()

    log.Printf("[WEBHOOK] Sent to %s | Status=%d", webhookURL, resp.StatusCode)
}

func SendEmail(tx model.Transaction) {
    from := os.Getenv("NOTIFY_EMAIL_FROM")
    to := os.Getenv("NOTIFY_EMAIL_TO")
    smtpHost := os.Getenv("SMTP_HOST")
    smtpPort := os.Getenv("SMTP_PORT")
    smtpUser := os.Getenv("SMTP_USER")
    smtpPass := os.Getenv("SMTP_PASS")

    subject := fmt.Sprintf("💰 Transaction Alert: %s", tx.TransactionID)

    tmpl, err := template.New("email").Parse(emailTemplate)
    if err != nil {
        log.Printf("[ERROR] parse template: %v", err)
        return
    }

    var bodyBuf bytes.Buffer
    if err := tmpl.Execute(&bodyBuf, tx); err != nil {
        log.Printf("[ERROR] template execute: %v", err)
        return
    }

    message := strings.Join([]string{
        fmt.Sprintf("From: %s", from),
        fmt.Sprintf("To: %s", to),
        fmt.Sprintf("Subject: %s", subject),
        "MIME-Version: 1.0",
        "Content-Type: text/plain; charset=UTF-8",
        "",
        bodyBuf.String(),
    }, "\r\n")

    maxRetries := 3
    for attempt := 1; attempt <= maxRetries; attempt++ {
        auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
        err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
        if err != nil {
            log.Printf("[RETRY-%d] send email failed: %v", attempt, err)
            time.Sleep(time.Duration(attempt*2) * time.Second)
            continue
        }

        log.Printf("[EMAIL] Sent to %s successfully (attempt %d)", to, attempt)
        return
    }

    log.Printf("[FAIL] All retry attempts failed for email to %s", to)
}
