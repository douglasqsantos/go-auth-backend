package main

import (
	"app/database"
	"app/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// Database connection
	database.Connect()

	// Creating a new fiber handler
	app := fiber.New()

	// Set up the routes
	routes.Setup(app)

	// Start the Server
	err := app.Listen(":8000")
	if err != nil {
		return 
	}
}
