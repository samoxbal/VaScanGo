package resolvers

import (
	"github.com/go-bongo/bongo"
	"github.com/graphql-go/graphql"
	"gopkg.in/mgo.v2/bson"
)

func ExperimentsListResolver(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	userId, _ := params.Args["user"].(string)
	connection := rootValue["connection"].(*bongo.Connection)
	results := connection.Collection("experiments").Find(bson.M{"user":userId})
	for results.Next() {

	}
}
