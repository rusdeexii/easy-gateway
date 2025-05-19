// ✅ Middleware + Routing สำหรับ Gateway รองรับ JWT Auth + Rate Limit + CORS ปลอดภัย พร้อม dynamic path (strip prefix)

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"github.com/golang-jwt/jwt/v4"
)

// 🔐 Middleware ตรวจ JWT (จาก cookie หรือ header)
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("access_token")
		if err != nil || tokenStr == "" {
			tokenStr = c.GetHeader("Authorization")
			if strings.HasPrefix(tokenStr, "Bearer ") {
				tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
			}
		}

		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		secret := os.Getenv("JWT_SECRET_KEY")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Next()
	}
}

// ⚙️ Middleware RateLimiter ต่อ IP
var rateLimiters = sync.Map{} // map[ip]*rate.Limiter
func RateLimiterMiddleware(rps float64, burst int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiterInterface, _ := rateLimiters.LoadOrStore(ip, rate.NewLimiter(rate.Limit(rps), burst))
		limiter := limiterInterface.(*rate.Limiter)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			return
		}
		c.Next()
	}
}

// 🌐 Middleware CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// 🔁 Proxy Handler (strip prefix แบบ dynamic)
func proxy(target string, routePrefix string) gin.HandlerFunc {
	remote, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(remote)
	return func(c *gin.Context) {
		// Strip prefix เช่น /api/auth → /login
		prefix := "/api" + routePrefix
		if strings.HasPrefix(c.Request.URL.Path, prefix) {
			c.Request.URL.Path = strings.TrimPrefix(c.Request.URL.Path, prefix)
			if c.Request.URL.Path == "" {
				c.Request.URL.Path = "/"
			}
		}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}


// 🚀 main.go สำหรับ Gateway
func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.SetTrustedProxies(nil) // ✅ ปลอดภัยใน Docker network

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())
	r.Use(RateLimiterMiddleware(5, 10))

	api := r.Group("/api")
	{
		// 🔓 Public routes
		api.Any("/auth/*path", proxy("http://auth_service:8081", "/auth"))
		api.Any("/callback/*path", proxy("http://callback_service:8082", "/callback"))
		api.Any("/notification/*path", proxy("http://notification_service:8083", "/notification"))
		api.Any("/agent/*path", proxy("http://agent_service:8085", "/agent"))

		// 🔐 Protected routes
		authRequired := api.Group("/")
		authRequired.Use(JWTMiddleware())
		{
			authRequired.Any("/merchant/*path", proxy("http://merchant_service:8084", "/merchant"))
			authRequired.Any("/statement/*path", proxy("http://statement_service:8086", "/statement"))
		}
	}

	log.Println("✅ Gateway listening on :8080")
	r.Run(":8080")
}
