package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/adaptors/core/models"
	"github.com/vimalkuriensam/auth_gear_nft/gateway-service/internals/ports"
)

type Adaptor struct {
	config ports.ConfigPort
}

func Initialize(config ports.ConfigPort) *Adaptor {
	return &Adaptor{
		config: config,
	}
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
			config.Logger.Println("rabbitmq not connected")
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
	channel.QueueBind("queue1_response", "queue2_key", "st-exchange", false, nil)
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
	channel.QueueBind("queue1_request", "queue1_key", "st-exchange", false, nil)
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
}

func (qAd *Adaptor) Listen() {
	config := qAd.config.GetConfig()
	channel := config.Queue.Channel
	defer channel.Close()
	msgs, err := channel.Consume("queue1_response", "", true, false, false, false, nil)
	if err != nil {
		config.Logger.Println(err)
		os.Exit(1)
	}
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			payload := models.Payload{}
			json.Unmarshal(msg.Body, &payload)
			qAd.config.SetMessage(payload.Id, payload)
		}
	}()
	fmt.Println("Waiting for messages...")
	<-forever
}

func (qAd *Adaptor) Emit(event models.Payload) error {
	channel := qAd.config.GetConfig().Queue.Channel
	payloadAsbytes, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return channel.PublishWithContext(context.Background(), "st-exchange", "queue1_key", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        payloadAsbytes,
	})

}
