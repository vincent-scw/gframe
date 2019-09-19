package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	
	s "github.com/vincent-scw/gframe/admin_svc/singleton"
)

// GameModel represents a game in database
type GameModel struct {
	// Use generated id instead of bson ObjectID
	ID			string 	`bson:"_id" json:"id,omitempty"`
	CreatedBy 	string	`bson:"createdBy" json:"createdBy"`
	Name 		string	`bson:"name" json:"name"`
}

// GameRepository represents game
type GameRepository struct {
	mongoClient *mongo.Client
	ctx 		context.Context
}

// NewGameRepository returns GameRepository
func NewGameRepository() *GameRepository {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	repo := GameRepository{
		mongoClient: s.GetMongoClient(),
		ctx: ctx,
	}
	return &repo
}

// CreateGame creates a new game
func (repo *GameRepository) CreateGame(model *GameModel) error {
	coll := repo.getCollection()
	_, err := coll.InsertOne(repo.ctx, model)
	if err != nil {
		log.Printf("create game error: %v", err)
	}
	return err
}

// UpdateGame updates a game
func (repo *GameRepository) UpdateGame(model *GameModel) error {
	coll := repo.getCollection()
	filter := bson.M{"_id": bson.M{"$eq": model.ID}}
	_, err := coll.UpdateOne(repo.ctx, filter, bson.M{"$set": model})
	if err != nil {
		log.Printf("update game error: %v", err)
	}
	return err
}

func (repo *GameRepository) getCollection() *mongo.Collection{
	return repo.mongoClient.Database("gframe").Collection("games")
}