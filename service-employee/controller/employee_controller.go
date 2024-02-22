package controller

import (
	"fmt"
	"net/http"
	"service-employee/config"
	"service-employee/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

var user_uri string = "http://localhost:3001/user"

type WebResponse struct {
	Code int
	Status string
	Data interface{}
}

func CreateEmployee(c *fiber.Ctx) error {
	db := config.GetMongoDatabase().Collection("employee")
	var requestBody model.Employee

	c.BodyParser(&requestBody)

	requestBody.Id = uuid.New().String()

	access_token := c.Get("access_token")
	if len(access_token) == 0 {
		return c.Status(401).SendString("Invalid token: Access token missing")
	}

	req, err := http.NewRequest("GET", user_uri + "/auth", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		panic(err)
	}

	// Set headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access_token", access_token)

	// Send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		panic(err)
	}
	defer resp.Body.Close()

	// Print the response
	// fmt.Println("Response Status:", resp.Status)
	// fmt.Println("Response Headers:", resp.Header)

	if resp.Status != "200 OK" {
		c.Status(401).SendString("invalid token")
	}

	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err = db.InsertOne(ctx, bson.M{
		"name": requestBody.Name,
	})

	if err != nil {
		panic(err)
	}

	return c.JSON(WebResponse{
		Code: 201,
		Status: "OK",
		Data: requestBody,
	})
}