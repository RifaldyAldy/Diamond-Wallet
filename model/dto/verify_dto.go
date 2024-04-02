package dto

import "time"

type VerifyUser struct {
	UserId       string    `json:"user_id"`
	Nik          string    `json:"nik"`
	JenisKelamin string    `json:"jenis_kelamin"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Umur         int       `json:"umur"`
	Photo        string    `json:"photo"`
	Pin          string    `json:"pin"`
}
