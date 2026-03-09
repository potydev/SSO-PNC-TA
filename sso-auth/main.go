package main

import (
	"log"
	"os"
	"sso-auth/config"
	"sso-auth/controllers"
	"sso-auth/models"
	"sso-auth/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	config.ConnectDatabase()

	db := config.GetDB()
	if err := db.AutoMigrate(
		&models.User{},
		&models.Mahasiswa{},
		&models.Dosen{},
		&models.ProgramStudi{},
		&models.TahunAjaran{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	controllers.InitGoogleOAuth()

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL"), os.Getenv("TRPL_BACKEND_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(r)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "sso-auth"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Printf("SSO Auth server starting on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
