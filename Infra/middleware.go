package infra

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: jwtError,
		// อ่าน JWT จาก cookie "jwt" ก่อน, ถ้าไม่มีลอง Authorization header
		TokenLookup: "cookie:jwt,header:Authorization:Bearer ",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": false,
			"data":   nil,
			"error":  "Missing or malformed JWT",
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status": false,
		"data":   nil,
		"error":  "Invalid or expired JWT",
	})
}
