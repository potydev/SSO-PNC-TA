package models

import (
	"time"

	"gorm.io/gorm"
)

type TahunAjaran struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Nama      string         `gorm:"not null" json:"nama"`
	Aktif     bool           `gorm:"default:false" json:"aktif"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
