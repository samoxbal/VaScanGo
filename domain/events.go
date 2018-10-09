package domain

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

type CreateExperimentEventData struct {
	UserID 			string `bson:"userID"`
	Name 			string `bson:"name"`
	Description 	string `bson:"description"`
	StartDate		string `bson:"startDate"`
	EndDate			string `bson:"endDate"`
}