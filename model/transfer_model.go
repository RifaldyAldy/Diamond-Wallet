package model

type Transfer struct {
	Id             string `json:"id"`
	UserId         string `json:"user_id"`
	TujuanTransfer string `json:"tujuan_transfer"`
	JumlahTransfer int    `json:"jumlah_transfer"`
	JenisTransfer  string `json:"jenis_transfer"`
}
