package eventbus

import (
	"VaScanGo/models"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

type EventStore struct {
	Connection *bongo.Connection
}

func (e *EventStore) Save(event models.Event, version int) error {
	if version != 0 {
		aggregateRecord := &models.AggregateRecord{
			AggregateID: event.AggregateID,
			AggregateType: event.AggregateType,
			Events: []models.Event{event},
		}
		err := e.Connection.Collection(event.AggregateType).Save(aggregateRecord)
		if err != nil {
			return err
		}
	} else {
		aggregateRecord := &models.AggregateRecord{}
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