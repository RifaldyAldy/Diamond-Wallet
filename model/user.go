package model

import "time"

type User struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Password    string    `json:"password,omitempty"`
	Role        string    `json:"role"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserResponse struct {
	Name        string    `json:"name"`
	Role        string    `json:"role"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type JwtClaims struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type UserSaldo struct {
	User  UserResponse `json:"user"`
	Saldo int          `json:"saldo"`
}
