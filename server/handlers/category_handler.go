package handlers

import (
	"database/sql"
	"strconv"

	"github.com/Saubhagya170025/rbac-blog-app/database/repository"
	"github.com/gofiber/fiber/v3"
)

func CreateCategoryHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req struct {
			CategoryName string `json:"category_name"`
		}

		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		categoryID, err := repository.CreateCategory(db, req.CategoryName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"category_id": categoryID,
			"message":     "Category created successfully",
		})
	}
}

func GetCategoryHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		categoryID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid category ID",
			})
		}

		category, err := repository.GetCategoryByID(db, categoryID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if category == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Category not found",
			})
		}

		return c.JSON(category)
	}
}

func GetAllCategoriesHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		categories, err := repository.GetAllCategories(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(categories)
	}
}

func UpdateCategoryHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		categoryID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid category ID",
			})
		}

		var req struct {
			CategoryName string `json:"category_name"`
		}

		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		err = repository.UpdateCategory(db, categoryID, req.CategoryName)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Category not found",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "Category updated successfully",
		})
	}
}

func DeleteCategoryHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		categoryID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid category ID",
			})
		}

		err = repository.DeleteCategory(db, categoryID)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Category not found",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "Category deleted successfully",
		})
	}
}