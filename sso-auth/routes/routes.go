package routes

import (
	"sso-auth/controllers"
	"sso-auth/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/google", controllers.GoogleLogin)
			auth.GET("/google/callback", controllers.GoogleCallback)
			auth.GET("/me", middleware.AuthMiddleware(), controllers.GetCurrentUser)
			auth.GET("/verify", middleware.AuthMiddleware(), controllers.VerifyToken)
			auth.POST("/logout", middleware.AuthMiddleware(), controllers.Logout)
		}
	}
}
