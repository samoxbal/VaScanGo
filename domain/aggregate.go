package domain

import (
	"VaScanGo/eventbus"
	"VaScanGo/models"
	"fmt"
	"github.com/satori/go.uuid"
)

const ExperimentAggregateType = "Experiment"

type Aggregate interface {
	StoreEvent(event models.Event, store *eventbus.EventStore, version int) error
}

type BaseAggregate struct {
	ID 		string
	Type 	string
	Events  []models.Event
}

func (ba *BaseAggregate) StoreEvent(event models.Event, store *eventbus.EventStore, version int) error {
	err := store.Save(event, version)
	if err != nil {
		return err
	}
	return nil
}

type ExperimentAggregate struct {
	BaseAggregate
}

func (ea *ExperimentAggregate) HandleCommand(cmd Command, store *eventbus.EventStore, handler eventbus.EventHandle) error {
	switch cmd := cmd.(type) {
	case *CreateExperimentCommand:
		event := models.Event{
			uuid.NewV4().String(),
			CreateExperimentEvent,
			ea.Type,
			ea.ID,
			CreateExperimentEventData{
				cmd.UserID,
				cmd.Name,
				cmd.Description,
				cmd.StartDate,
				cmd.EndDate,
			},
		}
		ea.StoreEvent(event, store, 0)
		handler.HandleEvent(event)
		return nil
	}
	return fmt.Errorf("don't find command")
}
