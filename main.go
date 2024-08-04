package main

import (
	"Traveloka/internal/api/routers/admin"
	"Traveloka/pkg/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := config.GetConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return
	}

	r := gin.Default()
	admin.SetupAdminRouter(r.Group("/"))
	r.Run()
}
