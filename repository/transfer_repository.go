package repository

import "database/sql"

type TransferRepository interface {
}

type transferRepository struct {
	db *sql.DB
}

// tulis code kalian disini

func NewTransferRepository(db *sql.DB) TransferRepository {
	return &transferRepository{db: db}
}
