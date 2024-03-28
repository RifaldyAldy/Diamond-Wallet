package repository

import (
	"database/sql"
	"time"

	"github.com/RifaldyAldy/diamond-wallet/model"
)

type UserRepository interface {
	Create(payload model.User) (model.User, error)
}

type userRepository struct {
	db *sql.DB
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

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
