package graphql

import (
	"log"

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
			Fields: graphql.Fields{
				"getGames": &getGamesField,
			},
		},
	)

	mutationType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"injectPlayers": &injectPlayersField,
				"startGame":     &startGameField,
				"cancelGame":    &cancelGameField,
				"updateGame":    &updateGameField,
			},
		},
	)

	var err error
	Schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    queryType,
			Mutation: mutationType,
		},
	)
	if err != nil {
		log.Panicf("error when create GraphQL schema: %v", err)
	}
}
