package main

import (
	"VaScanGo/controllers"
	"VaScanGo/domain"
	"VaScanGo/eventbus"
	"VaScanGo/models"
	"fmt"
	"github.com/go-bongo/bongo"
	"github.com/kataras/iris"
	"gopkg.in/go-playground/validator.v9"
	"os"
)

func main() {
	app := iris.New()
	bongoConfig := &bongo.Config{
		ConnectionString: "localhost",
		Database:         "VaScan",
	}
	connection, err := bongo.Connect(bongoConfig)
	eventStore := &eventbus.EventStore{
		connection,
	}
	eventConsumer := eventbus.MakeEventConsumer()
	eventConsumer.RegisterHandler(domain.CreateExperimentEvent, &domain.EventHandler{
		Projector: &domain.ExperimentProjector{
			Connection: connection,
		},
		ReadModel: &models.Experiment{},
	})
	eventConsumer.Start()
	if err != nil {
		fmt.Printf("Error MongodbConnection: %s", err)
	}
	var validate *validator.Validate
	validate = validator.New()
	app.Logger().SetOutput(os.Stdout)
	app.Post("/graphql", controllers.GraphQlController(eventStore, eventConsumer, validate))
	app.Post("/token", controllers.TokenController(connection, validate))
	app.Run(
		iris.Addr(":8080"),
		iris.WithConfiguration(iris.TOML("./config/iris.toml")),
		iris.WithoutServerError())
}
