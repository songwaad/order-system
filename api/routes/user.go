package routes

import (
	"kornkk/api/handlers"
	"kornkk/usecases/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app fiber.Router, service user.Service) {
	app.Post("/users", handlers.AddUser(service))
}
