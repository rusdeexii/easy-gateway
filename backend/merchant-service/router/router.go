package router

import (
    "github.com/gin-gonic/gin"
    "merchant-service/handler"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/merchants", handler.CreateMerchant)
    r.GET("/merchants/:id/webhook-url", handler.GetWebhookURL)

    return r
}
