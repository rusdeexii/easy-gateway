package router

import (
    "github.com/gin-gonic/gin"
    "callback-service/handler"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/callback/:merchantId/:bank", handler.HandleCallback)
    return r
}
