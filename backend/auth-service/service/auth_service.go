// service/auth_service.go
package service

import (
	"auth-service/model"
	"auth-service/repository"
	"auth-service/utils"
	"auth-service/config"
	"errors"
	"time"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(req model.RegisterRequest) (*model.User, error)
	Login(req model.LoginRequest) (*model.User, string, string, error)
	GetUserByUUID(uuid string) (*model.User, error)
	RevokeAllRefreshTokens(refreshToken string) error
	ValidateRefreshToken(token, userID string) (*model.RefreshToken, error)
	SaveRefreshToken(user *model.User, token string, expires time.Time) error
	ListUsers() ([]*model.User, error)
	UpdateUser(userID string, data model.UpdateUserRequest) (*model.User, error)
	DeleteUser(userID string) error
	


}


type authService struct {
	repo repository.UserRepository
}


func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Register(req model.RegisterRequest) (*model.User, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	user := &model.User{
		UserId:   utils.GenerateUUID(),
		Username: req.Username,
		Password: string(hashed),
		Role:     req.Role,
		IsActive: true,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *authService) RevokeAllRefreshTokens(refreshToken string) error {
	var rt model.RefreshToken
	err := config.DB.Where("token = ?", refreshToken).First(&rt).Error
	if err != nil {
		return nil // silent fail
	}
	return config.DB.Model(&model.RefreshToken{}).
		Where("user_id = ?", rt.UserId).
		Update("is_revoked", true).Error
}

func (s *authService) Login(req model.LoginRequest) (*model.User, string, string, error) {
	user, err := s.repo.FindByUsername(req.Username)
	if err != nil || user == nil {
		return nil, "", "", errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, "", "", errors.New("invalid username or password")
	}

	// ✅ generate access + refresh + refresh_exp
	access, refresh, exp := utils.GenerateTokens(user.UserId, user.Role)

	// ✅ save refresh token to DB
	if err := s.SaveRefreshToken(user, refresh, exp); err != nil {
		return nil, "", "", errors.New("failed to save refresh token")
	}

	return user, access, refresh, nil
}


func (s *authService) GetUserByUUID(uuid string) (*model.User, error) {
	user, err := s.repo.FindByUserId(uuid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *authService) SaveRefreshToken(user *model.User, token string, expires time.Time) error {
	rt := model.RefreshToken{
		UserId:    user.Id,
		Token:     token,
		ExpiresAt: expires,
		IsRevoked: false,
	}
	return config.DB.Create(&rt).Error
}


func (s *authService) ValidateRefreshToken(token, userID string) (*model.RefreshToken, error) {
	var rt model.RefreshToken
	err := config.DB.Where("token = ? AND is_revoked = false", token).First(&rt).Error
	if err != nil {
		return nil, errors.New("refresh token not found or revoked")
	}

	if rt.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("refresh token expired")
	}

	// optional: ตรวจสอบว่า userID ตรงกัน
	user, _ := s.repo.FindByID(rt.UserId)
	if user.UserId != userID {
		return nil, errors.New("token-user mismatch")
	}

	return &rt, nil
}


func (s *authService) ListUsers() ([]*model.User, error) {
	return s.repo.FindAll()
}

func (s *authService) UpdateUser(userID string, data model.UpdateUserRequest) (*model.User, error) {
	user, err := s.repo.FindByUserId(userID)
	if err != nil {
		return nil, err
	}
	if data.Username != "" {
		user.Username = data.Username
	}
	if data.IsActive {
		user.IsActive = data.IsActive
	}
	if data.Role != "" {
		user.Role = data.Role
	}
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *authService) DeleteUser(userID string) error {
	return s.repo.DeleteByUserId(userID)
}
