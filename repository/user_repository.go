package repository

import (
	"database/sql"

	"github.com/RifaldyAldy/diamond-wallet/model"
)

type UserRepository interface {
	GetByUsername(username string) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

// TODO: buat query user login
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
		&user.UpdatedAt)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
