package routes

import (
	"github.com/gin-gonic/gin"
	"be-golang/controller"
)

func ProjectRoutes(r *gin.Engine) {
	org := r.Group("api/v1/projects")
	{
		org.POST("", handlers.CreateProject)
		org.DELETE("/:id", handlers.DeleteProject)
	}
}
