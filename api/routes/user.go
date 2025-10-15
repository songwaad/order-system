package routes

import (
	"kornkk/api/handlers"
	"kornkk/usecases/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app fiber.Router, service user.Service) {
	app.Post("/users", handlers.AddUser(service))
	app.Get("/users", handlers.GetAllUsers(service))
	app.Get("/users/:id", handlers.GetUserByID(service))
	app.Patch("/users/:id", handlers.UpdateUser(service))
	app.Delete("/users/:id", handlers.DeleteUser(service))
}
