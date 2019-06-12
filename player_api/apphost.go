package main

import (
	"fmt"
	"log"

	"github.com/kataras/iris"

	"github.com/vincent-scw/gframe/player_api/kafka"
)

type Message struct {
	Content string `json:"content"`
}

func main() {
	app := iris.Default()

	app.Get("health", func(ctx iris.Context) {
		ctx.Text("I am good.")
	})

	app.Post("message", func(ctx iris.Context) {
		var messages []Message
		err := ctx.ReadJSON(&messages)
		if err != nil {
			log.Fatalln(err)
			ctx.StatusCode(iris.StatusBadRequest)
		}

		p := kafka.NewProducer()
		defer p.Dispose()
		for _, msg := range messages {
			log.Printf(msg.Content)
			p.Emit(fmt.Sprintf("address-%s", ctx.RemoteAddr()), msg.Content)
		}
	})

	fmt.Println("Start player api...")
	app.Run(iris.Addr("127.0.0.1:8080"))
}
