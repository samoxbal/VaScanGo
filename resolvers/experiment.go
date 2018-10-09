package resolvers

import (
	"VaScanGo/domain"
	"VaScanGo/eventbus"
	"VaScanGo/models"
	"github.com/go-bongo/bongo"
	"github.com/graphql-go/graphql"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

func ExperimentsListResolver(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	userId, _ := params.Args["user"].(string)
	connection := rootValue["connection"].(*bongo.Connection)
	results := connection.Collection("experiments").Find(bson.M{"user":userId})
	experiment := &models.Experiment{}
	var experiments []*models.Experiment
	for results.Next(experiment) {
		experiments = append(experiments, experiment)
	}
	return experiments, nil
}

func CreateExperimentResolver(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	eventStore := rootValue["eventStore"].(*eventbus.EventStore)

	userId, _ := params.Args["user"].(string)
	name, _ := params.Args["name"].(string)
	description, _ := params.Args["description"].(string)
	startDate, _ := params.Args["startDate"].(string)
	endDate, _ := params.Args["endDate"].(string)

	createExperimentCommand := &domain.CreateExperimentCommand{}
	createExperimentCommand.Type = domain.CreateExperimentCommandType
	createExperimentCommand.UserID = userId
	createExperimentCommand.Name = name
	createExperimentCommand.Description = description
	createExperimentCommand.StartDate = startDate
	createExperimentCommand.EndDate = endDate

	experimentAggregate := &domain.ExperimentAggregate{}
	experimentAggregate.ID = uuid.NewV4().String()
	experimentAggregate.Type = domain.ExperimentAggregateType
	experimentAggregate.HandleCommand(createExperimentCommand, eventStore)
	return nil, nil
}
