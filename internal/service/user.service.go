package service

import (
	"Traveloka/helper"
	"Traveloka/internal/models"
	"Traveloka/pkg/config"
	"errors"
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

func GetAllUsers(filter *FilterUser, sort *SortUser, pgnt *helper.Pagination) ([]models.Users, error) {
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
	//pagination
	db = db.Limit(pgnt.Limit)
	db = db.Offset(pgnt.Skip)
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

func ChangeStatus(id string, status string) (*models.Users, error) {
	var user models.Users
	db := config.DB

	if err := db.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&user).Update("status", status).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func SoftDelete(id string) (*models.Users, error) {

	var user models.Users
	db := config.DB
	result := db.Model(&user).Where("user_id = ?", id).Update("deleted", true)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("Không dòng náo bị ảnh hưởng")
	}
	if err := db.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
