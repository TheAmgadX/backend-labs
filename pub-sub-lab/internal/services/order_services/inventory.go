package services

import (
	"fmt"
	"time"

	"github.com/TheAmgadX/backend-labs/pub-sub-lab/internal/broker"
	"github.com/TheAmgadX/backend-labs/pub-sub-lab/internal/models"
)

type InventoryService struct {
	topics []string
}

func NewInventoryService() *InventoryService {
	topics := []string{"order.created", // success scenarios
		"order.payment_failed", "order.shipping_failed", "order.cancelled", // fail scenarios
	}

	service := InventoryService{
		topics: make([]string, 0, len(topics)),
	}

	service.topics = append(service.topics, topics...)

	return &service
}

func (inv *InventoryService) Subscribe(broker *broker.Broker) {
	for _, topic := range inv.topics {
		channel := broker.Subscribe(topic)

		go inv.listen(channel)
	}
}

func (inv *InventoryService) listen(channel chan models.Order) {
	for order := range channel {
		go inv.process(&order)
	}
}

func (inv *InventoryService) process(order *models.Order) {
	time.Sleep(time.Millisecond * 250)
	fmt.Println("Inventory Updated, order: ", order.Id, " placed successfully")
}

var _ Service = &InventoryService{}
