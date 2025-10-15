package models

import "time"

type Permission struct {
	PermissionID   int       `json:"permission_id"`
	RoleID         int       `json:"role_id"`
	CreateBlog     bool      `json:"create_blog"`
	CreateUser     bool      `json:"create_user"`
	CreateCategory bool      `json:"create_category"`
	CreateRole     bool      `json:"create_role"`
	UpdateBlog     bool      `json:"update_blog"`
	UpdateUser     bool      `json:"update_user"`
	UpdateCategory bool      `json:"update_category"`
	UpdateRole     bool      `json:"update_role"`
	DeleteBlog     bool      `json:"delete_blog"`
	DeleteUser     bool      `json:"delete_user"`
	DeleteCategory bool      `json:"delete_category"`
	DeleteRole     bool      `json:"delete_role"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}