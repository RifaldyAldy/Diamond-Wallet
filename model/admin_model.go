package model

import "time"

type Admin struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
