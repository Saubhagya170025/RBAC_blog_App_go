package handlers

import (
	"database/sql"
	"strconv"

	// "github.com/Saubhagya170025/rbac-blog-app/config"
	"github.com/Saubhagya170025/rbac-blog-app/database/repository"
	// "github.com/Saubhagya170025/rbac-blog-app/utils"
	// "github.com/Saubhagya170025/rbac-blog-app/handlers"
	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

// CreateUserHandler - Register new user (used for /api/auth/register)
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

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to process password",
			})
		}

		// Pass hashed password to repository
		userID, err := repository.CreateUser(db, req.Name, req.Email, string(hashedPassword), req.RoleID)
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

// ============================================
// AUTH HANDLERS (JWT Authentication)
// ============================================

// // LoginHandler handles user login with JWT token generation
// func LoginHandler(db *sql.DB, cfg *config.Config) fiber.Handler {
// 	return func(c fiber.Ctx) error {
// 		var req struct {
// 			Email    string `json:"email"`
// 			Password string `json:"password"`
// 		}

// 		if err := c.Bind().Body(&req); err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "Invalid request body",
// 			})
// 		}

// 		// Get user by email
// 		user, err := repository.GetUserByEmail(db, req.Email)
// 		if err != nil {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Invalid email or password",
// 			})
// 		}

// 		if user == nil {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Invalid email or password",
// 			})
// 		}

// 		// Verify password
// 		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Invalid email or password",
// 			})
// 		}

// 		// Generate access token
// 		accessToken, err := utils.GenerateAccessToken(
// 			user.UserID,
// 			user.Email,
// 			user.RoleID,
// 			cfg.JWTAccessSecret,
// 			cfg.AccessTokenExpiry,
// 		)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Failed to generate access token",
// 			})
// 		}

// 		// Generate refresh token
// 		refreshToken, err := utils.GenerateRefreshToken(
// 			user.UserID,
// 			cfg.JWTRefreshSecret,
// 			cfg.RefreshTokenExpiry,
// 		)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Failed to generate refresh token",
// 			})
// 		}

// 		// Store refresh token in database
// 		if err := repository.StoreRefreshToken(db, user.UserID, refreshToken); err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Failed to store refresh token",
// 			})
// 		}

// 		return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 			"message":       "Login successful",
// 			"access_token":  accessToken,
// 			"refresh_token": refreshToken,
// 			"user": fiber.Map{
// 				"user_id": user.UserID,
// 				"name":    user.Name,
// 				"email":   user.Email,
// 				"role_id": user.RoleID,
// 			},
// 		})
// 	}
// }

// // RefreshTokenHandler generates new access token using refresh token
// func RefreshTokenHandler(db *sql.DB, cfg *config.Config) fiber.Handler {
// 	return func(c fiber.Ctx) error {
// 		var req struct {
// 			RefreshToken string `json:"refresh_token"`
// 		}

// 		if err := c.Bind().Body(&req); err != nil {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "Invalid request body",
// 			})
// 		}

// 		// Validate refresh token
// 		claims, err := utils.ValidateToken(req.RefreshToken, cfg.JWTRefreshSecret)
// 		if err != nil {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Invalid refresh token",
// 			})
// 		}

// 		// Check if refresh token matches the one in database
// 		storedToken, err := repository.GetRefreshToken(db, claims.UserID)
// 		if err != nil || storedToken != req.RefreshToken {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Invalid refresh token",
// 			})
// 		}

// 		// Get user details
// 		user, err := repository.GetUserByID(db, claims.UserID)
// 		if err != nil || user == nil {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "User not found",
// 			})
// 		}

// 		// Generate new access token
// 		accessToken, err := utils.GenerateAccessToken(
// 			user.UserID,
// 			user.Email,
// 			user.RoleID,
// 			cfg.JWTAccessSecret,
// 			cfg.AccessTokenExpiry,
// 		)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Failed to generate access token",
// 			})
// 		}

// 		return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 			"message":      "Token refreshed successfully",
// 			"access_token": accessToken,
// 			"user": fiber.Map{
// 				"user_id": user.UserID,
// 				"name":    user.Name,
// 				"email":   user.Email,
// 				"role_id": user.RoleID,
// 			},
// 		})
// 	}
// }

// // LogoutHandler handles user logout (invalidates refresh token)
// func LogoutHandler(db *sql.DB) fiber.Handler {
// 	return func(c fiber.Ctx) error {
// 		// Get user ID from context (set by auth middleware)
// 		userID, ok := c.Locals("user_id").(int)
// 		if !ok {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Unauthorized",
// 			})
// 		}

// 		// Delete refresh token from database
// 		if err := repository.DeleteRefreshToken(db, userID); err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Failed to logout",
// 			})
// 		}

// 		return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 			"message": "Logout successful",
// 		})
// 	}
// }