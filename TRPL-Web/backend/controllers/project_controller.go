package controllers

import (
	"net/http"
	"trpl-web-backend/config"
	"trpl-web-backend/models"

	"github.com/gin-gonic/gin"
)

// GET /api/projects
func GetProjects(c *gin.Context) {
	db := config.GetDB()
	
	var projects []models.Project
	query := db.Preload("Mahasiswa.User").Preload("Mahasiswa.ProgramStudi").
		Where("status = ?", "Published").Order("created_at desc")

	// Filters
	if kategori := c.Query("kategori"); kategori != "" {
		query = query.Where("kategori = ?", kategori)
	}
	if tahun := c.Query("tahun"); tahun != "" {
		query = query.Where("tahun_pembuatan = ?", tahun)
	}

	if err := query.Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

// GET /api/projects/:id
func GetProjectByID(c *gin.Context) {
	id := c.Param("id")
	
	db := config.GetDB()
	var project models.Project
	if err := db.Preload("Mahasiswa.User").Preload("Mahasiswa.ProgramStudi").
		First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Increment view count
	db.Model(&project).Update("view_count", project.ViewCount+1)

	c.JSON(http.StatusOK, gin.H{"project": project})
}

// POST /api/projects (Mahasiswa only)
func CreateProject(c *gin.Context) {
	role, _ := c.Get("role")
	if role != "Mahasiswa" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only mahasiswa can create projects"})
		return
	}

	userID, _ := c.Get("user_id")
	
	// Get mahasiswa ID from user
	db := config.GetDB()
	var mahasiswa models.Mahasiswa
	if err := db.Where("user_id = ?", userID).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa profile not found. Please complete your profile first."})
		return
	}

	var input struct {
		Judul          string `json:"judul" binding:"required"`
		Deskripsi      string `json:"deskripsi"`
		Kategori       string `json:"kategori"`
		TahunPembuatan int    `json:"tahun_pembuatan"`
		Teknologi      string `json:"teknologi"`
		URLDemo        string `json:"url_demo"`
		URLRepository  string `json:"url_repository"`
		URLGambar      string `json:"url_gambar"`
		Status         string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project := models.Project{
		MahasiswaID:    mahasiswa.ID,
		Judul:          input.Judul,
		Deskripsi:      input.Deskripsi,
		Kategori:       input.Kategori,
		TahunPembuatan: input.TahunPembuatan,
		Teknologi:      input.Teknologi,
		URLDemo:        input.URLDemo,
		URLRepository:  input.URLRepository,
		URLGambar:      input.URLGambar,
		Status:         input.Status,
	}

	if project.Status == "" {
		project.Status = "Draft"
	}

	if err := db.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"project": project})
}

// PUT /api/projects/:id (Owner only)
func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	db := config.GetDB()
	var project models.Project
	if err := db.Preload("Mahasiswa").First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Check ownership
	if project.Mahasiswa.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only edit your own projects"})
		return
	}

	var input models.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	updates := map[string]interface{}{
		"judul":           input.Judul,
		"deskripsi":       input.Deskripsi,
		"kategori":        input.Kategori,
		"tahun_pembuatan": input.TahunPembuatan,
		"teknologi":       input.Teknologi,
		"url_demo":        input.URLDemo,
		"url_repository":  input.URLRepository,
		"url_gambar":      input.URLGambar,
		"status":          input.Status,
	}

	if err := db.Model(&project).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"project": project})
}

// DELETE /api/projects/:id (Owner only)
func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	db := config.GetDB()
	var project models.Project
	if err := db.Preload("Mahasiswa").First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Check ownership
	if project.Mahasiswa.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own projects"})
		return
	}

	if err := db.Delete(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

// GET /api/my-projects (Mahasiswa only)
func GetMyProjects(c *gin.Context) {
	userID, _ := c.Get("user_id")
	
	db := config.GetDB()
	var mahasiswa models.Mahasiswa
	if err := db.Where("user_id = ?", userID).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa profile not found"})
		return
	}

	var projects []models.Project
	if err := db.Where("mahasiswa_id = ?", mahasiswa.ID).Order("created_at desc").Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

// GET /api/stats
func GetStats(c *gin.Context) {
	db := config.GetDB()
	
	var totalProjects int64
	var totalMahasiswa int64
	var categories []struct {
		Kategori string
		Count    int64
	}

	db.Model(&models.Project{}).Where("status = ?", "Published").Count(&totalProjects)
	db.Model(&models.Mahasiswa{}).Count(&totalMahasiswa)
	db.Model(&models.Project{}).Select("kategori, COUNT(*) as count").
		Where("status = ?", "Published").Group("kategori").Scan(&categories)

	c.JSON(http.StatusOK, gin.H{
		"total_projects":   totalProjects,
		"total_mahasiswa":  totalMahasiswa,
		"categories":       categories,
	})
}
