package main

import (
	"app/database"
	"app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// Database connection
	database.Connect()

	// Creating a new fiber handler
	app := fiber.New()

	// Cors
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Set up the routes
	routes.Setup(app)

	// Start the Server
	err := app.Listen(":8000")
	if err != nil {
		return
	}
}
