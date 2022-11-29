package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sender/api"
	"sender/repository"
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

	// Initialise a gin application
	router := api.InitRoutes()
	port := ":" + viper.GetString("server.port")
	router.Run(port)
}
