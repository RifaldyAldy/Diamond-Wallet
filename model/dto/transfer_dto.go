package dto

type TransferRequest struct {
	UserId         string `json:"user_id"`
	TujuanTransfer string `json:"tujuan_transfer"`
	Pin            string `json:"pin"`
	JumlahTransfer int    `json:"jumlah_transfer"`
}

type TransferRequestSwag struct {
	TujuanTransfer string `json:"tujuan_transfer"`
	Pin            string `json:"pin"`
	JumlahTransfer int    `json:"jumlah_transfer"`
}
