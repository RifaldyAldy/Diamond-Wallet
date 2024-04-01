package model

type Transfer struct {
	Id             string `json:"id"`
	SenderName     string `json:"nama_pengirim,omitempty"`
	UserId         string `json:"user_id"`
	Trx_id         string `json:"trx_id,omitempty"`
	Receiver       string `json:"nama_penerima,omitempty"`
	TujuanTransfer string `json:"tujuan_transfer"`
	JumlahTransfer int    `json:"jumlah_transfer"`
	JenisTransfer  string `json:"jenis_transfer"`
}
