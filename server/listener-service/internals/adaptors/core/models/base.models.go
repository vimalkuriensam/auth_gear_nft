package models

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	Env    map[string]any
	Logger *log.Logger
	Queue  *Queue
}

type Queue struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}
