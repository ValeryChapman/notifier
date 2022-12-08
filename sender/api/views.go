package api

import (
	"encoding/json"
	"net/http"
	"sender/api/core"
	"sender/models"
	"sender/repository"
	"sender/utils"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateNotification(c *gin.Context) {
	response := make(chan core.Response)
	go func(context *gin.Context) {
		// Validate request data
		newNotification, status := CreateNotificationSerializer(context)
		if !status {
			response <- core.ErrorResponse(core.ValidationError)
			return
		}

		// Add more info
		newNotification.Id = utils.GenHash([]byte(time.Now().UTC().String()))
		newNotification.Completed = false
		newNotification.CreatedAt = time.Now().UTC()

		// Insert a notification into the database
		if err := models.CreateNotification(newNotification); err != nil {
			log.Errorf("Failed to insert object >> %s", err.Error())
			response <- core.ErrorResponse(core.InternalServerError)
			return
		}

		// Convert <Notification> structure to byte type
		newMailBytes, _ := json.Marshal(newNotification)

		// Create a new message
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        newMailBytes,
		}

		// Publish a new message to the channel
		if err := repository.RabbitChannel.Publish(
			"",
			"notifications",
			false,
			false,
			message,
		); err != nil {
			log.Errorf("Failed to publish a message >> %s", err.Error())
			response <- core.ErrorResponse(core.InternalServerError)
			return
		}
		response <- core.SuccessResponse(http.StatusCreated, newNotification)
	}(c.Copy())
	core.SendResponse(c, <-response)
}

func GetNotification(c *gin.Context) {
	response := make(chan core.Response)
	go func(context *gin.Context) {
		id := context.Param("id")
		filter := bson.M{"_id": id}
		notification, err := models.GetNotification(filter)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				response <- core.ErrorResponse(core.ObjectNotFound)
			} else {
				log.Errorf("Failed to get objects >> %s", err.Error())
				response <- core.ErrorResponse(core.InternalServerError)
			}
			return
		}
		response <- core.SuccessResponse(http.StatusOK, notification)
	}(c.Copy())
	core.SendResponse(c, <-response)
}

func GetNotifications(c *gin.Context) {
	response := make(chan core.Response)
	go func(context *gin.Context) {
		// Validate request data
		requestData, status := GetNotificationsSerializer(context)
		if !status {
			response <- core.ErrorResponse(core.ValidationError)
			return
		}

		// Find objects
		filter := bson.M{}
		if len(requestData.to) != 0 {
			filter = bson.M{"to": bson.M{"$all": requestData.to}}
		}
		notifications, err := models.GetNotifications(
			filter,
			requestData.limit,
			requestData.offset,
		)
		if err != nil {
			log.Errorf("Failed to get objects >> %s", err.Error())
			response <- core.ErrorResponse(core.InternalServerError)
			return
		}
		response <- core.SuccessResponse(http.StatusOK, notifications)
	}(c.Copy())
	core.SendResponse(c, <-response)
}
