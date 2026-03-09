package models

import (
	"time"

	"gorm.io/gorm"
)

type Mahasiswa struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	UserID         uint           `gorm:"uniqueIndex;not null" json:"user_id"`
	NIM            string         `gorm:"uniqueIndex;not null" json:"nim"`
	TempatLahir    string         `json:"tempat_lahir"`
	TanggalLahir   *time.Time     `json:"tanggal_lahir"`
	JenisKelamin   string         `json:"jenis_kelamin"`
	ProgramStudiID *uint          `json:"program_studi_id"`
	TahunAjaranID  *uint          `json:"tahun_ajaran_id"`
	NoTelepon      string         `json:"no_telepon"`
	Alamat         string         `gorm:"type:text" json:"alamat"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	User         User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	ProgramStudi *ProgramStudi `gorm:"foreignKey:ProgramStudiID" json:"program_studi,omitempty"`
	TahunAjaran  *TahunAjaran  `gorm:"foreignKey:TahunAjaranID" json:"tahun_ajaran,omitempty"`
}
