package graphql

import (
	"github.com/functionalfoundry/graphqlws"
)

var subscriptionManager = graphqlws.NewSubscriptionManager(&Schema)

// GraphqlwsHandler is websocket handler
var GraphqlwsHandler = graphqlws.NewHandler(graphqlws.HandlerConfig{
	SubscriptionManager: subscriptionManager,
	// Authenticate: func(authToken string) (interface{}, error) {

	// }
})

// Broadcast send message to all clients
func Broadcast(msg string) string {
	subscriptions := subscriptionManager.Subscriptions()
	for conn := range subscriptions {
		//conn.ID()
		//conn.User()
		data := graphqlws.DataMessagePayload{
			Data: msg,
		}
		for _, sub := range subscriptions[conn] {
			sub.SendData(&data)
		}
	}
	return msg
}
