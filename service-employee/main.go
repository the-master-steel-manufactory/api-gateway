package main

import (
	"fmt"
	"service-employee/config"
	"service-employee/controller"

	"github.com/gofiber/fiber/v2"
)

func init() {
	config.GetMongoDatabase()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi from service-employee")
	})
	app.Post("/employee", controller.CreateEmployee)

	port := 3002
	fmt.Printf("Service employee is running on :%d...\n", port)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting Service employee: %v\n", err)
	}
}
