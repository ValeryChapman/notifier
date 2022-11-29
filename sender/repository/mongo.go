package repository

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var MongoCollection *mongo.Collection
var Ctx = context.TODO()

func InitMongo() {
	mongoServerAddress := os.Getenv("MONGO_SERVER_ADDRESS")

	// Create a new connection
	clientOptions := options.Client().ApplyURI(mongoServerAddress)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to create mongo connection >> %s", err.Error())
	}

	// Test ping
	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping mongo >> %s", err.Error())
	}
	MongoCollection = client.Database(
		viper.GetString("mongodb.database"),
	).Collection(
		viper.GetString("mongodb.collection"),
	)

	// System message
	log.Infof("Successfully connected to MongoDB")
}
