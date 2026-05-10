package services

import (
	"fmt"
	"time"

	"github.com/TheAmgadX/backend-labs/pub-sub-lab/broker"
	"github.com/TheAmgadX/backend-labs/pub-sub-lab/models"
)

type PaymentService struct {
	topics []string
}

func NewPaymentService() *PaymentService {
	topics := []string{"order.reserved", // success scenarios
		"order.shipping_failed", "order.cancelled", // fail scenarios
	}

	service := PaymentService{
		topics: make([]string, 0, len(topics)),
	}

	service.topics = append(service.topics, topics...)

	return &service
}

func (inv *PaymentService) Subscribe(broker *broker.Broker) {
	for _, topic := range inv.topics {
		channel := broker.Subscribe(topic)

		go inv.listen(channel)
	}
}

func (inv *PaymentService) listen(channel chan models.Order) {
	for order := range channel {
		go inv.process(&order)
	}
}

func (inv *PaymentService) process(order *models.Order) {
	time.Sleep(time.Millisecond * 250)
	fmt.Println("Payment for order: ", order.Id, " done successfully")
}

var _ Service = &PaymentService{}
