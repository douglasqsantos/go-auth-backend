package routes

import (
	"app/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App){
	app.Get("/", controllers.Home)
}

