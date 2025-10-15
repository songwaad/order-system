package handlers

import (
	"kornkk/api/presenter"
	"kornkk/entities"
	"kornkk/usecases/auth"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Login(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input entities.LoginInput
		err := c.BodyParser(&input)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		token, err := service.Login(input.Identity, input.Password)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		return c.JSON(&fiber.Map{
			"status": true,
			"data": &fiber.Map{
				"token": token,
			},
			"error": nil,
		})
	}
}
