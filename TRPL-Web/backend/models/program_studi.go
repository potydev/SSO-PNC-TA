package models

import (
	"time"

	"gorm.io/gorm"
)

type ProgramStudi struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Nama      string         `gorm:"not null" json:"nama"`
	Kode      string         `gorm:"uniqueIndex" json:"kode"`
	Jenjang   string         `json:"jenjang"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Mahasiswa []Mahasiswa `gorm:"foreignKey:ProgramStudiID" json:"mahasiswa,omitempty"`
	Dosen     []Dosen     `gorm:"foreignKey:ProgramStudiID" json:"dosen,omitempty"`
}
