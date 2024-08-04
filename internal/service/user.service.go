package service

import (
	"Traveloka/internal/models"
	"Traveloka/pkg/config"
)

func GetAllUsers() ([]models.Users, error) {
	var users []models.Users
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func GetUserById(id string) (*models.Users, error) {
	var user models.Users

	if err := config.DB.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
