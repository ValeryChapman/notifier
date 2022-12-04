package main

import (
	"consumer/models"
	"consumer/repository"
	"consumer/services"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"time"
)

func main() {
	// Viper config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to open config file >> %s", err.Error())
	}

	// Initialise a logger
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above
	log.SetLevel(log.WarnLevel)

	// Initialise a mongodb connection
	repository.InitMongo()

	// Initialise a rabbitmq connection
	repository.InitRabbit()

	defer repository.RabbitConnection.Close()
	defer repository.RabbitChannel.Close()

	// Make a channel to receive messages into infinite loop
	forever := make(chan bool)

	go func() {

		// Listen messages
		for message := range repository.RabbitMessages {
			go func() {
				// Convert byte message to <services.Mail> structure
				newMail := services.Mail{}
				if err := json.Unmarshal(message.Body, &newMail); err != nil {
					log.Errorf("Failed to decode message body >> %s", err.Error())
				}

				// Convert byte message to <Notification> structure
				newNotification := &models.Notification{}
				if err := json.Unmarshal(message.Body, newNotification); err != nil {
					log.Errorf("Failed to create <Notification> structure >> %s", err.Error())
				}

				// Send a mail
				mailCompleted := true
				if err := services.SendMail(newMail); err != nil {
					log.Errorf("Failed to send mail >> %s", err.Error())
					mailCompleted = false
				}

				// Update a notification
				filter := bson.M{"_id": newNotification.Id}
				update := bson.M{"$set": bson.M{
					"completed":  mailCompleted,
					"updated_at": time.Now().UTC(),
				}}
				if err := models.UpdateNotification(filter, update); err != nil {
					log.Errorf("Failed to update notification >> %s", err.Error())
				}
			}()
		}
	}()

	<-forever
}
