// routes/routes.go
package routes

import (
	"database/sql"

	"github.com/Saubhagya170025/rbac-blog-app/handlers"
	"github.com/Saubhagya170025/rbac-blog-app/config"
	"github.com/Saubhagya170025/rbac-blog-app/middlewares"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, db *sql.DB, cfg *config.Config) {
	// Health check route (public)
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// API group
	api := app.Group("/api")

	// ===========================================
	// PUBLIC ROUTES (No Authentication Required)
	// ===========================================
	
	// Auth Routes (Public)
	authRoutes := api.Group("/auth")
	{
		authRoutes.Post("/register", handlers.CreateUserHandler(db))
		authRoutes.Post("/login", handlers.LoginHandler(db, cfg))
		authRoutes.Post("/refresh", handlers.RefreshTokenHandler(db, cfg))
	}

	// ===========================================
	// PROTECTED ROUTES (Authentication Required)
	// ===========================================
	
	// Apply auth middleware to all routes below
	protected := api.Group("/", middleware.AuthMiddleware(cfg))

	// Auth Routes (Protected)
	protectedAuth := protected.Group("/auth")
	{
		protectedAuth.Post("/logout", handlers.LogoutHandler(db))
		protectedAuth.Get("/validate", handlers.ValidateHandler(db))
	}

	// User Routes (Protected)
	userRoutes := protected.Group("/users")
	{
		userRoutes.Get("/:id", handlers.GetUserHandler(db))           // ok
		userRoutes.Get("", handlers.GetAllUsersHandler(db))           // ok
		userRoutes.Put("/:id", handlers.UpdateUserHandler(db))        // ok but you have to pass the role id also to update
		userRoutes.Delete("/:id", handlers.DeleteUserHandler(db))     // ok
	}

	// Role Routes (Protected)
	roleRoutes := protected.Group("/roles")
	{
		roleRoutes.Post("", handlers.CreateRoleWithPermsHandler(db))    // ok
		roleRoutes.Get("/:id", handlers.GetRoleHandler(db))             // ok but i also want to fetch the permissions associated with the role
		roleRoutes.Get("", handlers.GetAllRolesHandler(db))             // ok
		roleRoutes.Delete("/:id", handlers.DeleteRoleHandler(db))       // ok  its great that ondeleting roles its associated permissions are also getting deleted
		// you should create a common handler to update role name as well as associated permissions
	}

	// Permission Routes (Protected)
	permissionRoutes := protected.Group("/permissions")
	{
		// permissionRoutes.Post("", handlers.CreatePermissionHandler(db))
		permissionRoutes.Get("/role/:roleId", handlers.GetPermissionByRoleHandler(db))  // ok
		permissionRoutes.Put("/:id", handlers.UpdatePermissionHandler(db))              // error
	}

	// Category Routes (Protected)
	categoryRoutes := protected.Group("/categories")
	{
		categoryRoutes.Post("", handlers.CreateCategoryHandler(db))             // ok
		categoryRoutes.Get("/:id", handlers.GetCategoryHandler(db))             // ok
		categoryRoutes.Get("", handlers.GetAllCategoriesHandler(db))            // ok
		categoryRoutes.Put("/:id", handlers.UpdateCategoryHandler(db))          // ok
		categoryRoutes.Delete("/:id", handlers.DeleteCategoryHandler(db))       // ok
	}

	// Blog Routes (Protected)
	blogRoutes := protected.Group("/blogs")
	{
		blogRoutes.Post("", handlers.CreateBlogHandler(db))                                 // ok
		blogRoutes.Get("/:id", handlers.GetBlogHandler(db))                                 // ok
		blogRoutes.Get("", handlers.GetAllBlogsHandler(db))                                 // ok
		blogRoutes.Put("/:id", handlers.UpdateBlogHandler(db))                              // ok
		blogRoutes.Delete("/:id", handlers.DeleteBlogHandler(db))                           // ok
		blogRoutes.Get("/user/:userId", handlers.GetBlogsByUserHandler(db))                 // ok
		blogRoutes.Get("/category/:categoryId", handlers.GetBlogsByCategoryHandler(db))     // ok
	}
}