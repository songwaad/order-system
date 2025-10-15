package handlers

import (
	"kornkk/api/presenter"
	"kornkk/entities"
	"kornkk/usecases/user"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func AddUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input entities.RegisterInput
		err := c.BodyParser(&input)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		createdUser, err := service.Register(&input)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		return c.JSON(presenter.UserSuccessResponse(createdUser))
	}
}

func GetUserByID(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		user, err := service.GetUserByID(uint(id))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		return c.JSON(presenter.UserSuccessResponse(user))
	}

}

func DeleteUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		err = service.DeleteUser(uint(id))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "User deleted successfully",
			"error":  nil,
		})
	}
}
