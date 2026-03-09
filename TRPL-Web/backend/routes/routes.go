package routes

import (
	"trpl-web-backend/controllers"
	"trpl-web-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Project routes
		projects := api.Group("/projects")
		{
			projects.GET("", controllers.GetProjects)            // Public
			projects.GET("/:id", controllers.GetProjectByID)     // Public
			projects.GET("/stats", controllers.GetStats)         // Public
			
			// Protected routes
			projects.POST("", middleware.AuthMiddleware(), controllers.CreateProject)
			projects.PUT("/:id", middleware.AuthMiddleware(), controllers.UpdateProject)
			projects.DELETE("/:id", middleware.AuthMiddleware(), controllers.DeleteProject)
		}

		// My projects (Mahasiswa only)
		api.GET("/my-projects", middleware.AuthMiddleware(), controllers.GetMyProjects)
	}
}
