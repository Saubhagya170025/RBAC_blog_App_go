package handlers

import (
	"database/sql"
	"strconv"

	"github.com/Saubhagya170025/rbac-blog-app/database/repository"
	"github.com/gofiber/fiber/v3"
)

func CreateUserHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
			RoleID   int    `json:"role_id"`
		}

		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		userID, err := repository.CreateUser(db, req.Name, req.Email, req.Password, req.RoleID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"user_id": userID,
			"message": "User created successfully",
		})
	}
}

func GetUserHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		user, err := repository.GetUserByID(db, userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if user == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		return c.JSON(user)
	}
}

func GetAllUsersHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		users, err := repository.GetAllUsers(db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(users)
	}
}

func UpdateUserHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		var req struct {
			Name   string `json:"name"`
			Email  string `json:"email"`
			RoleID int    `json:"role_id"`
		}

		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		err = repository.UpdateUser(db, userID, req.Name, req.Email, req.RoleID)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "User not found",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "User updated successfully",
		})
	}
}

func DeleteUserHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		userID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		err = repository.DeleteUser(db, userID)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "User not found",
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "User deleted successfully",
		})
	}
}

func LoginUserHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		user, err := repository.GetUserByEmail(db, req.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid email or password",
			})
		}

		// TODO: Add password verification and JWT token generation
		return c.JSON(fiber.Map{
			"message": "Login successful",
			"user_id": user.UserID,
		})
	}
}