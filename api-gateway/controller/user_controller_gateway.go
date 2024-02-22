package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var user_uri string = "http://localhost:3001/user"

type UserBodyReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code        int    `json:"Code"`
	Status      string `json:"Status"`
	AccessToken string `json:"AccessToken"`
	Data        struct {
		ID       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"Data"`
}

func UserLogin(c *fiber.Ctx) error {
	var bodyRequest UserBodyReq
	c.BodyParser(&bodyRequest)

	payload, err := json.Marshal(bodyRequest)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(user_uri+"/login", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println("Response Status:", resp.Status)
	// fmt.Println("Response Body:", string(body))

	var res LoginResponse

	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}

	return c.JSON(res)
}