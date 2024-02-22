package controller

import (
	"context"
	"errors"
	"service-user/helpers"
	"service-user/model"

	"service-user/config"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type WebResponse struct {
	Code int
	Status string
	Data interface{}
}

func Register(c *fiber.Ctx) error {
	var requestBody model.User
	db := config.GetMongoDatabase().Collection("user")

	requestBody.Id = uuid.New().String()

	ctx, cancel := config.NewMongoContext()
	defer cancel()

	c.BodyParser(&requestBody)

	_, err := db.InsertOne(ctx, bson.M{
		"email": requestBody.Email,
		"password": helpers.HashPassword([]byte(requestBody.Password)),
	})

	if err != nil {
		panic(err)
	}

	return c.JSON(WebResponse{
		Code: 201,
		Status: "OK",
		Data: requestBody.Email,
	})
}

func Login(c *fiber.Ctx) error {
	db := config.GetMongoDatabase().Collection("user")

	var requestBody model.User
	var result model.User
 
	c.BodyParser(&requestBody)

	err := db.FindOne(context.TODO(), bson.D{{"email", requestBody.Email}}).Decode(&result)
	if err != nil {
		return c.JSON(WebResponse{
			Code: 401,
			Status: "BAD_REQUEST",
			Data: err.Error(),
		})
	}

	checkPassword := helpers.ComparePassword([]byte(result.Password), []byte(requestBody.Password))
	if !checkPassword {
		return c.JSON(WebResponse{
			Code: 401,
			Status: "BAD_REQUEST",
			Data: errors.New("invalid password").Error(),
		})
	}

	access_token := helpers.SignToken(requestBody.Email)

	return c.JSON(struct{
		Code int 
		Status string
		AccessToken string
		Data interface{}
	}{
		Code: 200,
		Status: "OK",
		AccessToken: access_token,
		Data: result,
	})
}

func Auth(c *fiber.Ctx) error {
	return c.JSON("OK")
}
