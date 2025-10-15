package main

import (
	"fmt"
	"log"

	"github.com/Saubhagya170025/rbac-blog-app/database"
	"github.com/Saubhagya170025/rbac-blog-app/routes"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v3"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "rbac_final"
)

func main() {

	// ----------------------------Databse Connection ---------------------------------That has to be placed in a separate file
	// This is a simple example of connecting to a PostgreSQL database using Go's database/sql package.
	// Make sure to import the pq driver with _ "github.com/lib/pq" to register it.
	// You can use the following code to connect to your PostgreSQL database.

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := database.Connect(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Successfully connected with database")

	// --------------------migrations-------------------------------------

	err = database.RunMigrations(db, "./migrations")
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// -----------------------routes----------------------------------------------
	// Create Fiber app
	app := fiber.New()

	// Adding CORS middleware with specific origin
	app.Use(cors.New(cors.Config{
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Content-Type,Authorization,Accept,Origin,Access-Control-Request-Method,Access-Control-Request-Headers,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Access-Control-Allow-Methods,Access-Control-Expose-Headers,Access-Control-Max-Age,Access-Control-Allow-Credentials",
		AllowCredentials: true,
	}))
	// Setup routes
	routes.SetupRoutes(app, db)
	// -----------------------Starting the server----------------------------------------------

	log.Printf("Server started on 0.0.0.0:8080")
	if err := app.Listen("0.0.0.0:8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
