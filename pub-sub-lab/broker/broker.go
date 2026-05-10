package broker

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/TheAmgadX/backend-labs/pub-sub-lab/models"
)

type Subscriber chan models.Order

type Broker struct {
	subscribers map[string][]Subscriber // list of channels, channel for each subscriber.
	mux         sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]Subscriber),
	}
}

func (b *Broker) getChanSize() (int, error) {
	chan_size := os.Getenv("BROKER_CHANNEL_SIZE")
	size, err := strconv.Atoi(chan_size)

	if err != nil {
		log.Fatal("Error: cannot convert channel size to integer value, check your .env file")
		return 0, err
	}

	return size, nil
}

func (b *Broker) Subscribe(topic string) Subscriber {
	size, err := b.getChanSize()

	if err != nil {
		return nil
	}

	channel := make(Subscriber, size)

	b.mux.Lock()
	defer b.mux.Unlock()

	b.subscribers[topic] = append(b.subscribers[topic], channel)

	return channel
}

func (b *Broker) Publish(topic string, order models.Order) {
	b.mux.Lock()
	defer b.mux.Unlock()

	for _, ch := range b.subscribers[topic] {
		go func(channel chan models.Order) {
			channel <- order
		}(ch)
	}
}
