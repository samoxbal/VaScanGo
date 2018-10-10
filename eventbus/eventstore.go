package eventbus

import (
	"VaScanGo/domain"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type EventStore struct {
	Connection *bongo.Connection
}

func (e *EventStore) Save(event domain.Event, version int) error {
	if version != 0 {
		aggregateRecord := &domain.AggregateRecord{
			AggregateID: event.AggregateID,
			AggregateType: event.AggregateType,
			Events: []domain.Event{event},
		}
		err := e.Connection.Collection(event.AggregateType).Save(aggregateRecord)
		if err != nil {
			return err
		}
	} else {
		aggregateRecord := &domain.AggregateRecord{}
		err := e.Connection.Collection(event.AggregateType).FindOne(bson.M{"AggregateID":event.AggregateID}, aggregateRecord)
		if err != nil {
			return err
		}
		aggregateRecord.Events = append(aggregateRecord.Events, event)
		e.Connection.Collection(event.AggregateType).Save(aggregateRecord)
		return nil
	}
	return nil
 }