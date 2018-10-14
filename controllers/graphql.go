package controllers

import (
	"VaScanGo/eventbus"
	"VaScanGo/models"
	"VaScanGo/schema"
	"VaScanGo/utils"
	"github.com/go-bongo/bongo"
	"github.com/graphql-go/graphql"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
)

func GraphQlController(eventStore *eventbus.EventStore, eventConsumer *eventbus.EventConsumer, connection *bongo.Connection) iris.Handler {
	return func(ctx iris.Context) {
		var validate *validator.Validate
		validate = validator.New()
		var req models.GraphQlRequest
		validate.RegisterStructValidation(utils.ValidateQueryStruct)
		if err := ctx.ReadJSON(&req); err != nil {
			ctx.Application().Logger().Infof("Error read request: %s", err)
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}
		validationError := validate.Struct(req)
		if validationError != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}
		rootObject := map[string]interface{}{
			"connection": connection,
			"eventStore": eventStore,
			"eventConsumer": eventConsumer,
			"ctx": ctx,
		}
		rootSchema, _ := graphql.NewSchema(graphql.SchemaConfig{
			Query: schema.Query,
			Mutation: schema.Mutation,
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
