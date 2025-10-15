package repository

import (
	"database/sql"
	// "time"
	"github.com/Saubhagya170025/rbac-blog-app/models"
)

func CreateRole(db *sql.DB, roleName string) (int, error) {
	var roleID int
	err := db.QueryRow(
		`INSERT INTO roles (role_name) VALUES ($1) RETURNING role_id`,
		roleName,
	).Scan(&roleID)
	return roleID, err
}

func GetAllRoles(db *sql.DB) ([]models.Role, error) {
	rows, err := db.Query(
		`SELECT role_id, role_name, created_at, updated_at
		 FROM roles ORDER BY role_id ASC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var r models.Role
		if err := rows.Scan(&r.RoleID, &r.RoleName, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, err
		}
		roles = append(roles, r)
	}
	return roles, rows.Err()
}

func GetRoleByID(db *sql.DB, roleID int) (*models.Role, error) {
	role := &models.Role{}
	err := db.QueryRow(
		`SELECT role_id, role_name, created_at, updated_at FROM roles WHERE role_id = $1`,
		roleID,
	).Scan(&role.RoleID, &role.RoleName, &role.CreatedAt, &role.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return role, err
}

func DeleteRole(db *sql.DB, roleID int) error {
	result, err := db.Exec("DELETE FROM roles WHERE role_id = $1", roleID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}