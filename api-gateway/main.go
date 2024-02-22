package main

import (
	"api-gateway/controller"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi from api gateway")
	})
	app.Post("/login", controller.UserLogin)
	app.Post("/employee", controller.CreateEmployee)

	port := 3000
	fmt.Printf("api gateway is running on :%d...\n", port)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting api gateway: %v\n", err)
	}
}
