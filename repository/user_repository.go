package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model"
	"github.com/RifaldyAldy/diamond-wallet/model/dto"
	"github.com/lib/pq"
)

type UserRepository interface {
	Get(id string) (model.User, error)
	Create(payload dto.UserRequestDto) (model.User, error)
	GetBalance(user_id string) (model.UserSaldo, error)
	GetByUsername(username string) (model.User, error)
	Update(id string, payload model.User) (model.User, error)
	Verify(payload dto.VerifyUser) (dto.VerifyUser, error)
	UpdatePin(payload dto.UpdatePinRequest) (dto.UpdatePinResponse, error)
	GetInfoUser(Info string, limit, offset int) ([]model.User, error)
	GetRekening(id string) (model.Rekening, error)
	CreateRekening(payload model.Rekening) (model.Rekening, error)
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

func (u *userRepository) Create(payload dto.UserRequestDto) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`
  INSERT INTO mst_user
    (name, username, password, email, phone_number, created_at, updated_at)
  VALUES
    ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id, name, username, password, role, email, phone_number, created_at, updated_at
    `,
		payload.Name,
		payload.Username,
		payload.Password,
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
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code == "23505" {
			if pgErr.Constraint == "mst_user_email_key" {
				return model.User{}, fmt.Errorf("email sudah terdaftar")
			} else if pgErr.Constraint == "mst_user_phone_number_key" {
				return model.User{}, fmt.Errorf("phonenumber sudah terdaftar,pakai nomor lain")
			} else if pgErr.Constraint == "mst_user_username_key" {
				return model.User{}, fmt.Errorf("username sudah terdaftar")
			}
		}
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
		s.saldo,
		s.pin
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
		&response.Pin,
	)
	if response.Pin == "" {
		return model.UserSaldo{}, fmt.Errorf("1")
	}
	if err != nil {
		return model.UserSaldo{}, err
	}

	return response, nil
}

func (u *userRepository) Update(id string, payload model.User) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`
  UPDATE mst_user SET
    name = $1, email=$2, phone_number=$3, updated_at=$4 
	WHERE id=$5
    RETURNING id, name, role, email, phone_number, created_at, updated_at
		
    `,
		payload.Name,
		payload.Email,
		payload.PhoneNumber,
		time.Now(),
		id,
	).Scan(
		&user.Id,
		&user.Name,
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

func (u *userRepository) Verify(payload dto.VerifyUser) (dto.VerifyUser, error) {
	tx, err := u.db.Begin()
	if err != nil {
		return dto.VerifyUser{}, err
	}
	if payload.Pin == "" {
		return dto.VerifyUser{}, fmt.Errorf("pin harus diisi")
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
		pgErr, ok := err.(*pq.Error)
		if ok && pgErr.Code == "23505" {
			if pgErr.Constraint == "mst_user_datas_nik_key" {
				return dto.VerifyUser{}, fmt.Errorf("nik sudah terdaftar: %s", pgErr.Detail)
			} else if pgErr.Constraint == "mst_user_datas_user_id_key" {
				return dto.VerifyUser{}, fmt.Errorf("user ini sudah verifikasi")
			}
		}
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

func (u *userRepository) GetInfoUser(info string, limit, offset int) ([]model.User, error) {
	query := `
		SELECT u.id, u.name, u.role, u.email, u.phone_number, u.created_at, u.updated_at,
			   s.saldo
		FROM mst_user u
		LEFT JOIN mst_saldo s ON u.id = s.user_id
		WHERE ` + info + ` LIMIT $1 OFFSET $2`

	rows, err := u.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Role,
			&user.Email,
			&user.PhoneNumber,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Saldo,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) UpdatePin(payload dto.UpdatePinRequest) (dto.UpdatePinResponse, error) {
	var response dto.UpdatePinResponse
	err := u.db.QueryRow(`
   UPDATE 
    mst_saldo 
  SET 
    pin = $1 
  WHERE 
    user_id = $2
  RETURNING user_id, pin
    `, payload.NewPin,
		payload.UserId,
	).Scan(
		&response.UserId,
		&response.Pin,
	)
	if err != nil {
		return dto.UpdatePinResponse{}, nil
	}
	return response, nil
}

func (u *userRepository) GetRekening(id string) (model.Rekening, error) {
	response := model.Rekening{}
	err := u.db.QueryRow(`SELECT * FROM mst_rekening_user WHERE user_id = $1`, id).Scan(&response.Id,
		&response.UserId,
		&response.Rekening,
		&response.Created_at,
		&response.Updated_at,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Rekening{}, fmt.Errorf("1")
		}
		return model.Rekening{}, err
	}

	return response, nil
}

func (u *userRepository) CreateRekening(payload model.Rekening) (model.Rekening, error) {
	err := u.db.QueryRow(`INSERT INTO mst_rekening_user (user_id,rekening,created_at,updated_at)
	VALUES
		($1,$2,$3,$4)
	RETURNING id,created_at,updated_at
	`, payload.UserId, payload.Rekening, time.Now(), time.Now()).Scan(&payload.Id, &payload.Created_at, &payload.Updated_at)
	if err != nil {
		return model.Rekening{}, err
	}

	return payload, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
