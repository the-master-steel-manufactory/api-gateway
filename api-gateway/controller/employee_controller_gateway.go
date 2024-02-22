package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var employee_uri string = "http://localhost:3002"

type EmployeeBodyReq struct {
	Name string `json:"name"`
}

type EmployeeResponse struct {
	Code   int    `json:"Code"`
	Status string `json:"Status"`
	Data   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"Data"`
}

func CreateEmployee(c *fiber.Ctx) error {
	var bodyRequest EmployeeBodyReq
	c.BodyParser(&bodyRequest)

	payload, err := json.Marshal(bodyRequest)
	if err != nil {
		panic(err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", employee_uri+"/employee", bytes.NewBuffer(payload))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating HTTP request")
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")

	access_token := c.Get("access_token")
	if len(access_token) == 0 {
		return c.Status(401).SendString("Invalid token: Access token missing")
	}

	// Add additional headers as needed
	req.Header.Set("access_token", access_token)

	// Make the HTTP POST request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error making HTTP POST request")
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusInternalServerError).SendString("Non-OK status code received")
	}

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error reading response body")
	}

	var res EmployeeResponse
    err = json.Unmarshal(responseBody, &res)
    if err != nil {
        return err
    }

    return c.JSON(res)
}