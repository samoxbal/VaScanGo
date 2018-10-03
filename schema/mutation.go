package schema

import (
	"VaScanGo/models"
	"VaScanGo/resolvers"
	"github.com/graphql-go/graphql"
)

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createExperiment": &graphql.Field{
			Type: models.ExperimentType,
			Description: "Create new experiment",
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"startDate": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"endDate": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolvers.CreateExperimentResolver,
		},
	},
})
