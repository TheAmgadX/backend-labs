package main

import (
	"log"

	"github.com/TheAmgadX/backend-labs/pub-sub-lab/broker"
	"github.com/TheAmgadX/backend-labs/pub-sub-lab/services"
	"github.com/joho/godotenv"
)

/*
	Phase 1 (Subscription Phase):
		all services subscribe to the topics they need.

	Phase 2 (Publish Phase):
		user clicks buy a product, the code calls broker.Publish("order.created", order)

	Phase 3 (The Dispatch):
		The broker sends the messsages to the targeted channels

	Phase 4 (Service Phase):
		Now the services process the event
*/

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Failed to load .env file")
		return
	}

	broker := broker.NewBroker()

	analytics_service := services.NewAnalyticsService()
	inven_service := services.NewInventoryService()
	payment_service := services.NewPaymentService()
	ship_service := services.NewShippingService()

	analytics_service.Subscribe(broker)
	inven_service.Subscribe(broker)
	payment_service.Subscribe(broker)
	ship_service.Subscribe(broker)
}
