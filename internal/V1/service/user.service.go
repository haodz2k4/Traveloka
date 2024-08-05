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

func Restore(id string) (*models.Users, error) {
	var user models.Users
	db := config.DB
	result := db.Model(&user).Where("user_id = ?", id).Update("deleted", false)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("Không dòng nào bị ảnh hưởng")
	}
	if err := db.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func DeletePermantely(id string) error {
	var user models.Users
	db := config.DB
	result := db.Where("user_id = ? AND deleted = ?", id, true).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Người dùng này chưa bị xóa mềm hoặc đã bị xóa vĩnh viễn")
	}
	return nil
}

func EditUserById(id string, body models.Users) (*models.Users, error) {

	var user models.Users
	db := config.DB

	result := db.Model(&user).Where("user_id = ? AND deleted = ?", id, false).Updates(body)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("Không dong nào được cập nhật")
	}
	if err := db.Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func ChangeMultiDelete(ids []string) error {

	var user models.Users
	db := config.DB
	result := db.Model(&user).Where("user_id IN ?", ids).Update("deleted", true)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Không có dòng nào được cập nhật")
	}

	return nil
}