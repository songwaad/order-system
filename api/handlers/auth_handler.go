package handlers

import (
	"kornkk/api/presenter"
	"kornkk/entities"
	"kornkk/usecases/auth"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

		// Set JWT token in HTTP-only cookie
		cookie := new(fiber.Cookie)
		cookie.Name = "jwt"
		cookie.Value = token
		cookie.HTTPOnly = true                          // Prevent JavaScript access (XSS protection)
		cookie.Secure = false                           // Set to true in production with HTTPS
		cookie.SameSite = "Lax"                         // CSRF protection
		cookie.Expires = time.Now().Add(24 * time.Hour) // Match JWT expiration
		cookie.Path = "/"                               // Available for all routes

		c.Cookie(cookie)

		return c.JSON(&fiber.Map{
			"status": true,
			"data": &fiber.Map{
				"message": "Login successful",
			},
			"error": nil,
		})
	}
}

func Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Clear the JWT cookie
		cookie := new(fiber.Cookie)
		cookie.Name = "jwt"
		cookie.Value = ""
		cookie.HTTPOnly = true
		cookie.Secure = false // Set to true in production with HTTPS
		cookie.SameSite = "Lax"
		cookie.Expires = time.Now().Add(-1 * time.Hour) // Expire immediately
		cookie.Path = "/"

		c.Cookie(cookie)

		return c.JSON(&fiber.Map{
			"status": true,
			"data": &fiber.Map{
				"message": "Logout successful",
			},
			"error": nil,
		})
	}
}

func Me() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userToken := c.Locals("user").(*jwt.Token)
		claims := userToken.Claims.(jwt.MapClaims)

		return c.JSON(&fiber.Map{
			"status": true,
			"data": &fiber.Map{
				"user": claims,
			},
			"error": nil,
		})
	}
}
