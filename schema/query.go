package schema

import (
	"VaScanGo/models"
	"VaScanGo/resolvers"
	"github.com/graphql-go/graphql"
)

var Query = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"experiments": &graphql.Field{
			Type: graphql.NewList(models.ExperimentType),
			Description: "Get all experiments list",
			Args: graphql.FieldConfigArgument{
				"user": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: resolvers.ExperimentsListResolver,
		},
	},
})
