package model

type Employee struct {
	Id   string `json:"id" bson:"_id"`
	Name string `json:"name"`
}