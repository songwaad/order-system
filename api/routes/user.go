package routes

import (
	infra "kornkk/Infra"
	"kornkk/api/handlers"
	"kornkk/usecases/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app fiber.Router, service user.Service) {
	app.Post("/users", handlers.AddUser(service))

	app.Get("/users", infra.Protected(), handlers.GetAllUsers(service))
	app.Get("/users/:id", infra.Protected(), handlers.GetUserByID(service))
	app.Patch("/users/:id", infra.Protected(), handlers.UpdateUser(service))
	app.Delete("/users/:id", infra.Protected(), handlers.DeleteUser(service))
}
