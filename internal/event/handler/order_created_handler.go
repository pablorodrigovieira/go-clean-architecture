package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/pablorodrigovieira/go-clean-architecture/pkg/events"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	payload := event.GetPayload()
	fmt.Printf("Orders listed event, payload size maybe: %T\n", payload)
	jsonOutput, _ := json.Marshal(payload)

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
