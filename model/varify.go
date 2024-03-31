package model

import "time"

type Verify struct {
	UserDatas UserDatas
	Saldo     Saldo
}

type UserDatas struct {
	UserId       string    `json:"user_id"`
	Nik          string    `json:"nik"`
	JenisKelamin string    `json:"jenis_kelamin"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Umur         int       `json:"umur"`
	Photo        string    `json:"photo"`
}

type Saldo struct {
	UserId string  `json:"user_id"`
	Saldo  float64 `json:"saldo"`
	Pin    string  `json:"pin"`
}
