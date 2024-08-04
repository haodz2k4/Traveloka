package service

import (
	"Traveloka/internal/models"
	"Traveloka/pkg/config"
	"fmt"
)

type FilterUser struct {
	Status  string
	Keyword string
	Email   string
}
type SortUser struct {
	SortKey   string
	SortValue string
}

func GetAllUsers(filter *FilterUser, sort *SortUser) ([]models.Users, error) {
	//Filter
	db := config.DB
	//status
	if filter.Status != "" {
		db = db.Where("status = ?", filter.Status)
	}
	//find
	if filter.Keyword != "" {
		db = db.Where("last_name like ? ", "%"+filter.Keyword+"%")
	}
	//email
	if filter.Email != "" {
		db = db.Where("email like ? ", "%"+filter.Email+"%")
	}

	//sort
	if sort.SortValue != "" && sort.SortValue != "" {
		string := fmt.Sprintf("%s %s", sort.SortKey, sort.SortValue)
		fmt.Println(string)
		db = db.Order(string)
	}

	var users []models.Users
	if err := db.Find(&users).Error; err != nil {
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
