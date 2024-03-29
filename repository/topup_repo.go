package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/utils/common"
)

type TopupRepository interface {
	Create(payload model.TopupModel) (any, error)
}

type topupRepository struct {
	db *sql.DB
}

func (t *topupRepository) Create(payload model.TopupModel) (any, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return "", err
	}
	deskripsi := fmt.Sprintf("%s ingin melakukan top up saldo sebesar %d", payload.User.Name, payload.TransactionDetails.GrossAmt)
	err = tx.QueryRow(`INSERT INTO trx_topup_method_payment(user_id,status,deskripsi,created_at,updated_at) VALUES ($1,$2,$3,$4,$5) RETURNING id`, payload.User.Id, "Menunggu pembayaran", deskripsi, time.Now(), time.Now()).Scan(&payload.TransactionDetails.OrderID)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	midtransResponse, err := common.GenerateMidtrans(payload.TransactionDetails)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	_, err = tx.Query(`UPDATE trx_topup_method_payment SET 
		token_midtrans=$1, 
		url_payment=$2 WHERE id=$3`, midtransResponse.Token, midtransResponse.UrlPayment, payload.TransactionDetails.OrderID)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()

	return midtransResponse, nil
}

func NewTopUpRepository(db *sql.DB) TopupRepository {
	return &topupRepository{db: db}
}
