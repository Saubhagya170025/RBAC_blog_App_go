package handlers

import (
	"database/sql"
	"strconv"

	"github.com/Saubhagya170025/rbac-blog-app/database/repository"
	"github.com/Saubhagya170025/rbac-blog-app/models"
	"github.com/gofiber/fiber/v3"
)

func CreatePermissionHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var perm models.Permission
		if err := c.Bind().Body(&perm); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		permID, err := repository.CreatePermission(db, &perm)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"permission_id": permID,
			"message":       "Permission created successfully",
		})
	}
}

func GetPermissionByRoleHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		roleID, err := strconv.Atoi(c.Params("roleId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid role ID",
			})
		}

		perm, err := repository.GetPermissionByRoleID(db, roleID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if perm == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Permission not found for this role",
			})
		}

		return c.JSON(perm)
	}
}

func UpdatePermissionHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		permID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid permission ID",
			})
		}

		var perm models.Permission
		if err := c.Bind().Body(&perm); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		err = repository.UpdatePermission(db, permID, &perm)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Permission not found",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "Permission updated successfully",
		})
	}
}