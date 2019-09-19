package singleton

import (
	"context"
	"sync"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
	"github.com/vincent-scw/gframe/admin_svc/config"
)

var (
	once        sync.Once
	mongoClient *mongo.Client
)

// GetMongoClient return mongo db client
func GetMongoClient() *mongo.Client {
	if config.GetMongoConnection() == "" {
		return nil
	}
	once.Do(func() {
		var err error
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(config.GetMongoConnection()))
		if err != nil {
			log.Fatalf("error connect to mongo %v", err)
		}
	})
	return mongoClient
}
