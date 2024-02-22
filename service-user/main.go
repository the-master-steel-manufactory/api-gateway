package main

import (
	"fmt"
	"service-user/config"
	"service-user/controller"
	"service-user/middleware"

	"github.com/gofiber/fiber/v2"
)

func init() {
	config.GetMongoDatabase()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi from service-user")
	})
	app.Post("/user/register", controller.Register)
	app.Post("/user/login", controller.Login)
	app.Get("/user/auth", middleware.Authentication, controller.Auth)

	port := 3001
	fmt.Printf("Service user is running on :%d...\n", port)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting Service user: %v\n", err)
	}
}
