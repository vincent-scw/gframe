package main

import (
	"fmt"
	"log"
	"net/http"

	gorilla "github.com/gorilla/websocket"
	corsmid "github.com/iris-contrib/middleware/cors"
	_ "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"

	r "github.com/vincent-scw/gframe/redisctl"

	"github.com/vincent-scw/gframe/admin_svc/config"
	"github.com/vincent-scw/gframe/admin_svc/simulator"
	"github.com/vincent-scw/gframe/admin_svc/subscriber"
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
	log.Println("Starting admin service...")

	pubsub := r.NewPubSubClient(config.GetRedisServer())
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
	go subscriber.SubscribePlayer(pubsub, func(msg string) string {
		srv.Broadcast(nil, websocket.Message{Namespace: "default", Event: "console", Body: []byte(msg)})
		return msg
	})
	go subscriber.SubscribeGroup(pubsub, func(msg string) string {
		srv.Broadcast(nil, websocket.Message{Namespace: "default", Event: "console", Body: []byte(msg)})
		return msg
	})

	app := iris.New()

	cors := corsmid.New(corsmid.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	app.Use(cors)
	app.AllowMethods(iris.MethodOptions)

	app.Get("/health", func(ctx iris.Context) {
		ctx.Text("I am good.")
	})

	_ = app.Get("/console", websocket.Handler(srv))

	api := app.Party("/api")
	sim := api.Party("/simulator")
	{
		sim.Post("/inject-players/{count:int}", func(ctx iris.Context) {
			count, _ := ctx.Params().GetInt("count")
			simulator.InjectPlayers(count)
		})
	}

	log.Println(fmt.Sprintf("Serve at %s...", config.GetPort()))
	app.Run(iris.Addr(config.GetPort()), iris.WithoutServerError(iris.ErrServerClosed))
}
