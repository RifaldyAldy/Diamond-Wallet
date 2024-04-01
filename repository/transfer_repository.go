package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
)

type TransferRepository interface {
	Create(payload dto.TransferRequest, send, receive model.User) (model.Transfer, error)
}

type transferRepository struct {
	db *sql.DB
}

// tulis code kalian disini
func (t *transferRepository) Create(payload dto.TransferRequest, send, receive model.User) (model.Transfer, error) {
	response := model.Transfer{}
	tx, err := t.db.Begin()
	if err != nil {
		tx.Rollback()
		return model.Transfer{}, err
	}
	send.Saldo -= payload.JumlahTransfer
	receive.Saldo += payload.JumlahTransfer
	if send.Saldo < 0 {
		tx.Rollback()
		return model.Transfer{}, fmt.Errorf("saldo anda tidak mendcukupi untuk transfer %d", payload.JumlahTransfer)
	}

	// buat catatan penerima ke database
	err = tx.QueryRow(`INSERT INTO trx_send_transfer (
		user_id,
		tujuan_transfer,
		jumlah_transfer,
		jenis_transfer
		)
	VALUES (
		$1,
		$2,
		$3,
		$4
	) RETURNING id`, send.Id, receive.Id, payload.JumlahTransfer, "mengirim").Scan(&response.Id)
	if err != nil {
		tx.Rollback()
		return model.Transfer{}, err
	}

	_, err = tx.Exec(`UPDATE mst_saldo SET saldo=$1 WHERE user_id=$2`, receive.Saldo, receive.Id)
	if err != nil {
		tx.Rollback()
		return model.Transfer{}, err
	}
	_, err = tx.Exec(`UPDATE mst_saldo SET saldo = $1 WHERE user_id=$2`, send.Saldo, send.Id)
	if err != nil {
		tx.Rollback()
		return model.Transfer{}, err
	}
	response.JenisTransfer = "mengirim"
	_, err = tx.Query(`INSERT INTO trx_receive_transfer (
		user_id,
		trx_id,
		tujuan_transfer,
		jumlah_transfer,
		jenis_transfer,
		transfer_at)
	VALUES ($1,$2,$3,$4,$5,$6)`, send.Id, response.Id, receive.Id, payload.JumlahTransfer, "menerima", time.Now())
	if err != nil {
		tx.Rollback()
		return model.Transfer{}, err
	}

	response.UserId = send.Id
	response.TujuanTransfer = receive.Id
	response.JumlahTransfer = payload.JumlahTransfer
	tx.Commit()

	return response, nil
}

func NewTransferRepository(db *sql.DB) TransferRepository {
	return &transferRepository{db: db}
}
