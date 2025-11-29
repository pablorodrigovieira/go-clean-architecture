package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/pablorodrigovieira/go-clean-architecture/pkg/events"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrdersListedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrdersListedHandler(rabbitMQChannel *amqp.Channel) *OrdersListedHandler {
	return &OrdersListedHandler{RabbitMQChannel: rabbitMQChannel}
}

func (h *OrdersListedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
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
		"",           // routing key
		false,
		false,
		msgRabbitmq,
	)
}
