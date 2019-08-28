package connection

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
			h.clients[client.ID] = client
		case clientID := <-h.unregister:
			if client, ok := h.clients[clientID]; ok {
				delete(h.clients, clientID)
				close(client.send)
			}
		}
	}
}

// SendToClient send message to given client
func (h *Hub) SendToClient(clientID string, message *Message) {
	if client, ok := h.clients[clientID]; ok {
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(h.clients, clientID)
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
