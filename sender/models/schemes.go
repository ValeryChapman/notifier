package models

import "time"

type Notification struct {
	Id        string    `json:"id" bson:"_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	To        []string  `json:"to" bson:"to" binding:"required,min=1,max=100"`
	Subject   string    `json:"subject" bson:"subject" binding:"required,min=1,max=250"`
	Body      string    `json:"body" bson:"body" binding:"required,min=1,max=1000000"`
	Completed bool      `json:"completed" bson:"completed"`
}
