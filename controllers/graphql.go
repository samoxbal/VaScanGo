package controllers

import (
	"VaScanGo/utils"
	"github.com/go-bongo/bongo"
	"github.com/graphql-go/graphql"
	"github.com/kataras/iris"
	"VaScanGo/schema"
	"VaScanGo/models"
	"gopkg.in/go-playground/validator.v9"
)

func GraphQlController(connection *bongo.Connection, validate *validator.Validate) iris.Handler {
	return func(ctx iris.Context) {
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
