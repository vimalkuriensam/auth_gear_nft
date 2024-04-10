package models

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	Env      map[string]any
	DataChan chan any
	Logger   *log.Logger
	Response *JSONResponse
	Error    *ErrorResponse
	Queue    *Queue
}

type JSONResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Status    int       `json:"status"`
	Path      string    `json:"path"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type ReadValue struct {
	B []byte
	D interface{}
}

type Queue struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Messages   map[string]chan Payload
}
