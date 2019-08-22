package graphql

import (
	"github.com/graphql-go/graphql"

	"github.com/vincent-scw/gframe/admin_svc/simulator"
)

// InjectPlayers type
type InjectPlayers struct {
	// Amount of players
	Amount int `json:"amount"`
}

var injectPlayersType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "InjectPlayers",
		Fields: graphql.Fields{
			"amount": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var injectPlayersField = graphql.Field{
	Type:        injectPlayersType,
	Description: "Inject players by given amount",
	Args: graphql.FieldConfigArgument{
		"amount": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: injectPlayersHandler,
}

func injectPlayersHandler(params graphql.ResolveParams) (interface{}, error) {
	amount, _ := params.Args["amount"].(int)
	simulator.InjectPlayers(amount)
	return InjectPlayers{Amount: amount}, nil
}
