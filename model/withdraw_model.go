package model

import "time"

type Withdraw struct {
	Id         string    `json:"id"`
	UserId     string    `json:"user_id"`
	Withdraw   int       `json:"withdraw"`
	Created_at time.Time `json:"created_at"`
}
