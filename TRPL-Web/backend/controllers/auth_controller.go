package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"
	"trpl-web-backend/config"
	"trpl-web-backend/models"
	"trpl-web-backend/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config

func InitGoogleOAuth() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

// GET /api/auth/google
func GoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.JSON(http.StatusOK, gin.H{"url": url})
}

// GET /api/auth/google/callback
func GoogleCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
		return
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	defer resp.Body.Close()

	var userInfo struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		Picture       string `json:"picture"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user info"})
		return
	}

	// Validate domain
	allowedDomain := os.Getenv("ALLOWED_DOMAIN")
	if !strings.HasSuffix(userInfo.Email, "@"+allowedDomain) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only @" + allowedDomain + " emails are allowed"})
		return
	}

	// Find or create user
	db := config.GetDB()
	var user models.User
	result := db.Where("email = ?", userInfo.Email).First(&user)

	now := time.Now()
	if result.RowsAffected == 0 {
		// Create new user - default role Mahasiswa
		user = models.User{
			Email:           userInfo.Email,
			Name:            userInfo.Name,
			GoogleID:        userInfo.ID,
			Role:            "Mahasiswa",
			EmailVerifiedAt: &now,
			FotoProfile:     &userInfo.Picture,
		}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
	} else {
		// Update existing user
		user.GoogleID = userInfo.ID
		if user.EmailVerifiedAt == nil {
			user.EmailVerifiedAt = &now
		}
		db.Save(&user)
	}

	// Generate JWT
	jwtToken, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	c.Redirect(http.StatusTemporaryRedirect, frontendURL+"/auth/callback?token="+jwtToken)
}

// GET /api/auth/me
func GetCurrentUser(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	db := config.GetDB()
	var user models.User
	if err := db.Preload("Mahasiswa.ProgramStudi").Preload("Mahasiswa.TahunAjaran").
		Preload("Dosen.ProgramStudi").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// POST /api/auth/logout
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
