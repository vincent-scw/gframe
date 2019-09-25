package graphql

import (
	"errors"
	"time"

	"github.com/graphql-go/graphql"

	repo "github.com/vincent-scw/gframe/admin_svc/repository"
)

var playerType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Player",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var gameType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Game",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"createdBy": &graphql.Field{
				Type: graphql.String,
			},
			"createdTime": &graphql.Field{
				Type: graphql.DateTime,
			},
			"registerTime": &graphql.Field{
				Type: graphql.DateTime,
			},
			"startTime": &graphql.Field{
				Type: graphql.DateTime,
			},
			"winner": &graphql.Field{
				Type: playerType,
			},
			"type": &graphql.Field{
				Type: graphql.Int,
			},
			"isCancelled": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var getGamesField = graphql.Field{
	Type:        graphql.NewList(gameType),
	Description: "Get games created by owner",
	Args: graphql.FieldConfigArgument{
		"owner": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		owner, ok := p.Args["owner"].(string)
		if ok {
			gameRepo := repo.NewGameRepository()
			return gameRepo.FindGames(owner)
		}
		return nil, nil
	},
}

var startGameField = graphql.Field{
	Type:        graphql.Boolean,
	Description: "Start a game",
	Args: graphql.FieldConfigArgument{
		"gameId": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		gameID, ok := p.Args["gameId"].(string)
		if ok {
			gameRepo := repo.NewGameRepository()
			game, err := checkGame(gameRepo, gameID)
			if err != nil {
				return false, err
			}

			game.StartTime = time.Now().UTC()
			err = gameRepo.UpdateGame(game)

			// TODO: publish this event
			return err == nil, err
		}
		return nil, nil
	},
}

var cancelGameField = graphql.Field{
	Type:        graphql.Boolean,
	Description: "Cancel a game",
	Args: graphql.FieldConfigArgument{
		"gameId": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		gameID, ok := p.Args["gameId"].(string)
		if ok {
			gameRepo := repo.NewGameRepository()
			game, err := checkGame(gameRepo, gameID)
			if err != nil {
				return false, err
			}

			game.IsCancelled = true
			err = gameRepo.UpdateGame(game)
			return err == nil, err
		}
		return nil, nil
	},
}

var updateGameField = graphql.Field{
	Type:        graphql.Boolean,
	Description: "Update a game",
	Args: graphql.FieldConfigArgument{
		"game": &graphql.ArgumentConfig{
			Type: gameType,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		gameID, ok := p.Args["game"].(string)
		if ok {
			gameRepo := repo.NewGameRepository()
			game, err := checkGame(gameRepo, gameID)
			if err != nil {
				return false, err
			}

			// If game can be update, it must has not started
			game.IsCancelled = false
			game.StartTime = time.Time{}
			err = gameRepo.UpdateGame(game)
			return err == nil, err
		}
		return nil, nil
	},
}

func checkGame(gameRepo *repo.GameRepository, gameID string) (*repo.GameModel, error) {
	game, err := gameRepo.GetOne(gameID)
	if err != nil {
		return nil, err
	}
	if game.IsCancelled {
		return nil, errors.New("the game has been cancelled")
	}
	if !game.StartTime.IsZero() {
		return nil, errors.New("the game has already started")
	}
	return game, nil
}
