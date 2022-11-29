package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sender/repository"
)

func CreateNotification(n *Notification) error {
	_, err := repository.MongoCollection.InsertOne(repository.Ctx, n)
	return err
}

func GetNotification(filter primitive.M) (Notification, error) {
	var result Notification
	err := repository.MongoCollection.FindOne(repository.Ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, err
}

func GetNotifications(filter primitive.M, limit int64, offset int64) ([]Notification, error) {
	var result []Notification
	options := options.Find().SetLimit(limit).SetSkip(offset)
	cursor, err := repository.MongoCollection.Find(repository.Ctx, filter, options)
	if err = cursor.All(repository.Ctx, &result); err != nil {
		return result, err
	}
	return result, nil
}
