package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	Email           string         `gorm:"uniqueIndex;not null" json:"email"`
	Name            string         `gorm:"not null" json:"name"`
	Role            string         `gorm:"not null" json:"role"`
	GoogleID        string         `gorm:"uniqueIndex" json:"google_id"`
	FotoProfile     *string        `json:"foto_profile"`
	EmailVerifiedAt *time.Time     `json:"email_verified_at"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	Mahasiswa *Mahasiswa `gorm:"foreignKey:UserID" json:"mahasiswa,omitempty"`
	Dosen     *Dosen     `gorm:"foreignKey:UserID" json:"dosen,omitempty"`
}
