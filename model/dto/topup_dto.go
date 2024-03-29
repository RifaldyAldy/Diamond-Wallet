package dto

type TopupRequest struct {
	Ammount int `json:"ammount"`
}

type ResponsePayment struct {
	OrderId           string `json:"order_id"`
	UserId            string `json:"user_id"`
	Saldo             int    `json:"saldo_sekarang"`
	StatusCode        int    `json:"status_code"`
	TransactionStatus string `json:"transaction_status"`
}
