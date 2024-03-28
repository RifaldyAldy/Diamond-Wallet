package repository

import "database/sql"

type PingRepository interface {
	Ping() error
}

type pingRepository struct {
	db *sql.DB
}

func (p *pingRepository) Ping() error {
	err := p.db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func NewPingRepository(db *sql.DB) PingRepository {
	return &pingRepository{
		db: db,
	}
}
