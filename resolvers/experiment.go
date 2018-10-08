package resolvers

import (
	"VaScanGo/models"
	"github.com/go-bongo/bongo"
	"github.com/graphql-go/graphql"
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
	userId, _ := params.Args["user"].(string)
	name, _ := params.Args["name"].(string)
	description, _ := params.Args["description"].(string)
	startDate, _ := params.Args["startDate"].(string)
	endDate, _ := params.Args["endDate"].(string)
	return nil, nil
}
