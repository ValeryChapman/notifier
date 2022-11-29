package models

import "time"

type Notification struct {
	Id        string    `json:"id" bson:"_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	To        []string  `json:"to" bson:"to"`
	Subject   string    `json:"subject" bson:"subject"`
	Body      string    `json:"body" bson:"body"`
	Completed bool      `json:"completed" bson:"completed"`
}
