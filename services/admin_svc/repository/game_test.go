package repository

import (
	"testing"

	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateGame(t *testing.T) {
	model := GameModel{"test", "bb", "bbdaaafs"}
	repo := NewGameRepository()
	repo.UpdateGame(&model)
}