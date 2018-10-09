package domain

const (
	CreateExperimentCommandType = "CreateExperimentCommand"
)

type Command interface {
	GetType() string
	GetAggregateID() string
}

type BaseCommand struct {
	Type 			string
	AggregateID   	string
}

type CreateExperimentCommand struct {
	BaseCommand
	UserID 			string
	Name 			string
	Description 	string
	StartDate		string
	EndDate			string
}

func (c *BaseCommand) GetType() string { return c.Type }
func (c *BaseCommand) GetAggregateID() string { return c.AggregateID }
