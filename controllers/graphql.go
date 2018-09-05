package controllers

import (
	"github.com/go-bongo/bongo"
	"github.com/graphql-go/graphql"
	"github.com/kataras/iris"
	"VaScanGo/schema"
	"VaScanGo/models"
)

func GraphQlController(connection *bongo.Connection) iris.Handler {
	return func(ctx iris.Context) {
		var req models.GraphQlRequest
		if err := ctx.ReadJSON(&req); err != nil {
			ctx.Application().Logger().Infof("Error read request: %s", err)
		}
		rootObject := map[string]interface{}{
			"connection": connection,
			"ctx": ctx,
		}
		rootSchema, _ := graphql.NewSchema(graphql.SchemaConfig{
			Query: schema.Query,
		})
		ctx.Application().Logger().Info(req.Query)
		result := graphql.Do(graphql.Params{
			RootObject: rootObject,
			Schema: rootSchema,
			RequestString: req.Query,
		})
		ctx.JSON(result)
	}
}
