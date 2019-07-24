package main

import (
	"log"

	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"

	e "github.com/vincent-scw/gframe/events"
	r "github.com/vincent-scw/gframe/redisctl"
)

const enableJWT = true
const namespace = "default"

var serverEvents = websocket.Namespaces{
	namespace: websocket.Events{
		websocket.OnNamespaceConnected: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// with `websocket.GetContext` you can retrieve the Iris' `Context`.
			ctx := websocket.GetContext(nsConn.Conn)

			log.Printf("[%s] connected to namespace [%s] with IP [%s]",
				nsConn, msg.Namespace,
				ctx.RemoteAddr())
			return nil
		},
		websocket.OnNamespaceDisconnect: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			log.Printf("[%s] disconnected from namespace [%s]", nsConn, msg.Namespace)
			return nil
		},
		"console": func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// room.String() returns -> NSConn.String() returns -> Conn.String() returns -> Conn.ID()
			log.Printf("[%s] sent: %s", nsConn, string(msg.Body))

			// Write message back to the client message owner with:
			// nsConn.Emit("console", msg)
			// Write message to all except this client with:
			nsConn.Conn.Server().Broadcast(nsConn, msg)
			return nil
		},
	},
}

func main() {
	log.Println("Starting player notification service...")

	pubsub := r.NewPubSubClient("40.83.112.48:6379")
	defer pubsub.Close()

	websocketServer := websocket.New(
		websocket.GorillaUpgrader(upgrader),
		serverEvents,
	)

	log.Println("Subscribe to Redis...")
	go pubsub.Subscribe(e.GroupChannel, func(msg string) {
		websocketServer.Broadcast(nil,
			websocket.Message{Namespace: namespace, Event: "console", Body: []byte(msg)})
	})

	app := iris.New()

	app.Get("/health", func(ctx iris.Context) {
		ctx.Text("I am good.")
	})

	_ = app.Get("/console", websocket.Handler(websocketServer))

	log.Println("Serve at localhost:9010...")
	app.Run(iris.Addr(":9010"), iris.WithoutServerError(iris.ErrServerClosed))
}
