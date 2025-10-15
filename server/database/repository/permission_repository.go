package repository

import (
	"database/sql"
	"github.com/Saubhagya170025/rbac-blog-app/models"
)

func CreatePermission(db *sql.DB, perm *models.Permission) (int, error) {
	var permID int
	err := db.QueryRow(
		`INSERT INTO permissions 
		 (role_id, create_blog, create_user, create_category, create_role, 
		  update_blog, update_user, update_category, update_role, 
		  delete_blog, delete_user, delete_category, delete_role) 
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) 
		 RETURNING permission_id`,
		perm.RoleID, perm.CreateBlog, perm.CreateUser, perm.CreateCategory, perm.CreateRole,
		perm.UpdateBlog, perm.UpdateUser, perm.UpdateCategory, perm.UpdateRole,
		perm.DeleteBlog, perm.DeleteUser, perm.DeleteCategory, perm.DeleteRole,
	).Scan(&permID)
	return permID, err
}

func GetPermissionByRoleID(db *sql.DB, roleID int) (*models.Permission, error) {
	perm := &models.Permission{}
	err := db.QueryRow(
		`SELECT permission_id, role_id, create_blog, create_user, create_category, create_role,
		        update_blog, update_user, update_category, update_role,
		        delete_blog, delete_user, delete_category, delete_role, created_at, updated_at
		 FROM permissions WHERE role_id = $1`,
		roleID,
	).Scan(&perm.PermissionID, &perm.RoleID, &perm.CreateBlog, &perm.CreateUser, &perm.CreateCategory, &perm.CreateRole,
		&perm.UpdateBlog, &perm.UpdateUser, &perm.UpdateCategory, &perm.UpdateRole,
		&perm.DeleteBlog, &perm.DeleteUser, &perm.DeleteCategory, &perm.DeleteRole, &perm.CreatedAt, &perm.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return perm, err
}


func UpdatePermission(db *sql.DB, permissionID int, perm *models.Permission) error {
	result, err := db.Exec(
		`UPDATE permissions 
		 SET create_blog = $1, create_user = $2, create_category = $3, create_role = $4,
		     update_blog = $5, update_user = $6, update_category = $7, update_role = $8,
		     delete_blog = $9, delete_user = $10, delete_category = $11, delete_role = $12,
		     updated_at = CURRENT_TIMESTAMP
		 WHERE permission_id = $13`,
		perm.CreateBlog, perm.CreateUser, perm.CreateCategory, perm.CreateRole,
		perm.UpdateBlog, perm.UpdateUser, perm.UpdateCategory, perm.UpdateRole,
		perm.DeleteBlog, perm.DeleteUser, perm.DeleteCategory, perm.DeleteRole,
		permissionID,
	)
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


