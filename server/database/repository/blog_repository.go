package repository

import (
	"database/sql"

	"github.com/Saubhagya170025/rbac-blog-app/models"
)

func CreateBlog(db *sql.DB, userID, categoryID int, title, description, content, filePath string) (int, error) {
	var blogID int
	err := db.QueryRow(
		`INSERT INTO blogs (user_id, category_id, title, description, content, file_path)
		 VALUES ($1, $2, $3, $4, $5, $6) RETURNING blog_id`,
		userID, categoryID, title, description, content, filePath,
	).Scan(&blogID)
	return blogID, err
}

func GetBlogByID(db *sql.DB, blogID int) (*models.Blog, error) {
	blog := &models.Blog{}
	err := db.QueryRow(
		`SELECT blog_id, user_id, category_id, title, description, content, file_path, created_at, updated_at
		 FROM blogs WHERE blog_id = $1`,
		blogID,
	).Scan(&blog.BlogID, &blog.UserID, &blog.CategoryID, &blog.Title, &blog.Description, &blog.Content, &blog.FilePath, &blog.CreatedAt, &blog.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return blog, err
}



func GetBlogsByUserID(db *sql.DB, userID int) ([]models.Blog, error) {
	rows, err := db.Query(
		`SELECT blog_id, user_id, category_id, title, description, content, file_path, created_at, updated_at
		 FROM blogs WHERE user_id = $1 ORDER BY created_at DESC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []models.Blog
	for rows.Next() {
		var b models.Blog
		if err := rows.Scan(&b.BlogID, &b.UserID, &b.CategoryID, &b.Title, &b.Description, &b.Content, &b.FilePath, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}
		blogs = append(blogs, b)
	}
	return blogs, rows.Err()
}

func GetBlogsByCategoryID(db *sql.DB, categoryID int) ([]models.Blog, error) {
	rows, err := db.Query(
		`SELECT blog_id, user_id, category_id, title, description, content, file_path, created_at, updated_at
		 FROM blogs WHERE category_id = $1 ORDER BY created_at DESC`,
		categoryID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []models.Blog
	for rows.Next() {
		var b models.Blog
		if err := rows.Scan(&b.BlogID, &b.UserID, &b.CategoryID, &b.Title, &b.Description, &b.Content, &b.FilePath, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}
		blogs = append(blogs, b)
	}
	return blogs, rows.Err()
}


func GetAllBlogs(db *sql.DB) ([]models.Blog, error) {
	rows, err := db.Query(
		`SELECT blog_id, user_id, category_id, title, description, content, file_path, created_at, updated_at
		 FROM blogs ORDER BY created_at DESC`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []models.Blog
	for rows.Next() {
		var b models.Blog
		if err := rows.Scan(&b.BlogID, &b.UserID, &b.CategoryID, &b.Title, &b.Description, &b.Content, &b.FilePath, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}
		blogs = append(blogs, b)
	}
	return blogs, rows.Err()
}

func UpdateBlog(db *sql.DB, blogID int, title, description, content string) error {
	result, err := db.Exec(
		`UPDATE blogs SET title = $1, description = $2, content = $3, updated_at = CURRENT_TIMESTAMP
		 WHERE blog_id = $4`,
		title, description, content, blogID,
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

func DeleteBlog(db *sql.DB, blogID int) error {
	result, err := db.Exec("DELETE FROM blogs WHERE blog_id = $1", blogID)
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