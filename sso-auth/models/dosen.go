package models

import (
	"time"

	"gorm.io/gorm"
)

type Dosen struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	UserID         uint           `gorm:"uniqueIndex;not null" json:"user_id"`
	NIDN           string         `gorm:"uniqueIndex" json:"nidn"`
	NIP            string         `gorm:"uniqueIndex" json:"nip"`
	TempatLahir    string         `json:"tempat_lahir"`
	TanggalLahir   *time.Time     `json:"tanggal_lahir"`
	JenisKelamin   string         `json:"jenis_kelamin"`
	ProgramStudiID *uint          `json:"program_studi_id"`
	Jabatan        string         `json:"jabatan"`
	NoTelepon      string         `json:"no_telepon"`
	Alamat         string         `gorm:"type:text" json:"alamat"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	User         User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	ProgramStudi *ProgramStudi `gorm:"foreignKey:ProgramStudiID" json:"program_studi,omitempty"`
}
