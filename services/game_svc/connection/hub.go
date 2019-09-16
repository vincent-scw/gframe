package connection

import "log"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[string]*Client

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan string
}

// NewHub create a hub
func NewHub() *Hub {
	hub := &Hub{
		register:   make(chan *Client),
		unregister: make(chan string),
		clients:    make(map[string]*Client),
	}

	go hub.run()
	return hub
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.User.Id] = client
		case userID := <-h.unregister:
			if client, ok := h.clients[userID]; ok {
				delete(h.clients, userID)
				close(client.send)
			}
		}
	}
}

func (h *Hub) findClient(connID string) *Client {
	for _, client := range h.clients {
		if client.conn.ID() == connID {
			return client
		}
	}
	log.Printf("conn %s cannot be found", connID)
	return nil
}

// SendToClient send message to given client
func (h *Hub) SendToClient(userID string, message *Message) {
	if client, ok := h.clients[userID]; ok {
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(h.clients, userID)
		}
	}
}

// Broadcast send message to all clients
func (h *Hub) Broadcast(message *Message) {
	for id, client := range h.clients {
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(h.clients, id)
		}
	}
}
