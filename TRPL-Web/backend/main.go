package main

import (
	"log"
	"os"
	"trpl-web-backend/config"
	"trpl-web-backend/controllers"
	"trpl-web-backend/models"
	"trpl-web-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Connect to database
	config.ConnectDatabase()

	// Auto migrate models
	db := config.GetDB()
	if err := db.AutoMigrate(
		&models.User{},
		&models.Mahasiswa{},
		&models.Dosen{},
		&models.ProgramStudi{},
		&models.TahunAjaran{},
		&models.Project{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize Google OAuth
	controllers.InitGoogleOAuth()

	// Setup Gin
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Setup routes
	routes.SetupRoutes(r)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
