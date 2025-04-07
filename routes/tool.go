package routes

import (
	"github.com/gin-gonic/gin"
	"be-golang/controller"
)

func ToolRoutes(r *gin.Engine) {
	org := r.Group("api/v1/tools")
	{
		org.POST("", handlers.CreateTool)
		org.DELETE("/:id", handlers.DeleteTool)
		org.GET("", handlers.GetTools)
	}
}
