package handlers

import (
	"database/sql"
	"strconv"

	"github.com/Saubhagya170025/rbac-blog-app/database/repository"
	"github.com/gofiber/fiber/v3"
)

func CreateBlogHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req struct {
			UserID      int    `json:"user_id"`
			CategoryID  int    `json:"category_id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Content     string `json:"content"`
			FilePath    string `json:"file_path"`
		}

		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		blogID, err := repository.CreateBlog(db, req.UserID, req.CategoryID, req.Title, req.Description, req.Content, req.FilePath)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"blog_id": blogID,
			"message": "Blog created successfully",
		})
	}
}

func GetBlogHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		blogID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid blog ID",
			})
		}

		blog, err := repository.GetBlogByID(db, blogID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if blog == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Blog not found",
			})
		}

		return c.JSON(blog)
	}
}

func GetAllBlogsHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		blogs, err := repository.GetAllBlogs(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(blogs)
	}
}

func UpdateBlogHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		blogID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid blog ID",
			})
		}

		var req struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Content     string `json:"content"`
		}

		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		err = repository.UpdateBlog(db, blogID, req.Title, req.Description, req.Content)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Blog not found",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "Blog updated successfully",
		})
	}
}

func DeleteBlogHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		blogID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid blog ID",
			})
		}

		err = repository.DeleteBlog(db, blogID)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Blog not found",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "Blog deleted successfully",
		})
	}
}

func GetBlogsByUserHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := strconv.Atoi(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		blogs, err := repository.GetBlogsByUserID(db, userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(blogs)
	}
}

func GetBlogsByCategoryHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		categoryID, err := strconv.Atoi(c.Params("categoryId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid category ID",
			})
		}

		blogs, err := repository.GetBlogsByCategoryID(db, categoryID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(blogs)
	}
}