package main

import (
	"github.com/kataras/iris"

	"github.com/vincent-scw/gframe/player_api/kafka"
)

func main() {
	app := iris.Default()

	app.Get("health", func(ctx iris.Context) {
		//ctx.Text("I am good.")
		p := kafka.NewProducer()
		p.Test()
	})

	app.Run(iris.Addr("127.0.0.1:8080"))
}
