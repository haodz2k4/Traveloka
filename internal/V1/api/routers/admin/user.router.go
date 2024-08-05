package admin

import (
	"Traveloka/internal/V1/api/handler/admin"
	"github.com/gin-gonic/gin"
)

func SetupUserRouter(r *gin.RouterGroup) {
	// /admin/users
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", admin.Index)
		userRoutes.GET("/detail/:id", admin.Detail)
		userRoutes.PATCH("/change/status/:id", admin.ChangeStatus)
		userRoutes.PATCH("/delete/soft/:id", admin.SoftDelete)
		userRoutes.PATCH("/delete/restore/:id", admin.Restore)
		userRoutes.DELETE("/delete/permantely/:id", admin.DeletePermantely)
		userRoutes.PATCH("/edit/:id", admin.EditUser)
	}
}
