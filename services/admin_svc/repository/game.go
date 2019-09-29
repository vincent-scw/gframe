package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	s "github.com/vincent-scw/gframe/admin_svc/singleton"
	c "github.com/vincent-scw/gframe/contracts"
	u "github.com/vincent-scw/gframe/util"
)

// GameModel represents a game in database
type GameModel struct {
	// Use generated id instead of bson ObjectID
	ID           string    `bson:"_id" json:"id,omitempty"`
	CreatedBy    string    `bson:"createdBy" json:"createdBy"`
	CreatedTime  time.Time `json:"createdTime" json:"createdTime"`
	Name         string    `bson:"name" json:"name"`
	RegisterTime time.Time `bson:"registerTime" json:"registerTime"`
	StartTime    time.Time `bson:"startTime" json:"startTime"`
	Winner       *c.User   `bson:"winner" json:"winner"`
	Type         int       `bson:"type" json:"type"`
	IsCancelled  bool      `bson:"isCancelled" json:"isCancelled"`
	IsCompleted  bool      `bson:"isCompleted" json:"isCompleted"`
	IsStarted    bool      `bson:"isStarted" json:"isStarted"`
}

// NewGame return a game model
func NewGame(name, createdBy string, reg time.Time) *GameModel {
	model := GameModel{
		ID:           u.NextRandom(),
		Name:         name,
		CreatedBy:    createdBy,
		CreatedTime:  time.Now().UTC(),
		RegisterTime: reg,
		Type:         1,
		IsCancelled:  false,
		IsCompleted:  false,
		IsStarted:    false,
	}
	return &model
}

// GameRepository represents game
type GameRepository struct {
	mongoClient *mongo.Client
	ctx         context.Context
}

// NewGameRepository returns GameRepository
func NewGameRepository() *GameRepository {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	repo := GameRepository{
		mongoClient: s.GetMongoClient(),
		ctx:         ctx,
	}
	return &repo
}

// CreateGame creates a new game
func (repo *GameRepository) CreateGame(model *GameModel) (*GameModel, error) {
	coll := repo.getCollection()
	result, err := coll.InsertOne(repo.ctx, model)
	if err != nil {
		log.Printf("create game error: %v", err)
	}
	newModel, err := repo.GetOne(result.InsertedID.(string))
	return newModel, err
}

// GetOne returns one game by id
func (repo *GameRepository) GetOne(id string) (*GameModel, error) {
	coll := repo.getCollection()
	filter := bson.M{"_id": id}
	var result GameModel
	err := coll.FindOne(repo.ctx, filter).Decode(&result)
	if err != nil {
		log.Printf("get one game error: %v", err)
	}
	return &result, err
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

// FindGames returns games for created user
func (repo *GameRepository) FindGames(createdBy string) ([]*GameModel, error) {
	coll := repo.getCollection()
	cur, err := coll.Find(repo.ctx, bson.M{"createdBy": createdBy})
	if err != nil {
		log.Printf("find games error: %v", err)
		return nil, err
	}
	defer cur.Close(repo.ctx)
	var games []*GameModel
	for cur.Next(repo.ctx) {
		var result GameModel
		err := cur.Decode(&result)
		if err != nil {
			log.Printf("find games error: %v", err)
			return nil, err
		}
		games = append(games, &result)
	}
	if err := cur.Err(); err != nil {
		log.Printf("error %v", err)
	}
	return games, err
}

func (repo *GameRepository) getCollection() *mongo.Collection {
	return repo.mongoClient.Database("gframe").Collection("games")
}
