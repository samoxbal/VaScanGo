package main

import (
	"github.com/go-bongo/bongo"
	"github.com/kataras/iris"
	"VaScanGo/controllers"
)

func main() {
	app := iris.New()
	bongoConfig := &bongo.Config{
		ConnectionString: "localhost",
		Database:         "vascan",
	}
	connection, err := bongo.Connect(bongoConfig)
	if err != nil {

	}
	app.Post("/schema", controllers.GraphQlController(connection))
	app.Run(
		iris.Addr(":8080"),
		iris.WithConfiguration(iris.TOML("./config/iris.toml")),
		iris.WithoutServerError())
}
