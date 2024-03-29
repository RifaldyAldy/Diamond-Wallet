package model

import (
	"time"

	"github.com/midtrans/midtrans-go"
)

type TopupModel struct {
	User               User
	TransactionDetails midtrans.TransactionDetails `json:"transaction_details"`
}

type TableTopupPayment struct {
	OrderId       string    `json:"id"`
	UserId        string    `json:"user_id"`
	TokenMidtrans string    `json:"token_midtrans"`
	UrlPayment    string    `json:"url_payment"`
	Status        string    `json:"status"`
	Deskripsi     string    `json:"deskripsi"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}
