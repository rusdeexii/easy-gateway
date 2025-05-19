package router

import (
    "github.com/gin-gonic/gin"
    "statement-service/handler"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/api/statements", handler.GetStatements)
    r.GET("/api/statements/summary", handler.GetSummary)
    r.GET("/api/statements/export", handler.ExportStatements)
    return r
}
