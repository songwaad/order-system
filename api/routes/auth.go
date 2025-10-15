package routes

import (
	"kornkk/api/handlers"
	"kornkk/usecases/auth"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app fiber.Router, service auth.Service) {
	app.Post("/login", handlers.Login(service))
}
