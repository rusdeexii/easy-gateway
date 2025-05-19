package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "callback-service/config"
    "callback-service/model"
    "callback-service/utils"
    "callback-service/queue" // 👈 alias ให้ตรง
)

func HandleCallback(c *gin.Context) {
    bank := c.Param("bank")
    merchantId := c.Param("merchantId")

    tx, err := utils.NormalizePayload(bank, c)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    tx.MerchantID = merchantId

    // 🔧 ใช้ config.GetDB() แล้วส่งเข้า model
    if err := model.SaveTransaction(config.GetDB(), tx); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save transaction"})
        return
    }

    if err := queue.PublishTransaction(tx); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish transaction"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"status": "received"})
}
