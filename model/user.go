package model

import "time"

type User struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username,omitempty"`
	Password    string    `json:"password,omitempty"`
	Role        string    `json:"role"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Saldo       int       `json:"saldo,omitempty"`
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
	Pin   string       `json:"pin"`
}

type Rekening struct {
	Id         string    `json:"id"`
	UserId     string    `json:"user_id"`
	Rekening   string    `json:"rekening"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
