package services

import (
	"fmt"
	"time"

	"github.com/TheAmgadX/backend-labs/pub-sub-lab/internal/broker"
	"github.com/TheAmgadX/backend-labs/pub-sub-lab/internal/models"
)

type AnalyticsService struct {
	topics []string
}

func NewAnalyticsService() *AnalyticsService {
	topics := []string{"order.created", "order.reserved", "order.paid", "order.shipped", // success scenarios
		"order.out_of_stock", "order.payment_failed", "order.shipping_failed", "order.cancelled", // fail scenarios
	}

	service := AnalyticsService{
		topics: make([]string, 0, len(topics)),
	}

	service.topics = append(service.topics, topics...)

	return &service
}

func (inv *AnalyticsService) Subscribe(broker *broker.Broker) {
	for _, topic := range inv.topics {
		channel := broker.Subscribe(topic)

		go inv.listen(channel)
	}
}

func (inv *AnalyticsService) listen(channel chan models.Order) {
	for order := range channel {
		go inv.process(&order)
	}
}

func (inv *AnalyticsService) process(order *models.Order) {
	time.Sleep(time.Millisecond * 250)
	fmt.Println("Order: ", order.Id, " anlytics done successfully")
}

var _ Service = &AnalyticsService{}
