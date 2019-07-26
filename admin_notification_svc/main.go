package main

import (
	"log"
	"net/http"

	gorilla "github.com/gorilla/websocket"
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"

	e "github.com/vincent-scw/gframe/events"
	r "github.com/vincent-scw/gframe/redisctl"
)

var upgrader = gorilla.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var serverEvents = websocket.Namespaces{
	"default": websocket.Events{
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
			log.Printf("[%s] sent: %s", nsConn, string(msg.Body))
			return nil
		},
	},
}

func main() {
	log.Println("Starting admin notification service...")

	pubsub := r.NewPubSubClient("40.83.112.48:6379")
	defer pubsub.Close()

	srv := websocket.New(
		websocket.GorillaUpgrader(upgrader),
		serverEvents,
	)

	srv.OnConnect = func(c *websocket.Conn) error {
		log.Printf("[%s] connected to server.", c.ID())
		return nil
	}

	srv.OnDisconnect = func(c *websocket.Conn) {
		log.Printf("[%s] disconnected from the server.", c.ID())
	}

	log.Println("Subscribe to Redis...")
	go pubsub.Subscribe(e.GroupChannel, func(msg string) {
		srv.Broadcast(nil, websocket.Message{Namespace: "default", Event: "console", Body: []byte(msg)})
	})

	app := iris.New()

	app.Get("/health", func(ctx iris.Context) {
		ctx.Text("I am good.")
	})

	_ = app.Get("/console", websocket.Handler(srv))

	log.Println("Serve at localhost:10010...")
	app.Run(iris.Addr(":10010"), iris.WithoutServerError(iris.ErrServerClosed))
}
