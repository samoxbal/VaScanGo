package domain

const (
	CreateExperimentCommandType = "CreateExperiment"
)

type Command interface {
	GetType() string
	GetAggregateType() string
	GetAggregateID() string
}

type BaseCommand struct {
	Type 			string
	AggregateID   	string
	AggregateType 	string
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
func (c *BaseCommand) GetAggregateType() string { return c.AggregateType }
func (c *BaseCommand) GetAggregateID() string { return c.AggregateID }
