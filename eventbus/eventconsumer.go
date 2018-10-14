package eventbus

import (
	"VaScanGo/domain"
	"github.com/nats-io/go-nats"
)

type EventConsumer struct {
	EventMap map[string]domain.EventHandle
}

func (ec *EventConsumer) RegisterHandler(eventType string, handler domain.EventHandle) {
	ec.EventMap[eventType] = handler
}

func (ec *EventConsumer) Start() {
	natsConn, _ := nats.Connect(nats.DefaultURL)
	connect, _ := nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)

	for eventType, eventHandler := range ec.EventMap {
		connect.Subscribe(eventType, eventHandler.ConsumeEvent)
	}
}

func (ec *EventConsumer) GetHandler(eventType string) domain.EventHandle {
	return ec.EventMap[eventType]
}

func MakeEventConsumer() *EventConsumer {
	return &EventConsumer{
		EventMap: make(map[string]domain.EventHandle),
	}
}