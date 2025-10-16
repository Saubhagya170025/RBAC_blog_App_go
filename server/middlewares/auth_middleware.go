package middleware

import (
	"github.com/Saubhagya170025/rbac-blog-app/config"
	"github.com/Saubhagya170025/rbac-blog-app/utils"
	"strings"

	"github.com/gofiber/fiber/v3"
)

// AuthMiddleware validates JWT access token
func AuthMiddleware(cfg *config.Config) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Get Authorization header
		// Prefer Authorization header, fall back to access_token cookie (HTTPOnly)
		authHeader := c.Get("Authorization")
		var tokenString string
		if authHeader != "" {
			// Check Bearer format
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid authorization header format"})
			}
			tokenString = parts[1]
		} else {
			// Try cookie
			tokenString = c.Cookies("access_token")
			if tokenString == "" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing authorization token"})
			}
		}

		// Validate token
		claims, err := utils.ValidateToken(tokenString, cfg.JWTAccessSecret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// Store user information in context for use in handlers
		c.Locals("user_id", claims.UserID)
		c.Locals("email", claims.Email)
		c.Locals("role_id", claims.RoleID)

		return c.Next()
	}
}

// OptionalAuthMiddleware validates token if present, but doesn't require it
func OptionalAuthMiddleware(cfg *config.Config) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Next()
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			tokenString := parts[1]
			if claims, err := utils.ValidateToken(tokenString, cfg.JWTAccessSecret); err == nil {
				c.Locals("user_id", claims.UserID)
				c.Locals("email", claims.Email)
				c.Locals("role_id", claims.RoleID)
			}
		}

		return c.Next()
	}
}