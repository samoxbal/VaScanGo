package main

import (
	"VaScanGo/controllers"
	"fmt"
	"github.com/go-bongo/bongo"
	"github.com/kataras/iris"
	"github.com/looplab/eventhorizon/aggregatestore/events"
	"gopkg.in/go-playground/validator.v9"
	"os"
	eventbus "github.com/looplab/eventhorizon/eventbus/local"
	eventstore "github.com/looplab/eventhorizon/eventstore/mongodb"
)

func main() {
	app := iris.New()
	bongoConfig := &bongo.Config{
		ConnectionString: "localhost",
		Database:         "VaScan",
	}
	connection, err := bongo.Connect(bongoConfig)
	eventStore, _ := eventstore.NewEventStore("localhost:27017", "VaScan")
	eventBus := eventbus.NewEventBus(nil)
	aggregateStore, _ := events.NewAggregateStore(eventStore, eventBus)

	if err != nil {
		fmt.Printf("Error MongodbConnection: %s", err)
	}
	var validate *validator.Validate
	validate = validator.New()
	app.Logger().SetOutput(os.Stdout)
	app.Post("/graphql", controllers.GraphQlController(connection, validate))
	app.Post("/token", controllers.TokenController(connection, validate))
	app.Run(
		iris.Addr(":8080"),
		iris.WithConfiguration(iris.TOML("./config/iris.toml")),
		iris.WithoutServerError())
}
