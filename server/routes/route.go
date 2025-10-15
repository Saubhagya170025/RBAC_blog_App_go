// routes/routes.go
package routes

import (
	"database/sql"

	"github.com/Saubhagya170025/rbac-blog-app/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	// Health check route
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// User Routes
	userRoutes := app.Group("/api/users")
	{
		userRoutes.Post("", handlers.CreateUserHandler(db))			//ok user is gettting created but refresh token is not getting stored on passing in  req.body 
		userRoutes.Get("/:id", handlers.GetUserHandler(db))			//ok
		userRoutes.Get("", handlers.GetAllUsersHandler(db))		     //ok
		userRoutes.Put("/:id", handlers.UpdateUserHandler(db))	      //ok but you have to pass the role id also to update
		userRoutes.Delete("/:id", handlers.DeleteUserHandler(db))		//ok
		userRoutes.Post("/login", handlers.LoginUserHandler(db))         //❌❌ not working
	}

	// Role Routes
	roleRoutes := app.Group("/api/roles")
	{
		roleRoutes.Post("", handlers.CreateRoleWithPermsHandler(db))    //ok
		roleRoutes.Get("/:id", handlers.GetRoleHandler(db))				//ok but i also want to fetch the permissions associated with the role
		roleRoutes.Get("", handlers.GetAllRolesHandler(db))				//ok
		roleRoutes.Delete("/:id", handlers.DeleteRoleHandler(db))		//ok  its great that ondeleting roles its associated permissions are also getting deleted 
		//you should create a common handler to update role name as well as associated permissions
	}

	// Permission Routes
	permissionRoutes := app.Group("/api/permissions")
	{
		// permissionRoutes.Post("", handlers.CreatePermissionHandler(db))
		permissionRoutes.Get("/role/:roleId", handlers.GetPermissionByRoleHandler(db))		//ok
		permissionRoutes.Put("/:id", handlers.UpdatePermissionHandler(db))					//error
	}

	// Category Routes
	categoryRoutes := app.Group("/api/categories")
	{
		categoryRoutes.Post("", handlers.CreateCategoryHandler(db))							//ok
		categoryRoutes.Get("/:id", handlers.GetCategoryHandler(db))							//ok
		categoryRoutes.Get("", handlers.GetAllCategoriesHandler(db))        				//ok
		categoryRoutes.Put("/:id", handlers.UpdateCategoryHandler(db)) 						//ok
		categoryRoutes.Delete("/:id", handlers.DeleteCategoryHandler(db))   				//ok
	}

	// Blog Routes
	blogRoutes := app.Group("/api/blogs")
	{
		blogRoutes.Post("", handlers.CreateBlogHandler(db))									//ok
		blogRoutes.Get("/:id", handlers.GetBlogHandler(db))									//ok 
		blogRoutes.Get("", handlers.GetAllBlogsHandler(db))									//ok
		blogRoutes.Put("/:id", handlers.UpdateBlogHandler(db))								//ok
		blogRoutes.Delete("/:id", handlers.DeleteBlogHandler(db))							//ok
		blogRoutes.Get("/user/:userId", handlers.GetBlogsByUserHandler(db))					//ok
		blogRoutes.Get("/category/:categoryId", handlers.GetBlogsByCategoryHandler(db))		//ok
	}
}
