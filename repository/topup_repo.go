package repository

import "database/sql"

type TopupRepository interface {
}

type topupRepository struct {
	db *sql.DB
}

func NewTopUpRepository(db *sql.DB) TopupRepository {
	return &topupRepository{db: db}
}
