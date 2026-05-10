package services

import (
	"fmt"
	"time"

	"github.com/TheAmgadX/backend-labs/pub-sub-lab/broker"
	"github.com/TheAmgadX/backend-labs/pub-sub-lab/models"
)

type ShippingService struct {
	topics []string
}

func NewShippingService() *ShippingService {
	topics := []string{"order.paid", // success scenarios
		"order.cancelled", // fail scenarios
	}

	service := ShippingService{
		topics: make([]string, 0, len(topics)),
	}

	service.topics = append(service.topics, topics...)

	return &service
}

func (inv *ShippingService) Subscribe(broker *broker.Broker) {
	for _, topic := range inv.topics {
		channel := broker.Subscribe(topic)

		go inv.listen(channel)
	}
}

func (inv *ShippingService) listen(channel chan models.Order) {
	for order := range channel {
		go inv.process(&order)
	}
}

func (inv *ShippingService) process(order *models.Order) {
	time.Sleep(time.Millisecond * 250)
	fmt.Println("Order: ", order.Id, " Shipping info done successfully, the orde shipping is now in the dashboard of shipping team.")
}

var _ Service = &ShippingService{}
