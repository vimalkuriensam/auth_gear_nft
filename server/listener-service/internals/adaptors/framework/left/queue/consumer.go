package queue

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/adaptors/core/models"
	"github.com/vimalkuriensam/auto_gear_nft/listener-service/internals/ports"
)

type Adaptor struct {
	config ports.ConfigPort
	routes ports.RoutesPort
}

var adaptor *Adaptor

func Initialize(config ports.ConfigPort, routes ports.RoutesPort) *Adaptor {
	adaptor = &Adaptor{
		config: config,
		routes: routes,
	}
	return adaptor
}

func GetAdaptor() *Adaptor {
	return adaptor
}

func (qAd *Adaptor) Connect() {
	config := qAd.config.GetConfig()
	env := config.Env
	user := env["queue_user"].(string)
	pass := env["queue_password"].(string)
	host := env["queue_host"].(string)
	port := env["queue_port"].(string)
	rabbitURL := fmt.Sprintf("amqp://%s:%s@%s:%s", user, pass, host, port)
	var counts int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection
	for {
		c, err := amqp.Dial(rabbitURL)
		if err != nil {
			config.Logger.Println("rabbitmq not connected", err)
			counts++
		} else {
			config.Logger.Println("connected to rabbitmq")
			connection = c
			break
		}
		if counts > 5 {
			config.Logger.Println(err)
			os.Exit(1)
		}
		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		config.Logger.Println("backing off...")
		time.Sleep(backoff)
	}
	config.Queue.Connection = connection
	channel, err := connection.Channel()
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
	config.Queue.Channel = channel
	err = channel.ExchangeDeclare("st-exchange", "direct", true, false, false, false, nil)
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
	_, err = channel.QueueDeclare("queue1_request", true, false, false, true, nil)
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
	_, err = channel.QueueDeclare("queue1_response", true, false, false, true, nil)
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
	channel.QueueBind("queue1_request", "queue1_key", "st-exchange", false, nil)
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
	channel.QueueBind("queue1_response", "queue2_key", "st-exchange", false, nil)
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
	_, err = channel.QueueDeclare("queue2_response", true, false, false, true, nil)
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
	_, err = channel.QueueDeclare("queue2_request", true, false, false, true, nil)
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
	channel.QueueBind("queue2_response", "queue3_key", "st-exchange", false, nil)
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
	channel.QueueBind("queue2_request", "queue4_key", "st-exchange", false, nil)
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
}

func (qAd *Adaptor) Listen() {
	config := qAd.config.GetConfig()
	channel := config.Queue.Channel
	defer channel.Close()
	msgs1, err := channel.Consume("queue1_request", "", false, false, false, false, nil)
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
	msgs2, err := channel.Consume("queue2_request", "", false, false, false, false, nil)
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
	processMessage := func(msgs <-chan amqp.Delivery, label string) {
		for msg := range msgs {
			payload := models.Payload{}
			json.Unmarshal(msg.Body, &payload)
			fmt.Println("Processing", label)
			qAd.routes.RouteRequest(payload)
			msg.Ack(false)
		}
	}
	go processMessage(msgs1, "queue1_request")
	go processMessage(msgs2, "queue2_request")
	fmt.Println("Waiting for messages...")
	select {}
}
