package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/RifaldyAldy/diamond-wallet/utils/common"
)

type TopupRepository interface {
	Create(payload model.TopupModel) (common.ResponseMidtrans, error)
	Getbyid(orderId string) (model.TableTopupPayment, error)
	Payment(payload dto.ResponsePayment) (dto.ResponsePayment, error)
}

type topupRepository struct {
	db *sql.DB
}

func (t *topupRepository) Create(payload model.TopupModel) (common.ResponseMidtrans, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return common.ResponseMidtrans{}, err
	}
	deskripsi := fmt.Sprintf("%s ingin melakukan top up saldo sebesar %d", payload.User.Name, payload.TransactionDetails.GrossAmt)
	err = tx.QueryRow(`INSERT INTO trx_topup_method_payment(user_id,ammount,status,deskripsi,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`, payload.User.Id, payload.TransactionDetails.GrossAmt, "Menunggu pembayaran", deskripsi, time.Now(), time.Now()).Scan(&payload.TransactionDetails.OrderID)
	if err != nil {
		tx.Rollback()
		return common.ResponseMidtrans{}, err
	}

	midtransResponse, err := common.GenerateMidtrans(payload.TransactionDetails)
	if err != nil {
		tx.Rollback()
		return common.ResponseMidtrans{}, err
	}
	_, err = tx.Query(`UPDATE trx_topup_method_payment SET 
		token_midtrans=$1, 
		url_payment=$2 WHERE id=$3`, midtransResponse.Token, midtransResponse.UrlPayment, payload.TransactionDetails.OrderID)
	if err != nil {
		tx.Rollback()
		return common.ResponseMidtrans{}, err
	}
	tx.Commit()

	return midtransResponse, nil
}

func (t *topupRepository) Getbyid(orderId string) (model.TableTopupPayment, error) {
	var tabel model.TableTopupPayment

	err := t.db.QueryRow(`SELECT 
	id,user_id,token_midtrans,url_payment,status,deskripsi,created_at,updated_at
	FROM 
	trx_topup_method_payment WHERE id =$1`, orderId).Scan(
		&tabel.OrderId,
		&tabel.UserId,
		&tabel.TokenMidtrans,
		&tabel.UrlPayment,
		&tabel.Status,
		&tabel.Deskripsi,
		&tabel.Created_at,
		&tabel.Updated_at,
	)
	if err != nil {
		return model.TableTopupPayment{}, err
	}

	return tabel, nil
}

func (t *topupRepository) Payment(payload dto.ResponsePayment) (dto.ResponsePayment, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return dto.ResponsePayment{}, err
	}
	var saldoTopup int
	tabel, err := t.Getbyid(payload.OrderId)
	if err != nil {
		return dto.ResponsePayment{}, err
	}
	if tabel.Status == "Pembayaran berhasil" {
		tx.Rollback()
		return dto.ResponsePayment{}, errors.New("link ini sudah tidak valid")
	}
	payload.UserId = tabel.UserId
	err = tx.QueryRow(`SELECT saldo FROM mst_saldo WHERE user_id = $1`, payload.UserId).Scan(&payload.Saldo)
	if err != nil {
		tx.Rollback()
		return dto.ResponsePayment{}, err
	}
	err = tx.QueryRow("UPDATE trx_topup_method_payment SET status=$1, updated_at=$2 WHERE id =$3 RETURNING user_id,ammount", "Pembayaran berhasil", time.Now(), payload.OrderId).Scan(&payload.UserId, &saldoTopup)
	if err != nil {
		tx.Rollback()
		return dto.ResponsePayment{}, err
	}

	payload.Saldo += saldoTopup
	_, err = tx.Query(`UPDATE mst_saldo SET saldo=$1 WHERE user_id=$2`, payload.Saldo, payload.UserId)
	if err != nil {
		tx.Rollback()
		return dto.ResponsePayment{}, err
	}

	tx.Commit()

	return payload, nil
}

func NewTopUpRepository(db *sql.DB) TopupRepository {
	return &topupRepository{db: db}
}
