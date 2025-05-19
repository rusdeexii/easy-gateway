// model/user.go
package model

import (
	"errors"
	"time"
)

const (
	RoleAdmin      = "admin"
	RoleSuperAdmin = "superadmin"
)

type User struct {
	Id        uint      `gorm:"primaryKey"`
	UserId    string    `gorm:"unique;not null"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"not null"`
	IsActive  bool      `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RefreshToken struct {
	Id        uint      `gorm:"primaryKey"`
	UserId    uint      `gorm:"not null"`
	Token     string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
	IsRevoked bool      `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
	Email    string `json:"email"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *RegisterRequest) Validate() error {
	if r.Username == "" || r.Password == "" || len(r.Password) < 6 {
		return errors.New("username and password required, password must be at least 6 characters")
	}
	if r.Role != RoleAdmin && r.Role != RoleSuperAdmin {
		return errors.New("invalid role")
	}
	return nil
}



func (u *User) SafeUser() map[string]interface{} {
	return map[string]interface{}{
		"id":         u.Id,
		"user_id":    u.UserId,
		"username":   u.Username,
		"role":       u.Role,
		"is_active":  u.IsActive,
		"created_at": u.CreatedAt,
	}
}
