package connection

import (
	"log"
	"time"

	"github.com/kataras/iris/websocket"
	c "github.com/vincent-scw/gframe/contracts"
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
	// User is user
	User *c.User
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send chan *Message
}

func (c *Client) writePump() {
	for {
		message := <-c.send
		if !c.conn.IsClosed() {
			log.Printf("send msg to %s: %s", c.User.Id, string(message.Content))
			c.conn.Write(websocket.Message{Namespace: "default", Event: string(message.Type), Body: message.Content})
		}
	}
}

func registerNewClient(hub *Hub, conn *websocket.Conn, user *c.User) {
	client := &Client{User: user, conn: conn, send: make(chan *Message, 256)}
	hub.register <- client

	go client.writePump()
}

func unregisterClient(hub *Hub, connID string) *c.User {
	client := hub.findClient(connID)
	hub.unregister <- connID
	if client != nil {
		return client.User
	}
	return nil
}
