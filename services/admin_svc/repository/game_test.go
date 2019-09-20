package repository

import (
	"testing"
	"time"

	c "github.com/vincent-scw/gframe/contracts"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewGame(t *testing.T) {
	model := NewGame("test game", "vincent", time.Time{})
	if model.Name != "test game" || model.CreatedBy != "vincent" || !model.RegisterTime.IsZero() {
		t.Error("NewGame model error")
	}
}

func TestCreateGame(t *testing.T) {
	model := GameModel{"test", "testg", "Test Game", 
	time.Date(2009, time.November, 10, 23, 12, 0, 0, time.UTC), 
		time.Date(2009, time.November, 10, 23, 12, 0, 0, time.UTC),
		&c.User{Id:"aabbcc", Name:"vincent"},1, false,
		}
	repo := NewGameRepository()
	repo.UpdateGame(&model)
}