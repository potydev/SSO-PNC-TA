package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	Email            string         `gorm:"uniqueIndex;not null" json:"email"`
	Name             string         `gorm:"not null" json:"name"`
	Role             string         `gorm:"type:enum('Mahasiswa','Dosen');not null" json:"role"`
	GoogleID         string         `gorm:"uniqueIndex" json:"google_id"`
	FotoProfile      *string        `json:"foto_profile"`
	EmailVerifiedAt  *time.Time     `json:"email_verified_at"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	
	// Relations
	Mahasiswa        *Mahasiswa     `gorm:"foreignKey:UserID" json:"mahasiswa,omitempty"`
	Dosen            *Dosen         `gorm:"foreignKey:UserID" json:"dosen,omitempty"`
}

type Mahasiswa struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	UserID          uint       `gorm:"uniqueIndex;not null" json:"user_id"`
	NIM             string     `gorm:"uniqueIndex;not null" json:"nim"`
	TempatLahir     string     `json:"tempat_lahir"`
	TanggalLahir    *time.Time `json:"tanggal_lahir"`
	JenisKelamin    string     `gorm:"type:enum('L','P')" json:"jenis_kelamin"`
	ProgramStudiID  *uint      `json:"program_studi_id"`
	TahunAjaranID   *uint      `json:"tahun_ajaran_id"`
	NoTelepon       string     `json:"no_telepon"`
	Alamat          string     `gorm:"type:text" json:"alamat"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	
	// Relations
	User            User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	ProgramStudi    *ProgramStudi  `gorm:"foreignKey:ProgramStudiID" json:"program_studi,omitempty"`
	TahunAjaran     *TahunAjaran   `gorm:"foreignKey:TahunAjaranID" json:"tahun_ajaran,omitempty"`
	Projects        []Project      `gorm:"foreignKey:MahasiswaID" json:"projects,omitempty"`
}

type Dosen struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	UserID          uint       `gorm:"uniqueIndex;not null" json:"user_id"`
	NIDN            string     `gorm:"uniqueIndex" json:"nidn"`
	NIP             string     `gorm:"uniqueIndex" json:"nip"`
	TempatLahir     string     `json:"tempat_lahir"`
	TanggalLahir    *time.Time `json:"tanggal_lahir"`
	JenisKelamin    string     `gorm:"type:enum('L','P')" json:"jenis_kelamin"`
	ProgramStudiID  *uint      `json:"program_studi_id"`
	Jabatan         string     `json:"jabatan"`
	NoTelepon       string     `json:"no_telepon"`
	Alamat          string     `gorm:"type:text" json:"alamat"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	
	// Relations
	User            User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	ProgramStudi    *ProgramStudi `gorm:"foreignKey:ProgramStudiID" json:"program_studi,omitempty"`
}

type ProgramStudi struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"not null" json:"nama"`
	Kode      string    `gorm:"uniqueIndex" json:"kode"`
	Jenjang   string    `json:"jenjang"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TahunAjaran struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"not null" json:"nama"`
	Aktif     bool      `gorm:"default:false" json:"aktif"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
