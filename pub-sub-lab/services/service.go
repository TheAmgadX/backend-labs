package services

import (
	"github.com/TheAmgadX/backend-labs/pub-sub-lab/broker"
	"github.com/TheAmgadX/backend-labs/pub-sub-lab/models"
)

type Service interface {
	Subscribe(broker *broker.Broker)
	listen(chan models.Order)
	process(*models.Order)
}
