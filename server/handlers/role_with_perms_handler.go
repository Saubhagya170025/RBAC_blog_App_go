package handlers

import (
	"database/sql"

	"github.com/Saubhagya170025/rbac-blog-app/database/repository"
	"github.com/Saubhagya170025/rbac-blog-app/models"
	"github.com/gofiber/fiber/v3"
)

// CreateRoleWithPermsRequest is the payload to create a role and its permissions together.
type CreateRoleWithPermsRequest struct {
	RoleName    string            `json:"role_name"`
	Permissions models.Permission `json:"permissions"`
}

// CreateRoleWithPermsHandler creates a role and its permissions. If permission creation fails,
// it attempts to delete the created role as a rollback (best-effort since repository functions
// don't accept a transaction object).
func CreateRoleWithPermsHandler(db *sql.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		var req CreateRoleWithPermsRequest
		if err := c.Bind().Body(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}
		if req.RoleName == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "role_name is required"})
		}

		// 1) Create role
		roleID, err := repository.CreateRole(db, req.RoleName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// 2) Create permission record tied to role
		perm := req.Permissions
		perm.RoleID = roleID
		permID, err := repository.CreatePermission(db, &perm)
		if err != nil {
			// best-effort rollback: delete the role we just created
			_ = repository.DeleteRole(db, roleID)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create permission, role creation rolled back"})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"role_id":       roleID,
			"permission_id": permID,
			"message":       "Role and permissions created successfully",
		})
	}
}
