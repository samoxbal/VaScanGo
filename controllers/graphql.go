package controllers

import (
	"github.com/go-bongo/bongo"
	"github.com/graphql-go/graphql"
	"github.com/kataras/iris"
	"VaScanGo/schema"
	"io/ioutil"
)

func GraphQlController(connection *bongo.Connection) iris.Handler {
	return func(ctx iris.Context) {
		rootObject := map[string]interface{}{
			"connection": connection,
		}
		rootSchema, _ := graphql.NewSchema(graphql.SchemaConfig{
			Query: schema.Query,
		})
		req := ctx.Request()
		body, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		result := graphql.Do(graphql.Params{
			RootObject: rootObject,
			Schema: rootSchema,
			RequestString: string(body[:]),
		})
		ctx.JSON(result)
	}
}
