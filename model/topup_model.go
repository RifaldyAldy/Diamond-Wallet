package model

import "github.com/midtrans/midtrans-go"

type TopupModel struct {
	User               User
	TransactionDetails midtrans.TransactionDetails `json:"transaction_details"`
}
