package repository

import (
	"database/sql"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
)

type AdminRepository interface {
	Register(payload model.Admin) (model.Admin, error)
	Get(username dto.LoginRequestDto) (model.Admin, error)
}

type adminRepository struct {
	db *sql.DB
}

func (a *adminRepository) Register(payload model.Admin) (model.Admin, error) {
	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()
	err := a.db.QueryRow(`INSERT INTO mst_admin 
			(name,role,username,password,email,created_at,updated_at)
		VALUES
			($1,$2,$3,$4,$5,$6,$7)
		RETURNING id
	`, payload.Name, "admin", payload.Username, payload.Password, payload.Email, payload.CreatedAt, payload.UpdatedAt).Scan(&payload.Id)
	if err != nil {
		return model.Admin{}, err
	}
	return payload, nil
}

func (a *adminRepository) Get(username dto.LoginRequestDto) (model.Admin, error) {
	response := model.Admin{}

	err := a.db.QueryRow(`
  SELECT 
    id, name, role,username,password, email, created_at, updated_at
  FROM
    mst_admin 
  WHERE
    username = $1
    `, username.Username,
	).Scan(
		&response.Id,
		&response.Name,
		&response.Role,
		&response.Username,
		&response.Password,
		&response.Email,
		&response.CreatedAt,
		&response.UpdatedAt,
	)
	if err != nil {
		return model.Admin{}, err
	}

	return response, nil
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	return &adminRepository{db: db}
}
