#!/bin/bash

while ! nc -z rabbitmq 5672
do
  echo "Failure connected to RabbitMQ"
  sleep 3
done

while ! nc -z mongodb 27017
do
  echo "Failure connected to MongoDB"
  sleep 3
done

go build -ldflags="-s -w" -o consumer .
./consumer