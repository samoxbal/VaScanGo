package main

import (
	"VaScanGo/controllers"
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
	if err != nil {

	}
	app.Logger().SetOutput(os.Stdout)
	app.Post("/schema", controllers.GraphQlController(connection))
	app.Run(
		iris.Addr(":8080"),
		iris.WithConfiguration(iris.TOML("./config/iris.toml")),
		iris.WithoutServerError())
}
