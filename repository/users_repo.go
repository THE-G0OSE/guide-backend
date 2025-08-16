package repository

import (
	"github.com/THE-G0OSE/guide-backend/models"
	"gorm.io/gorm"
)

type UserRepo struct{ DB *gorm.DB }

func (r *UserRepo) Find(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepo) Login(password string, username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	return &user, err
}

func (r *UserRepo) Create(user *models.User) error {
	err := r.DB.Create(&user).Error
	return err
}
