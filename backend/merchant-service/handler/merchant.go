package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "merchant-service/config"
    "merchant-service/model"
)

func CreateMerchant(c *gin.Context) {
    var merchant model.Merchant
    if err := c.ShouldBindJSON(&merchant); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.GetDB().Create(&merchant).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create merchant"})
        return
    }

    c.JSON(http.StatusOK, merchant)
}

func GetWebhookURL(c *gin.Context) {
    var merchant model.Merchant
    id := c.Param("id")

    if err := config.GetDB().First(&merchant, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Merchant not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"webhook_url": merchant.WebhookURL})
}
