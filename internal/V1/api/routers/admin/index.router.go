package admin

import "github.com/gin-gonic/gin"

func SetupAdminRouter(r *gin.RouterGroup) {

	adminRouter := r.Group("/admin")
	{
		SetupUserRouter(adminRouter)
	}

}
