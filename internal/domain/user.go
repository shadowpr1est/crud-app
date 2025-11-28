package domain

import "time"

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
