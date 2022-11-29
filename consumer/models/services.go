package models

import (
	"consumer/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateNotification(filter primitive.M, update primitive.M) error {
	_, err := repository.MongoCollection.UpdateOne(
		repository.Ctx,
		filter,
		update,
	)
	return err
}
