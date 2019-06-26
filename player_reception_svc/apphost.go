package main

import (
	"fmt"
	"log"

	"github.com/kataras/iris"

	"github.com/vincent-scw/gframe/player_api/kafka"
)

func main() {
	app := iris.Default()

	app.Get("health", func(ctx iris.Context) {
		ctx.Text("I am good.")
	})

	app.Post("message", func(ctx iris.Context) {
		var messages []kafka.TextMessage
		err := ctx.ReadJSON(&messages)
		if err != nil {
			log.Fatalln(err)
			ctx.StatusCode(iris.StatusBadRequest)
		}

		p := kafka.NewProducer()
		defer p.Dispose()
		for i := 0; i < len(messages); i++ {
			messages[i].Key = fmt.Sprintf("address-%s", ctx.RemoteAddr())
		}
		p.EmitMulti(messages)
	})

	fmt.Println("Start player api...")
	app.Run(iris.Addr("127.0.0.1:8080"))
}
