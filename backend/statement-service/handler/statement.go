package handler

import (
    "encoding/csv"
    "fmt"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
    "statement-service/config"
    "statement-service/model"
    "github.com/jung-kurt/gofpdf"
)

func GetStatements(c *gin.Context) {
    var transactions []model.Transaction
    db := config.GetDB()

    banks := c.QueryArray("bank")
    merchantID := c.Query("merchant_id")

    query := db.Table("transactions").Select("*")
    if len(banks) > 0 {
        query = query.Where("bank IN ?", banks)
    }
    if merchantID != "" {
        query = query.Where("merchant_id = ?", merchantID)
    }

    if err := query.Order("timestamp DESC").Find(&transactions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query transactions"})
        return
    }

    c.JSON(http.StatusOK, transactions)
}

func GetSummary(c *gin.Context) {
    db := config.GetDB()
    merchantID := c.Query("merchant_id")

    if merchantID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "merchant_id is required"})
        return
    }

    type Summary struct {
        Bank   string  `json:"bank"`
        Total  float64 `json:"total_amount"`
        Count  int64   `json:"transaction_count"`
    }

    var summaries []Summary
    if err := db.Table("transactions").
        Select("bank, COUNT(*) as transaction_count, SUM(amount) as total_amount").
        Where("merchant_id = ?", merchantID).
        Group("bank").
        Find(&summaries).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to summarize transactions"})
        return
    }

    c.JSON(http.StatusOK, summaries)
}

func ExportStatements(c *gin.Context) {
    db := config.GetDB()
    merchantID := c.Query("merchant_id")
    format := c.DefaultQuery("format", "csv")

    if merchantID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "merchant_id is required"})
        return
    }

    var transactions []model.Transaction
    if err := db.Table("transactions").
        Where("merchant_id = ?", merchantID).
        Order("timestamp DESC").
        Find(&transactions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to export transactions"})
        return
    }

    if format == "csv" {
        c.Header("Content-Type", "text/csv")
        c.Header("Content-Disposition", `attachment;filename="statement.csv"`)

        writer := csv.NewWriter(c.Writer)
        defer writer.Flush()

        writer.Write([]string{"TransactionID", "Bank", "Amount", "Timestamp"})

        for _, tx := range transactions {
            writer.Write([]string{
                tx.TransactionID,
                tx.Bank,
                fmt.Sprintf("%.2f", tx.Amount),
                tx.Timestamp.Format(time.RFC3339),
            })
        }
        return
    }
    if format == "pdf" {
        pdf := gofpdf.New("P", "mm", "A4", "")
        pdf.AddPage()
    
        // Title
        pdf.SetFont("Arial", "B", 16)
        pdf.CellFormat(0, 10, "Transaction Statement", "", 1, "C", false, 0, "")
    
        pdf.Ln(4)
    
        // Header
        pdf.SetFont("Arial", "B", 11)
        pdf.SetFillColor(200, 220, 255)
        pdf.CellFormat(50, 10, "Transaction ID", "1", 0, "C", true, 0, "")
        pdf.CellFormat(30, 10, "Bank", "1", 0, "C", true, 0, "")
        pdf.CellFormat(30, 10, "Amount", "1", 0, "C", true, 0, "")
        pdf.CellFormat(60, 10, "Timestamp", "1", 1, "C", true, 0, "")
    
        // Body
        pdf.SetFont("Arial", "", 10)
        for _, tx := range transactions {
            pdf.CellFormat(50, 8, truncateID(tx.TransactionID), "1", 0, "L", false, 0, "")
            pdf.CellFormat(30, 8, tx.Bank, "1", 0, "C", false, 0, "")
            pdf.CellFormat(30, 8, fmt.Sprintf("%.2f", tx.Amount), "1", 0, "R", false, 0, "")
            pdf.CellFormat(60, 8, tx.Timestamp.Format("2006-01-02 15:04"), "1", 1, "C", false, 0, "")
        }
    
        // Footer
        pdf.Ln(10)
        pdf.SetFont("Arial", "I", 8)
        pdf.CellFormat(0, 10, fmt.Sprintf("Generated on: %s", time.Now().Format("2006-01-02 15:04")), "", 0, "R", false, 0, "")
    
        c.Header("Content-Type", "application/pdf")
        c.Header("Content-Disposition", `attachment;filename="statement.pdf"`)
        if err := pdf.Output(c.Writer); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "PDF generation failed"})
        }
        return
    }
    
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid export format"})

}

 func truncateID(id string) string {
    if len(id) > 10 {
        return id[:10] + "..."
    }
    return id
}