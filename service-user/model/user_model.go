package model

type User struct {
	Id       string `json:"id" bson:"_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}