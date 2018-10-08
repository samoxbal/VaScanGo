package domain

import (
	"fmt"
	"github.com/satori/go.uuid"
)

const ExperimentAggregateType = "Experiment"

type BaseAggregate struct {
	ID 		string
	Type 	string
	Events  []Event
}

func (ba *BaseAggregate) StoreEvent(event Event) {
	ba.Events = append(ba.Events, event)
}

type ExperimentAggregate struct {
	BaseAggregate
	userID  		string
	name 			string
	description 	string
	startDate		string
	endDate			string
}

func (ea *ExperimentAggregate) HandleCommand(cmd Command) error {
	switch cmd := cmd.(type) {
	case *CreateExperimentCommand:
		ea.StoreEvent(Event{
			uuid.NewV4().String(),
			CreateExperimentCommandType,
			ea.Type,
			ea.ID,
			CreateExperimentEventData{
				cmd.UserID,
				cmd.Name,
				cmd.Description,
				cmd.StartDate,
				cmd.EndDate,
			},
		})
		return nil
	}
	return fmt.Errorf("don't find command")
}
