package eventbus

import (
	"VaScanGo/models"
	"github.com/nats-io/go-nats"
)

type EventHandle interface {
	HandleEvent(event models.Event)
	ConsumeEvent(event models.Event)
}

type EventConsumer struct {
	EventMap map[string]EventHandle
}

func (ec *EventConsumer) RegisterHandler(eventType string, handler EventHandle) {
	ec.EventMap[eventType] = handler
}

func (ec *EventConsumer) Start() {
	natsConn, _ := nats.Connect(nats.DefaultURL)
	connect, _ := nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)

	for eventType, eventHandler := range ec.EventMap {
		connect.Subscribe(eventType, eventHandler.ConsumeEvent)
	}
}

func (ec *EventConsumer) GetHandler(eventType string) EventHandle {
	return ec.EventMap[eventType]
}

func MakeEventConsumer() *EventConsumer {
	return &EventConsumer{
		EventMap: make(map[string]EventHandle),
	}
}