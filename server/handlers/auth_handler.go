package handlers

import (
	"database/sql"
	"github.com/Saubhagya170025/rbac-blog-app/config"
	"github.com/Saubhagya170025/rbac-blog-app/database/repository"
	"github.com/Saubhagya170025/rbac-blog-app/utils"

	"github.com/gofiber/fiber/v3"
)

// LoginHandler handles user login
func LoginHandler(db *sql.DB, cfg *config.Config) fiber.Handler {
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

		// Get user by email
		found_user, err := repository.GetUserByEmail(db, req.Email)
		// userID, name, hashedPassword, roleID, err := repository.GetUserByEmail(db, req.Email)
		userID := found_user.UserID
		name := found_user.Name
		hashedPassword := found_user.Password
		roleID := found_user.RoleID

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid credentials",
			})
		}

		// Check password
		if err := utils.CheckPassword(hashedPassword, req.Password); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid credentials",
			})
		}

		// Generate access token
		accessToken, err := utils.GenerateAccessToken(
			userID, req.Email, roleID,
			cfg.JWTAccessSecret,
			cfg.AccessTokenExpiry,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate access token",
			})
		}

		// Generate refresh token
		refreshToken, err := utils.GenerateRefreshToken(
			userID,
			cfg.JWTRefreshSecret,
			cfg.RefreshTokenExpiry,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate refresh token",
			})
		}

		// Store refresh token in database
		if err := repository.StoreRefreshToken(db, userID, refreshToken); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to store refresh token",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":       "Login successful",
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"user": fiber.Map{
				"user_id": userID,
				"name":    name,
				"email":   req.Email,
				"role_id": roleID,
			},
		})
	}
}

// RefreshTokenHandler generates new access token using refresh token
func RefreshTokenHandler(db *sql.DB, cfg *config.Config) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req struct {
			RefreshToken string `json:"refresh_token"`
		}

		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Validate refresh token
		claims, err := utils.ValidateToken(req.RefreshToken, cfg.JWTRefreshSecret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid refresh token",
			})
		}

		// Check if refresh token matches the one in database
		storedToken, err := repository.GetRefreshToken(db, claims.UserID)
		if err != nil || storedToken != req.RefreshToken {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid refresh token",
			})
		}

		// Get user details
		user_by_id, err := repository.GetUserByID(db, claims.UserID)
		// name, email, roleID, err := repository.GetUserByID(db, claims.UserID)
		name := user_by_id.Name
		email := user_by_id.Email
		roleID := user_by_id.RoleID

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found",
			})
		}

		// Generate new access token
		accessToken, err := utils.GenerateAccessToken(
			claims.UserID, email, roleID,
			cfg.JWTAccessSecret,
			cfg.AccessTokenExpiry,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate access token",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":      "Token refreshed successfully",
			"access_token": accessToken,
			"user": fiber.Map{
				"user_id": claims.UserID,
				"name":    name,
				"email":   email,
				"role_id": roleID,
			},
		})
	}
}

// LogoutHandler handles user logout
func LogoutHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Get user ID from context (set by auth middleware)
		userID := c.Locals("user_id").(int)

		// Delete refresh token from database
		if err := repository.DeleteRefreshToken(db, userID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to logout",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Logout successful",
		})
	}
}