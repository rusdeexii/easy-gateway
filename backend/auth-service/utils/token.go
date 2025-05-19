// utils/token.go - Update with specific token generation methods
package utils

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateTokens(userId, userRole string) (string, string, time.Time) {
	secret := os.Getenv("JWT_SECRET_KEY")

	atClaims := jwt.MapClaims{
		"user_id": userId,
		"role":    userRole, // ✅ สำคัญ
		"type":    "access",
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	rtExp := time.Now().Add(7 * 24 * time.Hour)
	rtClaims := jwt.MapClaims{
		"user_id": userId,
		"role":    userRole, // ✅ ต้องใส่ด้วย
		"type":    "refresh",
		"exp":     rtExp.Unix(),
	}

	access, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims).SignedString([]byte(secret))
	refresh, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims).SignedString([]byte(secret))

	return access, refresh, rtExp
}




func GenerateAccessToken(userId string) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	atClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
		"type":    "access",
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	return token.SignedString([]byte(secret))
}

func GenerateRefreshToken(userId string) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	rtClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
		"type":    "refresh",
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenStr string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
