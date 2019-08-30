package connection

import (
	"time"

	"github.com/kataras/iris/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	// ID is client id
	ID string
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan *Message
}

func (c *Client) writePump() {
	for {
		message := <-c.send
		if !c.conn.IsClosed() {
			c.conn.Write(websocket.Message{Namespace: "default", Event: string(message.Type), Body: message.Content})
		}
	}
}

func registerNewClient(hub *Hub, conn *websocket.Conn, id string) {
	client := &Client{ID: id, conn: conn, send: make(chan *Message, 256)}
	hub.register <- client

	go client.writePump()
}

func unregisterClient(hub *Hub, id string) {
	hub.unregister <- id
}
