package repository

import (
	"database/sql"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model"
)

type AdminRepository interface {
	Register(payload model.Admin) (model.Admin, error)
}

type adminRepository struct {
	db *sql.DB
}

func (a *adminRepository) Register(payload model.Admin) (model.Admin, error) {
	err := a.db.QueryRow(`INSERT INTO mst_admin 
			(name,username,password,email,created_at,updated_at)
		VALUES
			($1,$2,$3,$4,$5,$6)
		RETURNING id
	`, payload.Name, payload.Username, payload.Password, payload.Email, time.Now(), time.Now()).Scan(&payload.Id)

	if err != nil {
		return model.Admin{}, err
	}
	return payload, nil
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	return &adminRepository{db: db}
}
