// cmd/main.go
package main

import (
	"auth-service/config"
	"auth-service/controller"
	"auth-service/model"
	"auth-service/repository"
	"auth-service/service"
	"auth-service/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()
	config.DB.AutoMigrate(&model.User{}, &model.RefreshToken{})

	userRepo := repository.NewUserRepository()
	authService := service.NewAuthService(userRepo)

	r := gin.Default()
	r.Use(middleware.SecureHeaders()) // ✅ ปลอดภัยยิ่งขึ้น

	controller.RegisterAuthRoutes(r, authService)
	r.Run(":" + os.Getenv("PORT"))
}