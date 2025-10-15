package handlers

import (
	"database/sql"
	"strconv"

	"github.com/Saubhagya170025/rbac-blog-app/database/repository"
	"github.com/gofiber/fiber/v3"
)

func CreateRoleHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req struct {
			RoleName string `json:"role_name"`
		}

		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		roleID, err := repository.CreateRole(db, req.RoleName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"role_id": roleID,
			"message": "Role created successfully",
		})
	}
}

func GetRoleHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		roleID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid role ID",
			})
		}

		role, err := repository.GetRoleByID(db, roleID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if role == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Role not found",
			})
		}

		return c.JSON(role)
	}
}

func GetAllRolesHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		roles, err := repository.GetAllRoles(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(roles)
	}
}

func DeleteRoleHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		roleID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid role ID",
			})
		}

		err = repository.DeleteRole(db, roleID)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "Role not found",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "Role deleted successfully",
		})
	}
}