package presenter

import (
	"kornkk/entities"

	"github.com/gofiber/fiber/v2"
)

func UserSuccessResponse(data *entities.User) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func UsersSuccessResponse(data *[]entities.User) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
