package main

import (
	"Traveloka/internal/V1/api/routers/admin"
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
	admin.SetupAdminRouter(r.Group("/api/v1"))
	r.Run()
}
