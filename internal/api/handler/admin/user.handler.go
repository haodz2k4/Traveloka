package admin

import (
	"Traveloka/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	users, err := service.GetAllUsers()
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
