package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"os"
)

var RabbitConnection *amqp.Connection
var RabbitChannel *amqp.Channel
var RabbitMessages <-chan amqp.Delivery

func InitRabbit() {
	amqpServerAddress := os.Getenv("RABBIT_SERVER_ADDRESS")

	// Create a new connection
	connection, err := amqp.Dial(amqpServerAddress)
	if err != nil {
		log.Fatalf("Failed to create amqp connection >> %s", err.Error())
	}
	RabbitConnection = connection

	// Open the channel
	channel, err := RabbitConnection.Channel()
	if err != nil {
		log.Fatalf("Failed to open amqp channel >> %s", err.Error())
	}
	RabbitChannel = channel

	// Declare the <notifications> queue
	_, err = channel.QueueDeclare(
		viper.GetString("rabbitmq.queue"),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue >> %s", err.Error())
	}

	// Subscribe to the <notification> queue
	messages, err := RabbitChannel.Consume(
		viper.GetString("rabbitmq.queue"),
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to subscribe for the queue >> %s", err.Error())
	}
	RabbitMessages = messages

	// System message
	log.Infof("Successfully connected to RabbitMQ")
}
