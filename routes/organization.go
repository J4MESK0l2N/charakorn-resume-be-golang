package routes

import (
	"github.com/gin-gonic/gin"
	"be-golang/controller"
)

func OrganizationRoutes(r *gin.Engine) {
	org := r.Group("api/v1/organizations")
	{
		org.POST("", handlers.CreateOrganization)
		org.GET("", handlers.GetAllOrganizations)
		org.GET("/:id", handlers.GetOrganization)
		org.PUT("/:id", handlers.UpdateOrganization)
		org.DELETE("/:id", handlers.DeleteOrganization)
	}
}
