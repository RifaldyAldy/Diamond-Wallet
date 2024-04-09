package dto

type VerifyUser struct {
	UserId       string `json:"user_id"`
	Nik          string `json:"nik"`
	JenisKelamin string `json:"jenis_kelamin"`
	TanggalLahir string `json:"tanggal_lahir"`
	Umur         int    `json:"umur"`
	Photo        string `json:"photo"`
	Pin          string `json:"pin"`
}

type VerifyUserSwag struct {
	Nik          string `json:"nik"`
	JenisKelamin string `json:"jenis_kelamin"`
	TanggalLahir string `json:"tanggal_lahir"`
	Umur         int    `json:"umur"`
	Pin          string `json:"pin"`
}
