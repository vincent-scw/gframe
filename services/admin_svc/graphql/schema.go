package graphql

import (
	"github.com/graphql-go/graphql"
)

var (
	// Schema is graphql schema
	Schema graphql.Schema
)

func init() {
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
		},
	)

	mutationType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"injectPlayers": &injectPlayersField,
			},
		},
	)

	Schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		},
	)
}
