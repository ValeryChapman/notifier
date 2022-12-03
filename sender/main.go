package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"sender/api"
	"sender/repository"
	"syscall"
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

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	server := new(Server)
	router := api.InitRoutes()
	port := viper.GetString("server.port")
	go func() {
		if err := server.Run(port, router); err != nil {
			log.Fatalf("Failed while running http server: %s", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Print("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Errorf("Failed on server shutting down: %s", err.Error())
	}

	log.Print("Server shutdown")
}
