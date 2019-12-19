package model

import "time"

type Users struct {
	Id              int       `json:"id" gorm:"column:id"`
	Name            string    `json:"name" gorm:"column:name"`
	Email           string    `json:"email" gorm:"column:email"`
	EmailVerified   time.Time `json:"email_verified_at" gorm:"column:email_verified_at"`
	Password        string    `json:"password" gorm:"column:password"`
	NomorIdentitas  string    `json:"nomor_identitas" gorm:"column:nomor_identitas"`
	NomorTelpKantor string    `json:"nomor_telp_kantor" gorm:"column:nomor_telp_kantor"`
	IdUnit          int       `json:"id_unit" gorm:"column:id_unit"`
	NamaUnit        string    `json:"nama_unit"`
	IdJabatan       int       `json:"id_jabatan" gorm:"column:id_jabatan"`
	Token           string    `json:"remember_token" gorm:"column:remember_token"`
	CreatedAt       time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"column:updated_at"`
}
