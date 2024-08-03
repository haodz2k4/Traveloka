package main

import (
	"Traveloka/internal/api/routers/admin"
	"Traveloka/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	connect := config.GetConnection()
	r := gin.Default()
	fmt.Println("connect is :", connect)
	admin.SetupAdminRouter(r.Group("/"))

	r.Run()
}
