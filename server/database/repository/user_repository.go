package repository

import (
	"database/sql"
	"errors"

	"github.com/Saubhagya170025/rbac-blog-app/models"
)

// CreateUser creates a new user
func CreateUser(db *sql.DB, name, email, password string, roleID int) (int, error) {
	var userID int
	err := db.QueryRow(
		`INSERT INTO users (name, email, password, role_id) 
		 VALUES ($1, $2, $3, $4) RETURNING user_id`,
		name, email, password, roleID,
	).Scan(&userID)
	return userID, err
}

// GetAllUsers retrieves all users
func GetAllUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query(
		`SELECT user_id, name, email, password, refresh_token, role_id, created_at, updated_at
		 FROM users ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		var rt sql.NullString
		if err := rows.Scan(&u.UserID, &u.Name, &u.Email, &u.Password, &rt, &u.RoleID, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		if rt.Valid {
			u.RefreshToken = rt.String
		} else {
			u.RefreshToken = ""
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

// GetUserByID retrieves a user by ID
func 	GetUserByID(db *sql.DB, userID int) (*models.User, error) {
	user := &models.User{}
	var rt sql.NullString
	err := db.QueryRow(
		`SELECT user_id, name, email, password, refresh_token, role_id, created_at, updated_at
		 FROM users WHERE user_id = $1`,
		userID,
	).Scan(&user.UserID, &user.Name, &user.Email, &user.Password, &rt, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if rt.Valid {
		user.RefreshToken = rt.String
	} else {
		user.RefreshToken = ""
	}
	return user, nil
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	user := &models.User{}
	var rt sql.NullString
	err := db.QueryRow(
		`SELECT user_id, name, email, password, refresh_token, role_id, created_at, updated_at
		 FROM users WHERE email = $1`,
		email,
	).Scan(&user.UserID, &user.Name, &user.Email, &user.Password, &rt, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if rt.Valid {
		user.RefreshToken = rt.String
	} else {
		user.RefreshToken = ""
	}
	return user, nil
}

// UpdateUser updates user information
func UpdateUser(db *sql.DB, userID int, name, email string, roleID int) error {
	result, err := db.Exec(
		`UPDATE users SET name = $1, email = $2, role_id = $3, updated_at = CURRENT_TIMESTAMP
		 WHERE user_id = $4`,
		name, email, roleID, userID,
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

// DeleteUser deletes a user
func DeleteUser(db *sql.DB, userID int) error {
	result, err := db.Exec("DELETE FROM users WHERE user_id = $1", userID)
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

// ============================================
// JWT TOKEN MANAGEMENT FUNCTIONS
// ============================================

// StoreRefreshToken stores the refresh token in the database
func StoreRefreshToken(db *sql.DB, userID int, refreshToken string) error {
	_, err := db.Exec(
		`UPDATE users SET refresh_token = $1, updated_at = CURRENT_TIMESTAMP WHERE user_id = $2`,
		refreshToken, userID,
	)
	return err
}

// GetRefreshToken retrieves the stored refresh token for a user
func GetRefreshToken(db *sql.DB, userID int) (string, error) {
	var refreshToken sql.NullString
	err := db.QueryRow(
		`SELECT refresh_token FROM users WHERE user_id = $1`,
		userID,
	).Scan(&refreshToken)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user not found")
		}
		return "", err
	}

	if !refreshToken.Valid {
		return "", errors.New("no refresh token found")
	}

	return refreshToken.String, nil
}

// DeleteRefreshToken removes the refresh token (logout)
func DeleteRefreshToken(db *sql.DB, userID int) error {
	result, err := db.Exec(
		`UPDATE users SET refresh_token = NULL, updated_at = CURRENT_TIMESTAMP WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}