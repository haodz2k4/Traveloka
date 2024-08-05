package admin

import (
	"Traveloka/helper"
	"Traveloka/internal/V1/service"
	"Traveloka/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// [GET] "/admin/delete"
func Index(c *gin.Context) {
	//filter
	status := c.Query("status")
	keyword := c.Query("keyword")
	email := c.Query("email")
	filter := service.FilterUser{Status: status, Keyword: keyword, Email: email}
	//end filter

	//sort
	sortKey := c.Query("sortKey")
	sortValue := c.Query("sortValue")
	sort := service.SortUser{sortKey, sortValue}
	//end sort

	//Pagination
	pgnt := helper.GetPagination(c)

	users, err := service.GetAllUsers(&filter, &sort, pgnt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
func Detail(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	user, err := service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func ChangeStatus(c *gin.Context) {
	var inp models.Users
	id := c.Param("id")

	if err := c.BindJSON(&inp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user, err := service.ChangeStatus(id, inp.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})

}

func SoftDelete(c *gin.Context) {
	id := c.Param("id")
	user, err := service.SoftDelete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Restore(c *gin.Context) {
	id := c.Param("id")

	user, err := service.Restore(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func DeletePermantely(c *gin.Context) {
	id := c.Param("id")
	err := service.DeletePermantely(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "delete successfull"})
}

func EditUser(c *gin.Context) {
	id := c.Param("id")
	var user models.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	record, err := service.EditUserById(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": record})
}
