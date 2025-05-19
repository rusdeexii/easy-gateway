// repository/user.go

package repository

import (
	"auth-service/config"
	"auth-service/model"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByUsername(username string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
	FindByUserId(userId string) (*model.User, error) // ✅ เพิ่มตรงนี้
	FindAll() ([]*model.User, error)
	Update(user *model.User) error
	DeleteByUserId(userId string) error
}

type userRepo struct{}


func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (r *userRepo) Create(user *model.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepo) FindByUsername(username string) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByID(id uint) (*model.User, error) {
	var user model.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindByUserId(userId string) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("user_id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) FindAll() ([]*model.User, error) {
	var users []*model.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) Update(user *model.User) error {
	return config.DB.Save(user).Error
}

func (r *userRepo) DeleteByUserId(userId string) error {
	return config.DB.Where("user_id = ?", userId).Delete(&model.User{}).Error
}

