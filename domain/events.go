package domain

import (
	"VaScanGo/models"
	"github.com/nats-io/go-nats"
)

const (
	CreateExperimentEvent = "CreateExperimentEvent"
)

type EventHandler struct {
	Projector Projector
	ReadModel ReadModel
}

func (e *EventHandler) HandleEvent(event models.Event) {
	natsConn, _ := nats.Connect(nats.DefaultURL)
	connect, _ := nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)
	defer connect.Close()
	connect.Publish(event.Type, event)
}

func (e *EventHandler) ConsumeEvent(event models.Event) {
	e.Projector.Project(event, e.ReadModel)
}

type CreateExperimentEventData struct {
	UserID 			string `bson:"userID"`
	Name 			string `bson:"name"`
	Description 	string `bson:"description"`
	StartDate		string `bson:"startDate"`
	EndDate			string `bson:"endDate"`
}