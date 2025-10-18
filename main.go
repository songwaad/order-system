package main

import (
	"fmt"
	"kornkk/api/routes"
	"kornkk/database"
	"kornkk/entities"
	"kornkk/usecases/auth"
	"kornkk/usecases/user"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.GetDB()
	if err != nil {
		log.Fatal("Database Connection Error:", err)
	}

	fmt.Println("Database Connection successfully!")

	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.UserRole{})

	userRepo := user.NewRepo(db)
	userService := user.NewService(userRepo)

	authService := auth.NewService(userRepo)

	app := fiber.New()
	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     "http://localhost:5173",
			AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
			AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
			ExposeHeaders:    "Content-Length",
			MaxAge:           300,
			AllowCredentials: true,
		},
	))
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Korn KK"))
	})

	api := app.Group("/api")
	routes.UserRoute(api, userService)
	routes.AuthRoute(api, authService)
	log.Fatal(app.Listen(":8080"))

}
