package connection

import (
	"encoding/json"
	"log"
	"net/http"

	gorilla "github.com/gorilla/websocket"
	"github.com/kataras/iris/websocket"
	"github.com/kataras/neffos"

	"github.com/vincent-scw/gframe/contracts"
	"github.com/vincent-scw/gframe/game_svc/auth"
	"github.com/vincent-scw/gframe/game_svc/game"
)

var upgrader = gorilla.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// CreateWebsocket returns Websocket Server
func CreateWebsocket(hub *Hub, onConnect func(user *contracts.User) error,
	onDisconnect func(user *contracts.User)) *neffos.Server {
	gameSvc := game.NewService()
	serverEvents := websocket.Namespaces{
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
			"game": func(nsConn *websocket.NSConn, msg websocket.Message) error {
				//ctx := websocket.GetContext(nsConn.Conn)
				//auth.GetUserFromTokenForWS(ctx)
				log.Printf("[%s] sent: %s", nsConn, string(msg.Body))
				gameEvent := &contracts.GameEvent{}
				json.Unmarshal(msg.Body, gameEvent)
				gameEvent.Type = contracts.EventType_Play
				gameSvc.Play(gameEvent)
				return nil
			},
		},
	}

	srv := websocket.New(
		websocket.GorillaUpgrader(upgrader),
		serverEvents,
	)

	srv.OnConnect = func(c *websocket.Conn) error {
		ctx := websocket.GetContext(c)
		if err := auth.WSJwtHandler.CheckJWT(ctx); err != nil {
			return err
		}

		user, err := auth.GetUserFromTokenForWS(ctx)
		if err != nil {
			return err
		}

		log.Printf("[%s] connected to server.", c.ID())
		// register client
		registerNewClient(hub, c, user.Id)

		if onConnect != nil {
			return onConnect(user)
		}
		return nil
	}

	srv.OnDisconnect = func(c *websocket.Conn) {
		ctx := websocket.GetContext(c)
		user, err := auth.GetUserFromTokenForWS(ctx)
		if err != nil {
			log.Println(err)
			return
		}

		unregisterClient(hub, user.Id)

		log.Printf("[%s] disconnected from the server.", c.ID())
		if onDisconnect != nil {
			onDisconnect(user)
		}
	}

	return srv
}
