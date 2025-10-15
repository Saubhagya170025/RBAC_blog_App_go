package repository

import (
	"database/sql"
	// "time"
	"github.com/Saubhagya170025/rbac-blog-app/models"
)

func CreateCategory(db *sql.DB, categoryName string) (int, error) {
	var categoryID int
	err := db.QueryRow(
		`INSERT INTO categories (category_name) VALUES ($1) RETURNING category_id`,
		categoryName,
	).Scan(&categoryID)
	return categoryID, err
}

func GetCategoryByID(db *sql.DB, categoryID int) (*models.Category, error) {
	cat := &models.Category{}
	err := db.QueryRow(
		`SELECT category_id, category_name, created_at, updated_at FROM categories WHERE category_id = $1`,
		categoryID,
	).Scan(&cat.CategoryID, &cat.CategoryName, &cat.CreatedAt, &cat.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return cat, err
}

func DeleteCategory(db *sql.DB, categoryID int) error {
	result, err := db.Exec("DELETE FROM categories WHERE category_id = $1", categoryID)
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


func GetAllCategories(db *sql.DB) ([]models.Category, error) {
	rows, err := db.Query(
		`SELECT category_id, category_name, created_at, updated_at
		 FROM categories ORDER BY category_id ASC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.CategoryID, &c.CategoryName, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, rows.Err()
}

func UpdateCategory(db *sql.DB, categoryID int, categoryName string) error {
	result, err := db.Exec(
		`UPDATE categories SET category_name = $1, updated_at = CURRENT_TIMESTAMP
		 WHERE category_id = $2`,
		categoryName, categoryID,
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