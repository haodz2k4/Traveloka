package admin

import (
	"Traveloka/internal/api/handler/admin"
	"github.com/gin-gonic/gin"
)

func SetupUserRouter(r *gin.RouterGroup) {
	// /admin/users
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", admin.Index)
		userRoutes.GET("/detail/:id", admin.Detail)
		userRoutes.PATCH("/change/status/:id", admin.ChangeStatus)
	}
}
