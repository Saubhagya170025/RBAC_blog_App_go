package models

import "time"

type User struct {
	UserID       int       `json:"user_id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password,omitempty"`
	RefreshToken string    `json:"-"`
	RoleID       int       `json:"role_id"`
	Role         *Role     `json:"role,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}