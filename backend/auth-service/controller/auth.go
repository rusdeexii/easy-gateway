// contr
package controller

import (
	"auth-service/model"
	"auth-service/service"
	"auth-service/middleware"
	"auth-service/utils"
	"net/http"
	"time"
	"os"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (a *AuthController) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, err := a.AuthService.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user.SafeUser())
}

func (a *AuthController) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	user, access, refresh, err := a.AuthService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set tokens as HttpOnly cookies
	c.SetCookie("access_token", access, 900, "/", "", true, true)      // 15 min
	c.SetCookie("refresh_token", refresh, 604800, "/", "", true, true) // 7 days

	c.JSON(http.StatusOK, gin.H{
		"user":       user.SafeUser(),
		"expires_at": time.Now().Add(15 * time.Minute),
	})
}

func (a *AuthController) Refresh(c *gin.Context) {
	refreshTokenStr, err := c.Cookie("refresh_token")
	if err != nil || refreshTokenStr == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no refresh token"})
		return
	}

	claims, err := utils.ParseToken(refreshTokenStr, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}

	// ตรวจสอบและแปลง user_id
	userIDRaw, ok := claims["user_id"]
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing user_id in token"})
		return
	}
	userID, ok := userIDRaw.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id format"})
		return
	}

	// ตรวจสอบและแปลง role
	roleRaw, ok := claims["role"]
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing role in token"})
		return
	}
	role, ok := roleRaw.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid role format"})
		return
	}

	// ตรวจสอบใน DB ว่ามี refresh token นี้จริง และยังไม่ revoke
	_, err = a.AuthService.ValidateRefreshToken(refreshTokenStr, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// ✅ สร้าง access token ใหม่
	access, _, _ := utils.GenerateTokens(userID, role)

	// ✅ set access_token เป็น cookie
	c.SetCookie("access_token", access, 900, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "refreshed"})
}


func (a *AuthController) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", true, true)
	c.SetCookie("refresh_token", "", -1, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

func (a *AuthController) LogoutAll(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err == nil && refreshToken != "" {
		_ = a.AuthService.RevokeAllRefreshTokens(refreshToken)
	}

	c.SetCookie("access_token", "", -1, "/", "", true, true)
	c.SetCookie("refresh_token", "", -1, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out everywhere"})
}

// GET /users
func (a *AuthController) ListUsers(c *gin.Context) {
	users, err := a.AuthService.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}
	var result []map[string]interface{}
	for _, user := range users {
		result = append(result, user.SafeUser())
	}
	c.JSON(http.StatusOK, result)
}

func (a *AuthController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var update model.UpdateUserRequest
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	updatedUser, err := a.AuthService.UpdateUser(id, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser.SafeUser())
}


// DELETE /users/:id
func (a *AuthController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := a.AuthService.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}


func (a *AuthController) Me(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, err := a.AuthService.GetUserByUUID(userID.(string))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user.SafeUser()})
}

func RegisterAuthRoutes(r *gin.Engine, authService service.AuthService) {
	auth := NewAuthController(authService)

	r.POST("/login", auth.Login)
	r.POST("/register", auth.Register)
	r.POST("/refresh", auth.Refresh)
	r.POST("/logout", auth.Logout)
	r.POST("/logout_all", auth.LogoutAll)
	r.GET("/me", middleware.AuthMiddleware(), auth.Me)
	r.GET("/users", middleware.AuthMiddleware(), middleware.RequireRole("superadmin"), auth.ListUsers)
	r.PUT("/users/:id", middleware.AuthMiddleware(), middleware.RequireRole("superadmin"), auth.UpdateUser)
	r.DELETE("/users/:id", middleware.AuthMiddleware(), middleware.RequireRole("superadmin"), auth.DeleteUser)
	r.POST("/users", middleware.AuthMiddleware(), middleware.RequireRole("superadmin"), auth.Register)


}

