package domain

import (
	"VaScanGo/models"
	"github.com/go-bongo/bongo"
)

type ReadModel interface {
	GetModelID() string
}

type Projector interface {
	Project(event Event, readModel ReadModel) error
}

type ExperimentProjector struct {
	Connection *bongo.Connection
}

func (ep *ExperimentProjector) Project(event Event, readModel ReadModel) error {
	experiment := readModel.(*models.Experiment)
	eventData := event.Data.(CreateExperimentEventData)

	experiment.AggregateID = event.AggregateID
	experiment.User = eventData.UserID
	experiment.Name = eventData.Name
	experiment.Description = eventData.Description
	experiment.StartDate = eventData.StartDate
	experiment.EndDate = eventData.EndDate

	err := ep.Connection.Collection(event.AggregateType).Save(experiment)
	if err != nil {
		return err
	}
	return nil
}