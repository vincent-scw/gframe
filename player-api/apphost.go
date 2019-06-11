package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()

	app.Get("health", func(ctx iris.Context) {
		ctx.Text("I am good.")
	})

	app.Run(iris.Addr("127.0.0.1:8080"))
}
