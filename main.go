package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.Run(
		iris.Addr(":8080"),
		iris.WithConfiguration(iris.TOML("./config/iris.toml")),
		iris.WithoutServerError())
}
