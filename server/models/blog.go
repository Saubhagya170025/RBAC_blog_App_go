package models

import "time"

type Blog struct {
	BlogID      int       `json:"blog_id"`
	UserID      int       `json:"user_id"`
	User        *User     `json:"user,omitempty"`
	CategoryID  int       `json:"category_id"`
	Category    *Category `json:"category,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	FilePath    string    `json:"file_path"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}