package routes

import (
	"github.com/gin-gonic/gin"
	"be-golang/controller"
)

func ProfileRoutes(r *gin.Engine) {
	org := r.Group("api/v1/profile")
	{
		org.POST("", handlers.CreateProfile)
		org.GET("", handlers.GetProfile)
		org.POST("/:id", handlers.DeleteProfile)
	}
}
