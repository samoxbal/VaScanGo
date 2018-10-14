package domain

import "github.com/nats-io/go-nats"

const (
	CreateExperimentEvent = "CreateExperimentEvent"
)

type Event struct {
	ID 				string
	Type 			string
	AggregateType 	string
	AggregateID		string
	Data 			interface{}
}

type EventHandle interface {
	HandleEvent(event Event)
	ConsumeEvent(event Event)
}

type EventHandler struct {
	Projector Projector
	ReadModel ReadModel
}

func (e *EventHandler) HandleEvent(event Event) {
	natsConn, _ := nats.Connect(nats.DefaultURL)
	connect, _ := nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)
	defer connect.Close()
	connect.Publish(event.Type, event)
}

func (e *EventHandler) ConsumeEvent(event Event) {
	e.Projector.Project(event, e.ReadModel)
}

type CreateExperimentEventData struct {
	UserID 			string `bson:"userID"`
	Name 			string `bson:"name"`
	Description 	string `bson:"description"`
	StartDate		string `bson:"startDate"`
	EndDate			string `bson:"endDate"`
}