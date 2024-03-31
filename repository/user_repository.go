package repository

import (
	"database/sql"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
)

type UserRepository interface {
	Get(id string) (model.User, error)
	Create(payload model.User) (model.User, error)
	GetBalance(user_id string) (model.UserSaldo, error)
	GetByUsername(username string) (model.User, error)
	Verify(payload dto.VerifyUser) (dto.VerifyUser, error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Get(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`
	SELECT 
	  id, name, username, role, email, phone_number, created_at, updated_at
	FROM
	  mst_user 
	WHERE
	  id = $1
	  `, id).Scan(&user.Id,
		&user.Name,
		&user.Username,
		&user.Role,
		&user.Email,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) Create(payload model.User) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`
  INSERT INTO mst_user
    (name, username, password, role, email, phone_number, created_at, updated_at)
  VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id, name, username, password, role, email, phone_number, created_at, updated_at
    `,
		payload.Name,
		payload.Username,
		payload.Password,
		payload.Role,
		payload.Email,
		payload.PhoneNumber,
		time.Now(),
		time.Now(),
	).Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.Email,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userRepository) GetByUsername(username string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`
  SELECT 
    id, name, username, password, role, email, phone_number, created_at, updated_at
  FROM
    mst_user 
  WHERE
    username = $1
    `, username,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.Email,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u *userRepository) GetBalance(user_id string) (model.UserSaldo, error) {
	var response model.UserSaldo

	err := u.db.QueryRow(`SELECT 
		u.name,
		u.role,
		u.email,
		u.phone_number,
		u.created_at,
		u.updated_at,
		s.saldo
	FROM 
   		 mst_user AS u
	LEFT JOIN 
    	mst_saldo AS s ON u.id = s.user_id
	WHERE 
		u.id = $1;`, user_id).Scan(
		&response.User.Name,
		&response.User.Role,
		&response.User.Email,
		&response.User.PhoneNumber,
		&response.User.CreatedAt,
		&response.User.UpdatedAt,
		&response.Saldo,
	)
	if err != nil {
		return model.UserSaldo{}, err
	}

	return response, nil
}

func (u *userRepository) Verify(payload dto.VerifyUser) (dto.VerifyUser, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return dto.VerifyUser{}, err
	}

	userId, err := u.Get(payload.UserId)
	if err != nil {
		tx.Rollback()
		return dto.VerifyUser{}, err
	}

	_, err = tx.Exec(`
        INSERT INTO mst_user_datas
            (user_id, nik, jenis_kelamin, tanggal_lahir, umur, photo)
        VALUES
            ($1, $2, $3, $4, $5, $6)
    `, userId.Id, payload.Nik, payload.JenisKelamin, payload.TanggalLahir, payload.Umur, payload.Photo)
	if err != nil {
		tx.Rollback()
		return dto.VerifyUser{}, err
	}

	_, err = tx.Exec(`
        INSERT INTO mst_saldo
            (user_id, saldo, pin)
        VALUES
            ($1, $2, $3)
    `, userId.Id, 0, payload.Pin)
	if err != nil {
		tx.Rollback()
		return dto.VerifyUser{}, err
	}

	if err := tx.Commit(); err != nil {
		return dto.VerifyUser{}, err
	}

	return payload, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
