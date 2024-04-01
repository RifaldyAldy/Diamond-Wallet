package repository

import (
	"database/sql"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model"
)

type UserRepository interface {
	Get(id string) (model.User, error)
	Create(payload model.User) (model.User, error)
	GetBalance(user_id string) (model.UserSaldo, error)
	GetByUsername(username string) (model.User, error)
	Update(id string, payload model.User) (model.User, error)
	GetInfoUser(Info string, limit, offset int) ([]model.User, error)
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

func (u *userRepository) Update(id string, payload model.User) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`
  UPDATE mst_user SET
    name = $1, username = $2, password=$3, role=$4, email=$5, phone_number=$6, created_at=$7, updated_at=$8 
	WHERE id=$9
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
		id,
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

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
