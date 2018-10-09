package eventbus

import (
	"VaScanGo/domain"
	"github.com/go-bongo/bongo"
)

type EventStore struct {
	Connection *bongo.Connection
}

func (e *EventStore) Save(event domain.Event) error {

}