package routes

import (
	infra "kornkk/Infra"
	"kornkk/api/handlers"
	"kornkk/usecases/auth"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app fiber.Router, service auth.Service) {
	app.Post("/login", handlers.Login(service))
	app.Post("/logout", handlers.Logout())

	app.Get("/me", infra.Protected(), handlers.Me())
}
