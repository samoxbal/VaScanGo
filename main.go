package main

import (
	"VaScanGo/controllers"
	"VaScanGo/domain"
	"VaScanGo/eventbus"
	"VaScanGo/models"
	"fmt"
	"github.com/go-bongo/bongo"
	"github.com/kataras/iris"
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
		Connection: connection,
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
	app.Logger().SetOutput(os.Stdout)
	app.Post("/graphql", controllers.GraphQlController(eventStore, eventConsumer, connection))
	app.Post("/token", controllers.TokenController(connection))
	app.Run(
		iris.Addr(":8080"),
		iris.WithConfiguration(iris.TOML("./config/iris.toml")),
		iris.WithoutServerError())
}
