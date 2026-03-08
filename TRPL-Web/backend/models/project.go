package models

import (
	"time"
)

type Project struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	MahasiswaID     uint      `gorm:"not null" json:"mahasiswa_id"`
	Judul           string    `gorm:"not null" json:"judul"`
	Deskripsi       string    `gorm:"type:text" json:"deskripsi"`
	Kategori        string    `json:"kategori"` // Web, Mobile, Desktop, AI/ML, dll
	TahunPembuatan  int       `json:"tahun_pembuatan"`
	Teknologi       string    `gorm:"type:text" json:"teknologi"` // JSON array of tech stack
	URLDemo         string    `json:"url_demo"`
	URLRepository   string    `json:"url_repository"`
	URLGambar       string    `json:"url_gambar"` // Cover image
	Status          string    `gorm:"default:'Draft'" json:"status"`
	ViewCount       int       `gorm:"default:0" json:"view_count"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	
	// Relations
	Mahasiswa       Mahasiswa `gorm:"foreignKey:MahasiswaID" json:"mahasiswa,omitempty"`
}
